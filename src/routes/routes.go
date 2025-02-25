package routes

import (
	"josu-foruria/src/controllers"
	"josu-foruria/src/dao"
	"josu-foruria/src/database"

	"github.com/gin-gonic/gin"
)

func HandleUsuarios(r *gin.Engine, db *database.DB) {
	usuarioService := &controllers.UsuarioService{
		DAO: &dao.UsuarioDAO{
			DB: db,
		},
	}

	usuarios := r.Group("/usuarios")
	{
		usuarios.GET("", usuarioService.GetUsuarios)
		usuarios.POST("", usuarioService.CreateUsuario)
		usuarios.PUT("", usuarioService.UpdateUsuario)
		usuarios.DELETE("", usuarioService.DeleteUsuario)
	}
}
