// main.go
package main

import (
	"fmt"
	"net/http"
	"time"

	"RetoIronChip/database"
	"RetoIronChip/routes"

	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/middleware/stdlib"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

func main() {
	var db database.DB
	db.InitDB() // Inicializar la conexión a MongoDB
	defer db.CloseDB()

	rate := limiter.Rate{Period: 1 * time.Second, Limit: 10}
	store := memory.NewStore()
	instance := limiter.New(store, rate)
	rateLimiterMiddleware := stdlib.NewMiddleware(instance)

	fmt.Println("Servidor escuchando en el puerto 8080...")
	http.Handle("/usuarios", rateLimiterMiddleware.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		routes.HandleUsuarios(w, r, &db)
	})))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
