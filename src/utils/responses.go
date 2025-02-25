package utils

import "github.com/gin-gonic/gin"

type ErrorResponse struct {
	Message string `json:"Error"`
}

// RespondWithError envía una respuesta JSON con un mensaje de error usando Gin
func RespondWithError(c *gin.Context, code int, message string) {
	c.JSON(code, ErrorResponse{Message: message})
}

// RespondWithJSON envía una respuesta JSON con datos y código HTTP usando Gin
func RespondWithJSON(c *gin.Context, code int, payload interface{}) {
	c.JSON(code, payload)
}
