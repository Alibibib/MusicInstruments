package tests

import (
	"MusicInstruments/database"
	"MusicInstruments/handlers"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	err := database.InitDB()
	if err != nil {
		log.Fatalf("Ошибка инициализации БД: %v", err)
	}
	os.Exit(m.Run())
}

func TestRegisterHandler(t *testing.T) {
	router := gin.Default()
	router.POST("/register", handlers.RegisterHandler)

	payload := `{"username": "testuser", "password": "testpass"}`

	req, _ := http.NewRequest("POST", "/register", bytes.NewBufferString(payload))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
}

func TestLoginHandler(t *testing.T) {
	router := gin.Default()
	router.POST("/login", handlers.LoginHandler)

	payload := `{"username": "testuser", "password": "testpass"}`

	req, _ := http.NewRequest("POST", "/login", bytes.NewBufferString(payload))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "token")
}

func TestGetAllUsersHandler_Unauthorized(t *testing.T) {
	router := gin.Default()
	router.GET("/users", handlers.GetAllUsersHandler)

	req, _ := http.NewRequest("GET", "/users", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnauthorized, resp.Code)
}

func TestGetUserByIDHandler(t *testing.T) {
	router := gin.Default()
	router.GET("/users/:id", handlers.GetUserByIDHandler)

	req, _ := http.NewRequest("GET", "/users/1", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestCreateInstrumentHandler(t *testing.T) {
	router := gin.Default()
	handler := handlers.NewMusicalInstrumentHandler()
	router.POST("/instruments", handler.CreateHandler)

	instrument := map[string]string{
		"name":     "Guitar",
		"category": "Strings",
		"brand":    "Yamaha",
	}
	body, _ := json.Marshal(instrument)

	req, _ := http.NewRequest("POST", "/instruments", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
}

func TestGetAllInstrumentsHandler(t *testing.T) {
	router := gin.Default()
	handler := handlers.NewMusicalInstrumentHandler()
	router.GET("/instruments", handler.GetAllHandler)

	req, _ := http.NewRequest("GET", "/instruments", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetInstrumentByIDHandler(t *testing.T) {
	router := gin.Default()
	handler := handlers.NewMusicalInstrumentHandler()
	router.GET("/instruments/:id", handler.GetByIDHandler)

	req, _ := http.NewRequest("GET", "/instruments/1", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestDeleteInstrumentHandler(t *testing.T) {
	router := gin.Default()
	handler := handlers.NewMusicalInstrumentHandler()
	router.DELETE("/instruments/:id", handler.DeleteHandler)

	req, _ := http.NewRequest("DELETE", "/instruments/1", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestUpdateUserHandler(t *testing.T) {
	router := gin.Default()
	router.PUT("/users/:id", handlers.UpdateUserHandler)

	payload := `{"username": "updateduser"}`

	req, _ := http.NewRequest("PUT", "/users/1", bytes.NewBufferString(payload))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetAllCategoriesHandler(t *testing.T) {
	router := gin.Default()
	handler := handlers.NewCategoryHandler()
	router.GET("/categories", handler.GetAllHandler)

	req, _ := http.NewRequest("GET", "/categories", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}
