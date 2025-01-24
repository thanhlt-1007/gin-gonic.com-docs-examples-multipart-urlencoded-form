package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func postMessageHandler(context *gin.Context) {
    message := context.PostForm("message")
    name := context.DefaultPostForm("name", "anonymous")
    context.JSON(
        http.StatusOK,
        gin.H {
            "message": message,
            "name": name,
        },
    )
}

func main() {
    engine := gin.Default()
    engine.POST("/message", postMessageHandler)
    engine.Run()
}
