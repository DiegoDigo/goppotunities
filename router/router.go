package router

import "github.com/gin-gonic/gin"

func Initialize() {
	r := gin.Default()

	initializerRoutes(r)

	if err := r.Run("0.0.0.0:5000"); err != nil {
		panic("Error ao subir o servidor")
	}
}
