package routes

import (
    "net/http"
    "RetoIronChip/controllers"
)

func HandleUsuarios(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        controllers.GetUsuarios(w, r)
    } else {
        http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
    }
}
