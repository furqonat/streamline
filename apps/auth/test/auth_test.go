package test

import (
	auth_v1 "apps/auth/controllers/v1/auth"
	"apps/auth/dto"
	"apps/auth/services"
	"apps/auth/utils"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestAuthSignIn(t *testing.T) {
	r := SetUpRouter()
	authController := auth_v1.NewAuthController(utils.GetLogger(), services.NewAuthService(utils.NewJwt(utils.GetLogger(), utils.NewEnv(utils.GetLogger())), utils.GetLogger(), utils.NewDatabase(utils.GetLogger()), utils.NewPwHash()))
	r.POST("/signin", authController.SignIn)
	body := dto.SignInDto{
		Password: "admin",
		Username: "admin",
	}
	data, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/signin", bytes.NewBuffer(data))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestAuthSignUp(t *testing.T) {
	r := SetUpRouter()
	authController := auth_v1.NewAuthController(utils.GetLogger(), services.NewAuthService(utils.NewJwt(utils.GetLogger(), utils.NewEnv(utils.GetLogger())), utils.GetLogger(), utils.NewDatabase(utils.GetLogger()), utils.NewPwHash()))
	r.POST("/signup", authController.SignUp)
	ipAddress := "127.0.0.1"
	userAgent := "admin"
	device := "admin"
	body := dto.SignUpDto{
		Email: "admin",
		Name:  "admin",
		SignInDto: dto.SignInDto{
			Username:  "admin",
			Password:  "admin",
			IpAddress: &ipAddress,
			UserAgent: &userAgent,
			Device:    &device,
		},
	}
	data, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(data))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
