// database/db.go
package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	Client          *mongo.Client
	UsersCollection *mongo.Collection
}

func (db *DB) InitDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Error al conectar con MongoDB: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("No se pudo hacer ping a MongoDB: %v", err)
	}

	fmt.Println("Conectado a MongoDB")
	db.Client = client
	db.UsersCollection = client.Database("RetoIronChip").Collection("usuarios")
}

func (db *DB) CloseDB() {
	if db.Client != nil {
		db.Client.Disconnect(context.Background())
	}
}
