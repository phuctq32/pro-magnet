package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"pro-magnet/components/appcontext"
	"time"
)

func main() {
	// DB
	db, cancel := connectMongoDB()
	defer cancel()
	appCtx := appcontext.NewAppContext(db)
	log.Println("Database name: " + appCtx.GetDBConn().Name())

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	err := router.Run(fmt.Sprintf(":%v", os.Getenv("PORT")))
	if err != nil {
		log.Fatal(err)
	}
}

func connectMongoDB() (*mongo.Database, context.CancelFunc) {
	mongoUri := os.Getenv("MONGO_CONN_STRING")
	mongoDBName := os.Getenv("MONGO_DATABASE_NAME")

	opts := options.Client().ApplyURI(mongoUri)
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(mongoDBName)

	log.Println("DB connected!")

	return db, cancel
}
