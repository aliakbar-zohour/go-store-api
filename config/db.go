package config

import (
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)

var DB *mongo.Database

func ConnectDB() {
    mongoURI := "mongodb+srv://admin:admin@cluster0.fhr17mi.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    serverAPI := options.ServerAPI(options.ServerAPIVersion1)
    opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

    client, err := mongo.Connect(ctx, opts)
    if err != nil {
        log.Fatal("❌ MongoDB connection error:", err)
    }

    if err := client.Ping(ctx, readpref.Primary()); err != nil {
        log.Fatal("❌ MongoDB ping error:", err)
    }

    DB = client.Database("store")
    log.Println("✅ Connected to MongoDB Atlas!")
}
