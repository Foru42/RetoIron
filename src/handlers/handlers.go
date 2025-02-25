package handlers

import (
	"josu-foruria/src/controllers"
	"josu-foruria/src/dao"
	"josu-foruria/src/database"

	"github.com/gin-gonic/gin"
)

func InitUserHandlers(r *gin.Engine, db *database.DB) {

	usuarioService := &controllers.UsuarioService{
		DAO: &dao.UsuarioDAO{
			DB: db,
		},
	}

	r.GET("/usuarios", usuarioService.GetUsuarios)
	r.GET("/usuarios/:id", usuarioService.GetUsuarioId)
	r.POST("/usuarios", usuarioService.CreateUsuario)
	r.PUT("/usuarios", usuarioService.UpdateUsuario)
	r.DELETE("/usuarios", usuarioService.DeleteUsuario)
}
