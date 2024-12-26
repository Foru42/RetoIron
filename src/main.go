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
    // Inicializa la base de datos
    database.InitDB()
    defer database.DB.Close()

    // Configuracion del rate limiting: 10 solicitudes por segundo
    rate := limiter.Rate{
        Period: 1 * time.Second,
        Limit:  10,
    }
    store := memory.NewStore()
    instance := limiter.New(store, rate)
    rateLimiterMiddleware := stdlib.NewMiddleware(instance)

    // Configura el servidor
    fmt.Println("Servidor escuchando en el puerto 8080...")
    http.Handle("/usuarios", rateLimiterMiddleware.Handler(http.HandlerFunc(routes.HandleUsuarios)))

    // Inicia el servidor
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Error al iniciar el servidor:", err)
    }
}
