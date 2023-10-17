package asyncgroup

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
	"time"
)

var (
	ErrTimeout = errors.New("time out task group")
)

var (
	defaultTimeout   = time.Second * 5
	defaultMaxWorker = runtime.NumCPU()
)

type AsyncGroup interface {
	// Process runs the tasks concurrently.
	//
	// It returns the first non-nil error when processing tasks, otherwise it returns nil.
	Process(ctx context.Context, tasks ...func(context.Context) error) error

	// ProcessWithTimeout runs the tasks concurrently with the context timeout.
	//
	// It returns the first non-nil error when processing tasks, otherwise it returns nil.
	ProcessWithTimeout(ctx context.Context, timeout time.Duration, tasks ...func(context.Context) error) error
}

type asyncGroup struct {
	// semaphore ensure that the worker goroutines not exceed the limit.
	// Default value is runtime.NumCPU()
	semaphore chan struct{}
}

func New(maxWorker int) (AsyncGroup, func()) {
	if maxWorker < 1 {
		maxWorker = defaultMaxWorker
	}
	ag := &asyncGroup{
		semaphore: make(chan struct{}, maxWorker),
	}
	return ag, ag.cancel
}

func (ag *asyncGroup) cancel() {
	for i := 0; i < len(ag.semaphore); i++ {
		ag.semaphore <- struct{}{}
	}
	close(ag.semaphore)
}

func (ag *asyncGroup) Process(
	ctx context.Context,
	tasks ...func(context.Context) error,
) error {
	return ag.processWithTimeout(ctx, defaultTimeout, tasks...)
}

func (ag *asyncGroup) ProcessWithTimeout(
	ctx context.Context,
	timeout time.Duration,
	tasks ...func(context.Context) error,
) error {
	if timeout <= 0 {
		timeout = defaultTimeout
	}
	return ag.processWithTimeout(ctx, timeout, tasks...)
}

// processWithTimeout runs the tasks with context timeout.
//
// If any task exceeded the deadline, returns ErrTimeout.
func (ag *asyncGroup) processWithTimeout(
	ctx context.Context,
	timeout time.Duration,
	tasks ...func(context.Context) error,
) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	taskCount := len(tasks)

	results := make(chan error, taskCount)
	defer close(results)

	// Ensure that context cancel only call once when a non-nil error happened.
	cancelOnce := new(sync.Once)

	for i := 0; i < taskCount; i++ {
		ag.semaphore <- struct{}{} // acquire
		go func(i int) {
			defer func() {
				<-ag.semaphore // release
			}()

			// Check context before processing
			select {
			case <-ctx.Done():
				results <- context.Cause(ctx)
				return
			default:
			}

			errTaskCh := make(chan error)
			go func() {
				errTaskCh <- ag.processTask(ctx, tasks[i])
				close(errTaskCh)
			}()

			select {
			case <-ctx.Done():
				results <- context.Cause(ctx)
			case e := <-errTaskCh:
				if e != nil {
					// Cancel other tasks in the group which haven't done yet.
					cancelOnce.Do(func() {
						cancel()
					})
				}
				results <- e
			}
		}(i)
	}

	var err error
	for i := 0; i < taskCount; i++ {
		if e := <-results; e != nil && err == nil {
			err = e
		}
	}

	if errors.Is(err, context.DeadlineExceeded) {
		err = ErrTimeout
	}
	return err
}

// processTask executes the task and can handle the panic when task is processing.
func (ag *asyncGroup) processTask(ctx context.Context, task func(context.Context) error) (err error) {
	defer func() {
		if e := recover(); e != nil {
			errMsg := string(debug.Stack())
			fmt.Println(errMsg)

			i := 1
			_, file, line, ok := runtime.Caller(i) // skip the first frame (panic itself)
			for ok && strings.Contains(file, "runtime/") {
				i++
				_, file, line, ok = runtime.Caller(i)
			}

			// Include the file and line number info in the error, if runtime.Caller returned ok.
			if ok {
				err = fmt.Errorf("panic [%s:%d]: %v", file, line, e)
			} else {
				err = fmt.Errorf("panic: %v", e)
			}
		}
	}()

	return task(ctx)
}
