package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model
	Nome string `json:"nome" validate:"required"`
	CPF  string `json:"cpf" validate:"required,len=11,numeric"`
	RG   string `json:"rg" validate:"required,len=8,numeric"`
}

var validate *validator.Validate

func ValidadaDados(aluno *Aluno) error {
	validate = validator.New()
	if err := validate.Struct(aluno); err != nil {
		return err
	}
	return nil
}
