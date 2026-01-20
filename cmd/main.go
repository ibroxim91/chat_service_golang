package main

import (
    "github.com/gin-gonic/gin"
    "chat-service/internal/handlers"
    "chat-service/internal/hub"
)

func main() {
    r := gin.Default()
    h := hub.NewHub()
    go h.Run()

    r.GET("/ws", handlers.ChatHandler(h))

    r.Run(":8080")
}
