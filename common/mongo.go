package common

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

func WithMongodbTransaction(
	ctx context.Context,
	db *mongo.Database,
	fn func(context.Context) error,
) error {
	session, err := db.Client().StartSession()
	if err != nil {
		return NewServerErr(err)
	}
	defer session.EndSession(ctx)

	callback := func(sessCtx mongo.SessionContext) error {
		if e := session.StartTransaction(); e != nil {
			return NewServerErr(e)
		}

		if e := fn(sessCtx); e != nil {
			return e
		}

		if e := session.CommitTransaction(context.Background()); e != nil {
			return NewServerErr(e)
		}

		return nil
	}

	if err = mongo.WithSession(ctx, session, callback); err != nil {
		if abortErr := session.AbortTransaction(context.Background()); abortErr != nil {
			return err
		}
		return err
	}

	return nil
}
