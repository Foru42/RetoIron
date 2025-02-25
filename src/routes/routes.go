package routes

import (
	"josu-foruria/src/controllers"
	"josu-foruria/src/database"
	"net/http"
)

// Funcion para las rutas de los metodos GET/CREATE/UPDATE/DELETE
func HandleUsuarios(w http.ResponseWriter, r *http.Request, db *database.DB) {

	switch r.Method {
	case http.MethodGet:
		controllers.GetUsuarios(w, r, db)
	case http.MethodPost:
		controllers.CreateUsuario(w, r, db)
	case http.MethodPut:
		controllers.UpdateUsuario(w, r, db)
	case http.MethodDelete:
		controllers.DeleteUsuario(w, r, db)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}
