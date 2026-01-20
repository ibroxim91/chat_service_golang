package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
    "chat-service/internal/hub"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

func ChatHandler(h *hub.Hub) gin.HandlerFunc {
    return func(c *gin.Context) {
        conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
        if err != nil {
            return
        }
        client := &hub.Client{Conn: conn, Send: make(chan []byte, 256), Hub: h}
        h.Register <- client

        go client.WritePump()
        go client.ReadPump()
    }
}
