package handlers

import (
	"josu-foruria/src/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsuariosHandler(service *controllers.UsuarioService) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.JSON(http.StatusServiceUnavailable, gin.H{"status": "Error interno del servidor"})
			}
		}()

		service.GetUsuarios(c)
	}
}

func GetUsuarioByIdHandler(service *controllers.UsuarioService) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.JSON(http.StatusServiceUnavailable, gin.H{"status": "Error interno del servidor"})
			}
		}()

		service.GetUsuarioId(c)
	}
}

func PostUsuario(service *controllers.UsuarioService) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.JSON(http.StatusServiceUnavailable, gin.H{"status": "Error interno del servidor"})
			}
		}()

		service.CreateUsuario(c)
	}
}
func PutUsuario(service *controllers.UsuarioService) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.JSON(http.StatusServiceUnavailable, gin.H{"status": "Error interno del servidor"})
			}
		}()

		service.UpdateUsuario(c)
	}
}

func DeleteUsuario(service *controllers.UsuarioService) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.JSON(http.StatusServiceUnavailable, gin.H{"status": "Error interno del servidor"})
			}
		}()

		service.DeleteUsuario(c)
	}
}
