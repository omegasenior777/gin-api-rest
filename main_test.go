package main

import (
	"alura/gin-api-rest/controllers"
	"alura/gin-api-rest/database"
	"alura/gin-api-rest/models"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupRotasTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}
func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "Nome teste 0",
		RG:  "12345600",
		CPF: "12345678900"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}
func TestStatusCodeSaudacao(t *testing.T) {
	r := SetupRotasTeste()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/gui", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
	mockResposta := `{"API diz:":"E ai gui?"}`
	respostaBody, err := io.ReadAll(resposta.Body)
	if err != nil {
		t.Fatalf("Erro ao ler o corpo da resposta: %v", err)
	}

	assert.Equal(t, mockResposta, string(respostaBody))
}

func TestListaTodosAlunosHandler(t *testing.T) {
	database.ConectaBanco()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupRotasTeste()
	r.GET("/alunos", controllers.ExibeAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscaPorCPFHandler(t *testing.T) {
	database.ConectaBanco()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupRotasTeste()
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoCPF)
	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678900", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscaPorIDHandler(t *testing.T) {
	database.ConectaBanco()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupRotasTeste()
	r.GET("/alunos/:id", controllers.BuscaAluno)
	path := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", path, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var alunoMock models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMock)
	assert.Equal(t, "Nome teste 0", alunoMock.Nome)
	assert.Equal(t, "12345678900", alunoMock.CPF)
	assert.Equal(t, "12345600", alunoMock.RG)
}

func TestDeletaAlunoHandler(t *testing.T) {
	database.ConectaBanco()
	CriaAlunoMock()
	r := SetupRotasTeste()
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	path := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", path, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestEditaAlunoHandler(t *testing.T) {
	database.ConectaBanco()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupRotasTeste()
	r.PATCH("/alunos/:id", controllers.EditarAluno)
	aluno := models.Aluno{Nome: "Nome teste 000",
		RG:  "99999999",
		CPF: "12345678900"}
	valorJson, _ := json.Marshal(aluno)
	path := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", path, bytes.NewBuffer(valorJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var alunoMock models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMock)
	assert.Equal(t, "Nome teste 000", alunoMock.Nome)
	assert.Equal(t, "12345678900", alunoMock.CPF)
	assert.Equal(t, "99999999", alunoMock.RG)
	assert.Equal(t, http.StatusOK, resposta.Code)
}
