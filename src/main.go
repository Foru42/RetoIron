package main

import (
    "fmt"
    "net/http"
    "RetoIronChip/routes"
    "RetoIronChip/database"
)

func main() {
    // Inicializa la base de datos
    database.InitDB()
    defer database.DB.Close() // Cierra la conexión al salir

    fmt.Println("Servidor escuchando en el puerto 8080...")

    // Configura las rutas
    http.HandleFunc("/usuarios", routes.HandleUsuarios)

    // Inicia el servidor
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Error al iniciar el servidor:", err)
    }
}