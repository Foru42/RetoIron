package routes

import (
	"josu-foruria/src/controllers"
	"josu-foruria/src/dao"
	"josu-foruria/src/database"
	"josu-foruria/src/handlers"

	"github.com/gin-gonic/gin"
)

func RoutesUsuarios(r *gin.Engine, db *database.DB) {
	usuarioService := &controllers.UsuarioService{
		DAO: &dao.UsuarioDAO{
			DB: db,
		},
	}

	r.GET("/usuarios", handlers.GetUsuariosHandler(usuarioService))
	r.GET("/usuarios/:id", handlers.GetUsuarioByIdHandler(usuarioService))
	r.POST("/usuarios", handlers.PostUsuario(usuarioService))
	r.PUT("/usuarios", handlers.PutUsuario(usuarioService))
	r.DELETE("/usuarios", handlers.DeleteUsuario(usuarioService))
}
