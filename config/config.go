package config

import (
    "context"
    "log"
    "os"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
    mongoURI := os.Getenv("MONGO_URI")
    username := os.Getenv("MONGO_USER")
    password := os.Getenv("MONGO_PASSWORD")

    clientOptions := options.Client().ApplyURI(mongoURI)
    
    // Add authentication if username and password are provided
    if username != "" && password != "" {
        clientOptions.SetAuth(options.Credential{
            Username: username,
            Password: password,
        })
    }

    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal("❌ MongoDB Connection Error:", err)
    }

    // Ping the database to verify connection
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal("❌ MongoDB Ping Error:", err)
    }

    log.Println("✅ Connected to MongoDB!")
    DB = client.Database("go_gin_app")
}
