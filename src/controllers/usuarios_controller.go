package controllers

import (
    "encoding/json"
    "net/http"
    "regexp"
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


// Valida que un email sea válido
func isValidEmail(email string) bool {
    regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
    re := regexp.MustCompile(regex)
    return re.MatchString(email)
}

// Valida que un texto no esté vacío y no exceda cierta longitud
func isValidText(text string, maxLength int) bool {
    return len(text) > 0 && len(text) <= maxLength
}

func CreateUsuario(w http.ResponseWriter, r *http.Request) {
    var usuario models.Usuario
    if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
        http.Error(w, "Datos inválidos: formato JSON incorrecto", http.StatusBadRequest)
        return
    }

    // Validación de campos
    if !isValidText(usuario.Name, 50) {
        http.Error(w, "El nombre es obligatorio y debe tener menos de 50 caracteres", http.StatusBadRequest)
        return
    }
    if !isValidText(usuario.Surname, 50) {
        http.Error(w, "El apellido es obligatorio y debe tener menos de 50 caracteres", http.StatusBadRequest)
        return
    }
    if !isValidEmail(usuario.Email) {
        http.Error(w, "El email no tiene un formato válido", http.StatusBadRequest)
        return
    }

    db := database.GetDB()

    query := `INSERT INTO usuarios (name, surname, email) VALUES (?, ?, ?)`
    result, err := db.Exec(query, usuario.Name, usuario.Surname, usuario.Email)
    if err != nil {
        http.Error(w, "Error al crear el usuario. Es posible que el email ya esté registrado.", http.StatusInternalServerError)
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
        http.Error(w, "Datos inválidos: formato JSON incorrecto", http.StatusBadRequest)
        return
    }

    // Validación de campos
    if usuario.ID <= 0 {
        http.Error(w, "El ID del usuario es inválido", http.StatusBadRequest)
        return
    }
    if !isValidText(usuario.Name, 50) {
        http.Error(w, "El nombre es obligatorio y debe tener menos de 50 caracteres", http.StatusBadRequest)
        return
    }
    if !isValidText(usuario.Surname, 50) {
        http.Error(w, "El apellido es obligatorio y debe tener menos de 50 caracteres", http.StatusBadRequest)
        return
    }
    if !isValidEmail(usuario.Email) {
        http.Error(w, "El email no tiene un formato válido", http.StatusBadRequest)
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
