package controllers

import (
	"context"
	"josu-foruria/src/dao"
	"josu-foruria/src/models"
	"josu-foruria/src/utils"
	"josu-foruria/src/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsuarioService struct {
	DAO *dao.UsuarioDAO
}

func (service *UsuarioService) GetUsuarios(c *gin.Context) {
	ctx := context.Background()
	usuarios, err := service.DAO.GetUsuarios(ctx)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Error al obtener usuarios")
		return
	}
	utils.RespondWithJSON(c, http.StatusOK, usuarios)
}

func (service *UsuarioService) CreateUsuario(c *gin.Context) {
	var usuario models.Usuario
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
	err := service.DAO.CreateUsuario(ctx, usuario)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Error al crear el usuario")
		return
	}

	utils.RespondWithJSON(c, http.StatusCreated, usuario)
}

func (service *UsuarioService) UpdateUsuario(c *gin.Context) {
	var usuario models.Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Datos inválidos: formato JSON incorrecto")
		return
	}

	if !validators.IsValidEmail(usuario.Email) || !validators.IsValidText(usuario.Name, 50) || !validators.IsValidText(usuario.Surname, 50) {
		utils.RespondWithError(c, http.StatusBadRequest, "Datos inválidos")
		return
	}

	ctx := context.Background()
	err := service.DAO.UpdateUsuario(ctx, usuario)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithJSON(c, http.StatusOK, gin.H{"message": "Usuario actualizado correctamente"})
}

func (service *UsuarioService) DeleteUsuario(c *gin.Context) {
	name := c.Query("name")
	if !validators.IsValidText(name, 50) {
		utils.RespondWithError(c, http.StatusBadRequest, "El nombre del usuario es obligatorio")
		return
	}

	ctx := context.Background()
	err := service.DAO.DeleteUsuario(ctx, name)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	utils.RespondWithJSON(c, http.StatusOK, gin.H{"message": "Usuario eliminado correctamente"})
}
