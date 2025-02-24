package controllers

import (
	"RetoIronChip/database"
	"RetoIronChip/models"
	"RetoIronChip/validators"
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Funcion para recoger los usuarios
func GetUsuarios(w http.ResponseWriter, r *http.Request, db *database.DB) {
	var usuarios []models.Usuario
	ctx := context.Background()

	cursor, err := db.UsersCollection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Error al obtener usuarios", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var usuario models.Usuario
		if err := cursor.Decode(&usuario); err != nil {
			http.Error(w, "Error al leer los datos", http.StatusInternalServerError)
			return
		}
		usuarios = append(usuarios, usuario)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuarios)
}

func CreateUsuario(w http.ResponseWriter, r *http.Request, db *database.DB) {
	var usuario models.Usuario
	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		http.Error(w, "Datos inválidos: formato JSON incorrecto", http.StatusBadRequest)
		return
	}

	if !validators.IsValidEmail(usuario.Email) {
		http.Error(w, "Error El email tiene un formato invalido : xxx@xxx.xxx", http.StatusInternalServerError)
		return
	}

	if !validators.IsValidText(usuario.Name, 50) {
		http.Error(w, "El nombre es obligatorio y debe tener menos de 50 caracteres", http.StatusBadRequest)
		return
	}
	if !validators.IsValidText(usuario.Surname, 50) {
		http.Error(w, "El apellido es obligatorio y debe tener menos de 50 caracteres", http.StatusBadRequest)
		return
	}

	usuario.ID = primitive.NewObjectID()
	_, err := db.UsersCollection.InsertOne(context.Background(), usuario)
	if err != nil {
		http.Error(w, "Error al crear el usuario", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(usuario)
}

func UpdateUsuario(w http.ResponseWriter, r *http.Request, db *database.DB) {
	var usuario models.Usuario

	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		http.Error(w, "Datos inválidos: formato JSON incorrecto", http.StatusBadRequest)
		return
	}

	// Filtrar por el nombre del usuario
	filter := bson.M{"name": usuario.Name}

	// Verificar si el usuario existe
	count, err := db.UsersCollection.CountDocuments(context.Background(), filter)
	if err != nil {
		http.Error(w, "Error al verificar la existencia del usuario", http.StatusInternalServerError)
		return
	}

	if count == 0 {
		http.Error(w, "Usuario no existe", http.StatusNotFound)
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
		http.Error(w, "Error al actualizar usuario", http.StatusInternalServerError)
		return
	}

	if result.MatchedCount == 0 {
		http.Error(w, "No se pudo actualizar el usuario", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Usuario actualizado correctamente\n"))
}

func DeleteUsuario(w http.ResponseWriter, r *http.Request, db *database.DB) {
	name := r.URL.Query().Get("name")

	if name == "" {
		http.Error(w, "El nombre del usuario es obligatorio", http.StatusBadRequest)
		return
	}

	// Filtrar por el nombre del usuario
	filter := bson.M{"name": name}

	// Eliminar el usuario
	result, err := db.UsersCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		http.Error(w, "Error al eliminar usuario", http.StatusInternalServerError)
		return
	}

	if result.DeletedCount == 0 {
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Usuario eliminado correctamente\n"))
}
