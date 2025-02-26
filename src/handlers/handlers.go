package handlers

import (
	"context"
	"josu-foruria/src/controllers"
	"josu-foruria/src/models"
	"josu-foruria/src/utils"
	"josu-foruria/src/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsuariosHandler(service *controllers.UsuarioService) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				utils.RespondWithError(c, http.StatusServiceUnavailable, "Error interno del servidor")
			}
		}()

		ctx := context.Background()
		usuarios, err := service.GetUsuarios(ctx)
		if err != nil {
			utils.RespondWithError(c, http.StatusInternalServerError, "Error al obtener usuarios")
			return
		}

		utils.RespondWithJSON(c, http.StatusOK, usuarios)
	}
}
func GetUsuarioByIdHandler(service *controllers.UsuarioService) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				utils.RespondWithError(c, http.StatusServiceUnavailable, "Error interno del servidor")
			}
		}()

		id := c.Param("id")
		if id == "" {
			utils.RespondWithError(c, http.StatusBadRequest, "El ID del usuario es obligatorio")
			return
		}

		ctx := context.Background()
		usuario, err := service.GetUsuarioId(ctx, id)
		if err != nil {
			utils.RespondWithError(c, http.StatusNotFound, "Usuario no encontrado")
			return
		}

		utils.RespondWithJSON(c, http.StatusOK, usuario)
	}
}

func PostUsuario(service *controllers.UsuarioService) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				utils.RespondWithError(c, http.StatusServiceUnavailable, "Error interno del servidor")
			}
		}()

		var usuario *models.Usuario
		if err := c.ShouldBindJSON(&usuario); err != nil {
			utils.RespondWithError(c, http.StatusBadRequest, "Datos inválidos: formato JSON incorrecto")
			return
		}

		if !validators.IsValidEmail(usuario.Email) {
			utils.RespondWithError(c, http.StatusBadRequest, "El email tiene un formato inválido")
			return
		}

		if !validators.IsValidText(usuario.Name, 50) || !validators.IsValidText(usuario.Surname, 50) {
			utils.RespondWithError(c, http.StatusBadRequest, "Nombre y apellido deben tener menos de 50 caracteres")
			return
		}

		ctx := context.Background()
		usuario, err := service.CreateUsuario(ctx, *usuario)
		if err != nil {
			utils.RespondWithError(c, http.StatusInternalServerError, "Error al crear el usuario")
			return
		}

		utils.RespondWithJSON(c, http.StatusCreated, usuario)
	}
}

func PutUsuario(service *controllers.UsuarioService) gin.HandlerFunc {
	var usuario models.Usuario
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.JSON(http.StatusServiceUnavailable, gin.H{"status": "Error interno del servidor"})
			}
		}()
		id := c.Param("id")
		if id == "" {
			utils.RespondWithError(c, http.StatusBadRequest, "El ID del usuario es obligatorio")
			return
		}

		if err := c.ShouldBindJSON(&usuario); err != nil {
			utils.RespondWithError(c, http.StatusBadRequest, "Datos inválidos: formato JSON incorrecto")
			return
		}

		if !validators.IsValidEmail(usuario.Email) || !validators.IsValidText(usuario.Name, 50) || !validators.IsValidText(usuario.Surname, 50) {
			utils.RespondWithError(c, http.StatusBadRequest, "Datos inválidos")
			return
		}

		ctx := context.Background()
		updatedUsuario, err := service.UpdateUsuario(ctx, usuario, id)
		if err != nil {
			utils.RespondWithError(c, http.StatusNotFound, "Usuario no encontrado")
			return
		}

		utils.RespondWithJSON(c, http.StatusOK, updatedUsuario)
	}
}
func DeleteUsuario(service *controllers.UsuarioService) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.JSON(http.StatusServiceUnavailable, gin.H{"status": "Error interno del servidor"})
			}
		}()
		id := c.Param("id")
		if id == "" {
			utils.RespondWithError(c, http.StatusBadRequest, "El ID del usuario es obligatorio")
			return
		}
		ctx := context.Background()
		err := service.DeleteUsuario(ctx, id)
		if err != nil {
			utils.RespondWithError(c, http.StatusNotFound, "Usuario no encontrado")
			return
		}
		utils.RespondWithJSON(c, http.StatusOK, gin.H{"status": "Usuario eliminado"})
	}
}
