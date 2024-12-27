package routes

import (
    "net/http"
    "RetoIronChip/controllers"
)

// Funcion para las rutas de los metodos GET/CREATE/UPDATE/DELETE
func HandleUsuarios(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        controllers.GetUsuarios(w, r)
    case http.MethodPost:
        controllers.CreateUsuario(w, r)
    case http.MethodPut:
        controllers.UpdateUsuario(w, r)
    case http.MethodDelete:
        controllers.DeleteUsuario(w, r)
    default:
        http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
    }
}