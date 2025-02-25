package controllers

import (
	"context"
	"encoding/json"
	"josu-foruria/src/database"
	"josu-foruria/src/models"
	"josu-foruria/src/utils"
	"josu-foruria/src/validators"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Obtener todos los usuarios
func GetUsuarios(w http.ResponseWriter, r *http.Request, db *database.DB) {
	var usuarios []models.Usuario
	ctx := context.Background()

	cursor, err := db.UsersCollection.Find(ctx, bson.M{})
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error al obtener usuarios")
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var usuario models.Usuario
		if err := cursor.Decode(&usuario); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Error al leer los datos")
			return
		}
		usuarios = append(usuarios, usuario)
	}

	utils.RespondWithJSON(w, http.StatusOK, usuarios)
}

// Crear usuario
func CreateUsuario(w http.ResponseWriter, r *http.Request, db *database.DB) {
	var usuario models.Usuario
	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Datos inválidos: formato JSON incorrecto")
		return
	}

	if !validators.IsValidEmail(usuario.Email) {
		utils.RespondWithError(w, http.StatusBadRequest, "El email tiene un formato inválido: xxx@xxx.xxx")
		return
	}

	if !validators.IsValidText(usuario.Name, 50) {
		utils.RespondWithError(w, http.StatusBadRequest, "El nombre es obligatorio y debe tener menos de 50 caracteres")
		return
	}

	if !validators.IsValidText(usuario.Surname, 50) {
		utils.RespondWithError(w, http.StatusBadRequest, "El apellido es obligatorio y debe tener menos de 50 caracteres")
		return
	}

	// Verificar si ya existe un usuario con el mismo email
	ctx := context.Background()
	count, err := db.UsersCollection.CountDocuments(ctx, bson.M{"email": usuario.Email})
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error al verificar el email")
		return
	}

	if count > 0 {
		utils.RespondWithError(w, http.StatusBadRequest, "El email ya está registrado")
		return
	}

	// Asignar un ID único al nuevo usuario
	usuario.ID = primitive.NewObjectID()

	// Insertar el nuevo usuario
	_, err = db.UsersCollection.InsertOne(ctx, usuario)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error al crear el usuario")
		return
	}

	// Responder con el usuario creado
	utils.RespondWithJSON(w, http.StatusCreated, usuario)
}

// Actualizar usuario
func UpdateUsuario(w http.ResponseWriter, r *http.Request, db *database.DB) {
	var usuario models.Usuario

	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Datos invalidos, formato Json invalido")
		return
	}

	// Filtrar por el nombre del usuario
	filter := bson.M{"name": usuario.Name}

	if !validators.IsValidEmail(usuario.Email) {
		utils.RespondWithError(w, http.StatusBadRequest, "El email tiene un formato inválido: xxx@xxx.xxx")
		return
	}

	if !validators.IsValidText(usuario.Name, 50) {
		utils.RespondWithError(w, http.StatusBadRequest, "El nombre es obligatorio y debe tener menos de 50 caracteres")
		return
	}
	if !validators.IsValidText(usuario.Surname, 50) {
		utils.RespondWithError(w, http.StatusBadRequest, "El apellido es obligatorio y debe tener menos de 50 caracteres")
		return
	}

	// Verificar si el usuario existe
	count, err := db.UsersCollection.CountDocuments(context.Background(), filter)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Erros al verificar la existencia del usuario")
		return
	}

	if count == 0 {
		utils.RespondWithError(w, http.StatusBadRequest, "Usuario Inexistente")
		return
	}
	// Realizar el update
	update := bson.M{"$set": bson.M{
		"name":    usuario.Name,
		"surname": usuario.Surname,
		"email":   usuario.Email,
	}}

	result, err := db.UsersCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Error al actualizar usuario")
		return
	}

	if result.MatchedCount == 0 {
		utils.RespondWithError(w, http.StatusBadRequest, "No se pudo actualizar el usuario")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Usuario actualizado correctamente\n"))
}

// Borrar usuario
func DeleteUsuario(w http.ResponseWriter, r *http.Request, db *database.DB) {
	name := r.URL.Query().Get("name")

	if !validators.IsValidText(name, 50) {
		utils.RespondWithError(w, http.StatusBadRequest, "El nombre del usuario es obligatorio")
		return
	}

	// Filtrar por el nombre del usuario
	filter := bson.M{"name": name}

	// Eliminar el usuario
	result, err := db.UsersCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Error al eliminar usuario")
		return
	}

	if result.DeletedCount == 0 {
		utils.RespondWithError(w, http.StatusBadRequest, "Usuario no encontrado")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Usuario eliminado correctamente\n"))
}
