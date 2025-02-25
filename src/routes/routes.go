package routes

import (
	"josu-foruria/src/database"
	"josu-foruria/src/handlers"

	"github.com/gin-gonic/gin"
)

func HandleUsuarios(r *gin.Engine, db *database.DB) {
	handlers.InitUserHandlers(r, db)
}
