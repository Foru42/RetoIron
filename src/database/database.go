// database/db.go
package database

import (
	"context"
	"fmt"
	"log"
	"os"
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

	// Leer URI y nombre de BD desde las variables de entorno
	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")
	/*
		mongoURI = "mongodb://localhost:27017/"
		dbName = "RetoIronChip"
	*/
	if mongoURI == "" || dbName == "" {
		log.Fatal("Faltan variables de entorno necesarias")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("No se pudo hacer ping a la base de datos: %v", err)
	}

	fmt.Println("Conectado a MongoDB")
	db.Client = client
	db.UsersCollection = client.Database(dbName).Collection("usuarios")
}

func (db *DB) CloseDB() {
	if db.Client != nil {
		db.Client.Disconnect(context.Background())
	}
}
