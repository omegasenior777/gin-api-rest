package routes

import (
	"alura/gin-api-rest/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/alunos", controllers.ExibeAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriaAluno)
	r.GET("/alunos/:id", controllers.BuscaAluno)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PATCH("/alunos/:id", controllers.EditarAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoCPF)
	r.GET("/index", controllers.PaginaIndex)
	r.NoRoute(controllers.RotaNaoEncontrada)
	// define a PORTA r.Run(":5000")
	r.Run()
}
