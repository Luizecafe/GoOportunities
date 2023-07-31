package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// inicializa o Router utilizando as configurações Default do gin
	router := gin.Default()
	// Definindo uma Rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Estamos rodando nossa API na porta 8080 (padrão).
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
