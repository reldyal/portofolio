package main

import (
    "log"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "github.com/reldyal/backend/internal/handler"
)

func main() {
    // Load .env
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }

    r := gin.Default()

    // CORS middleware — penting biar Vue bisa hit API
    r.Use(func(c *gin.Context) {
        c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
        c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        c.Header("Access-Control-Allow-Headers", "Content-Type")
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        c.Next()
    })

    // Routing — group by feature
    api := r.Group("/api")
    {
        api.POST("/contact", handler.SendContact)
        api.GET("/github/repos", handler.GetRepos)
    }

    r.Run(":8080") // listen di port 8080
}