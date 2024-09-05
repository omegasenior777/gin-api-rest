package database

import (
	"alura/gin-api-rest/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaBanco() {
	stringConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringConexao))
	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados")
	}
	DB.AutoMigrate(&models.Aluno{})
}
