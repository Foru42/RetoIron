package dao

import (
	"context"
	"fmt"
	"josu-foruria/src/database"
	"josu-foruria/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UsuarioDAO struct {
	DB *database.DB
}

func (dao *UsuarioDAO) GetUsuarios(ctx context.Context) ([]models.Usuario, error) {
	var usuarios []models.Usuario
	cursor, err := dao.DB.UsersCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var usuario models.Usuario
		if err := cursor.Decode(&usuario); err != nil {
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}

	return usuarios, err
}

func (dao *UsuarioDAO) CreateUsuario(ctx context.Context, usuario models.Usuario) error {
	usuario.ID = primitive.NewObjectID()
	_, err := dao.DB.UsersCollection.InsertOne(ctx, usuario)
	return err
}

func (dao *UsuarioDAO) UpdateUsuario(ctx context.Context, usuario models.Usuario) error {
	count, err := dao.DB.UsersCollection.CountDocuments(ctx, bson.M{"name": usuario.Name})
	if err != nil || count == 0 {
		return fmt.Errorf("usuario no encontrado")
	}
	update := bson.M{
		"$set": bson.M{
			"name":    usuario.Name,
			"surname": usuario.Surname,
			"email":   usuario.Email,
		},
	}
	_, err = dao.DB.UsersCollection.UpdateOne(ctx, bson.M{"name": usuario.Name}, update)
	return err
}

func (dao *UsuarioDAO) DeleteUsuario(ctx context.Context, name string) error {
	count, err := dao.DB.UsersCollection.CountDocuments(ctx, bson.M{"name": name})
	if err != nil || count == 0 {
		return fmt.Errorf("usuario no encontrado")
	}

	_, err = dao.DB.UsersCollection.DeleteOne(ctx, bson.M{"name": name})
	return err
}
