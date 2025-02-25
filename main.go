// main.go
package main

import (
	"fmt"
	"josu-foruria/src/database"
	"josu-foruria/src/routes"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/ulule/limiter"
	"github.com/ulule/limiter/drivers/middleware/stdlib"
	"github.com/ulule/limiter/drivers/store/memory"
)

func main() {

	port := os.Getenv("PORT")

	var db database.DB
	db.InitDB()
	defer db.CloseDB()

	limit := os.Getenv("RATE_LIMIT")
	limitInt, _ := strconv.Atoi(limit)
	if port == "" || limit == "" {
		log.Fatal("Faltan variables de entorno necesarias")
	}

	rate := limiter.Rate{Period: 5 * time.Second, Limit: int64(limitInt)}
	store := memory.NewStore()
	instance := limiter.New(store, rate)
	rateLimiterMiddleware := stdlib.NewMiddleware(instance)

	fmt.Printf("Servidor escuchando en el puerto %s...\n", port)
	http.Handle("/usuarios", rateLimiterMiddleware.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		routes.HandleUsuarios(w, r, &db)
	})))

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
