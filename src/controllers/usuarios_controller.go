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


    // Consulta a la base de datos
    rows, err := db.Query("SELECT id, name, surname, email FROM usuarios")
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


func CreateUsuario(w http.ResponseWriter, r *http.Request) {
    var usuario models.Usuario
    if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
        http.Error(w, "Datos inválidos", http.StatusBadRequest)
        return
    }

    db := database.GetDB()


    query := `INSERT INTO usuarios (name, surname, email) VALUES (?, ?, ?)`
    result, err := db.Exec(query, usuario.Name, usuario.Surname, usuario.Email)
    if err != nil {
        http.Error(w, "Error al crear el usuario", http.StatusInternalServerError)
        return
    }

    id, _ := result.LastInsertId()
    usuario.ID = int(id)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(usuario)
}

func UpdateUsuario(w http.ResponseWriter, r *http.Request) {
    var usuario models.Usuario
    if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
        http.Error(w, "Datos inválidos", http.StatusBadRequest)
        return
    }

    db := database.GetDB()


    query := `UPDATE usuarios SET name = ?, surname = ?, email = ? WHERE id = ?`
    _, err := db.Exec(query, usuario.Name, usuario.Surname, usuario.Email, usuario.ID)
    if err != nil {
        http.Error(w, "Error al actualizar el usuario", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Usuario actualizado correctamente"))
}
func DeleteUsuario(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, "ID requerido", http.StatusBadRequest)
        return
    }

    db := database.GetDB()


    query := `DELETE FROM usuarios WHERE id = ?`
    _, err := db.Exec(query, id)
    if err != nil {
        http.Error(w, "Error al eliminar el usuario", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Usuario eliminado correctamente"))
}
