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

func (dao *UsuarioDAO) GetUsuarioId(ctx context.Context, id string) (models.Usuario, error) {
	var usuario models.Usuario
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return usuario, fmt.Errorf("ID inválido")
	}

	err = dao.DB.UsersCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&usuario)
	if err != nil {
		return usuario, fmt.Errorf("usuario no encontrado")
	}

	return usuario, nil
}

func (dao *UsuarioDAO) CreateUsuario(ctx context.Context, usuario *models.Usuario) (*models.Usuario, error) {
	usuario.ID = primitive.NewObjectID()

	_, err := dao.DB.UsersCollection.InsertOne(ctx, usuario)
	if err != nil {
		return nil, err
	}
	return usuario, nil
}

func (dao *UsuarioDAO) UpdateUsuario(ctx context.Context, usuario models.Usuario, id string) (*models.Usuario, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("ID inválido")
	}
	count, err := dao.DB.UsersCollection.CountDocuments(ctx, bson.M{"_id": objectID})
	if err != nil || count == 0 {
		return nil, fmt.Errorf("usuario no encontrado")
	}
	update := bson.M{
		"$set": bson.M{
			"name":    usuario.Name,
			"surname": usuario.Surname,
			"email":   usuario.Email,
		},
	}
	_, err = dao.DB.UsersCollection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		return nil, err
	}
	var updatedUsuario models.Usuario
	err = dao.DB.UsersCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&updatedUsuario)
	if err != nil {
		return nil, err
	}

	return &updatedUsuario, nil
}

func (dao *UsuarioDAO) DeleteUsuario(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("ID inválido")
	}

	count, err := dao.DB.UsersCollection.CountDocuments(ctx, bson.M{"_id": objectID})
	if err != nil || count == 0 {
		return fmt.Errorf("usuario no encontrado")
	}

	_, err = dao.DB.UsersCollection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}
