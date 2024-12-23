package controllers

import (
    "encoding/json"
    "net/http"
    "RetoIronChip/models"
    "RetoIronChip/database"
)

func GetUsuarios(w http.ResponseWriter, r *http.Request) {
    // Obtiene la conexión a la base de datos
    db := database.GetDB()
    defer db.Close()

    // Consulta a la base de datos
    rows, err := db.Query("SELECT id, name, surname,email FROM usuarios")
    if err != nil {
        http.Error(w, "Error al obtener los usuarios", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    // Recorre los resultados
    var usuarios []models.Usuario
    for rows.Next() {
        var usuario models.Usuario
        if err := rows.Scan(&usuario.ID, &usuario.Name, &usuario.Surname, &usuario.Email); err != nil {
            http.Error(w, "Error al leer los datos", http.StatusInternalServerError)
            return
        }
        usuarios = append(usuarios, usuario)
    }

    // Devuelve la respuesta como JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(usuarios)
}
