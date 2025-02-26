package controllers

import (
	"context"
	"josu-foruria/src/dao"
	"josu-foruria/src/models"
)

type UsuarioService struct {
	DAO *dao.UsuarioDAO
}

func (service *UsuarioService) GetUsuarios(ctx context.Context) ([]models.Usuario, error) {
	usuarios, err := service.DAO.GetUsuarios(ctx)
	if err != nil {
		return nil, err
	}
	return usuarios, nil
}

func (service *UsuarioService) GetUsuarioId(ctx context.Context, id string) (*models.Usuario, error) {
	usuario, err := service.DAO.GetUsuarioId(ctx, id)
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}

func (service *UsuarioService) CreateUsuario(ctx context.Context, usuario models.Usuario) error {
	err := service.DAO.CreateUsuario(ctx, usuario)
	if err != nil {
		return err
	}
	return nil
}

func (service *UsuarioService) UpdateUsuario(ctx context.Context, usuario models.Usuario) error {
	err := service.DAO.UpdateUsuario(ctx, usuario)
	if err != nil {
		return err
	}
	return nil
}

func (service *UsuarioService) DeleteUsuario(ctx context.Context, name string) error {
	err := service.DAO.DeleteUsuario(ctx, name)
	if err != nil {
		return err
	}
	return nil
}
