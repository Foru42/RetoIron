package main

import (
	"fmt"
	"josu-foruria/src/database"
	"josu-foruria/src/routes"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/time/rate"
)

func rateLimiter(limiter *rate.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "too many requests"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}

	port := os.Getenv("PORT")

	var db database.DB
	db.InitDB()
	defer db.CloseDB()

	limit := os.Getenv("RATE_LIMIT")
	limitInt, err := strconv.Atoi(limit)
	if err != nil || port == "" || limit == "" {
		log.Fatal("Faltan variables de entorno necesarias")
	}

	limiter := rate.NewLimiter(rate.Limit(limitInt), 5)

	r := gin.Default()

	r.Use(rateLimiter(limiter))

	routes.RoutesUsuarios(r, &db)

	fmt.Printf("Servidor escuchando en el puerto %s...\n", port)
	if err := r.Run(":" + port); err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
