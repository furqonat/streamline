package test

import (
	user_v1 "apps/auth/controllers/v1/user"
	"apps/auth/dto"
	"apps/auth/middlewares"
	"apps/auth/services"
	"apps/auth/utils"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alecthomas/assert"
)

func TestGetUserFail(t *testing.T) {
	r := SetUpRouter()
	userController := user_v1.NewUserController(utils.GetLogger(), services.NewUserService(utils.GetLogger(), utils.NewDatabase(utils.GetLogger()), utils.NewPwHash()))
	jwtMiddleware := middlewares.NewJwtMiddleware(utils.NewJwt(utils.GetLogger(), utils.NewEnv(utils.GetLogger())), utils.GetLogger())
	api := r.Group("/api").Use(jwtMiddleware.HandleAuthWithRoles(utils.ROLE_USER))
	api.GET("/user", userController.GetUser)

	req, _ := http.NewRequest("GET", "/api/user", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	t.Log(w.Body.String())
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestGetUserOk(t *testing.T) {
	r := SetUpRouter()
	userController := user_v1.NewUserController(utils.GetLogger(), services.NewUserService(utils.GetLogger(), utils.NewDatabase(utils.GetLogger()), utils.NewPwHash()))
	jwtMiddleware := middlewares.NewJwtMiddleware(utils.NewJwt(utils.GetLogger(), utils.NewEnv(utils.GetLogger())), utils.GetLogger())
	api := r.Group("/api").Use(jwtMiddleware.HandleAuthWithRoles(utils.ROLE_USER))
	api.GET("/user", userController.GetUser)
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTU3NTgwMTgsImlkIjoiY2x3N2d0c3A4MDAwMDgwdXhoZmwycnd0biIsIm5hbWUiOiJhZG1pbiIsInJvbGUiOlt7ImlkIjoiY2x3N2dzcjVmMDAwMHpybXl0cnRmZThwciIsIm5hbWUiOiJ1c2VyIn1dLCJ0b2tlbiI6IiIsInJlZnJlc2hfdG9rZW4iOiIifQ.BMXLQo3xzx_MkIUNP6gTAQDxtcz9Nqf0hOMn-OIutzw"
	req, _ := http.NewRequest("GET", "/api/user", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	t.Log(w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateUserName(t *testing.T) {
	r := SetUpRouter()
	userController := user_v1.NewUserController(utils.GetLogger(), services.NewUserService(utils.GetLogger(), utils.NewDatabase(utils.GetLogger()), utils.NewPwHash()))
	jwtMiddleware := middlewares.NewJwtMiddleware(utils.NewJwt(utils.GetLogger(), utils.NewEnv(utils.GetLogger())), utils.GetLogger())
	api := r.Group("/api").Use(jwtMiddleware.HandleAuthWithRoles(utils.ROLE_USER))
	api.PUT("/user", userController.UpdateUser)
	token := "" // TODO add token
	username := "user1aa"
	body := dto.UpdateUserDto{
		Username: &username,
	}
	data, _ := json.Marshal(body)
	req, _ := http.NewRequest("PUT", "/api/user", bytes.NewBuffer(data))
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	t.Log(w.Body.String())
	assert.Equal(t, http.StatusUnauthorized, w.Code)

}

func TestUpdateName(t *testing.T) {
	r := SetUpRouter()
	userController := user_v1.NewUserController(utils.GetLogger(), services.NewUserService(utils.GetLogger(), utils.NewDatabase(utils.GetLogger()), utils.NewPwHash()))
	jwtMiddleware := middlewares.NewJwtMiddleware(utils.NewJwt(utils.GetLogger(), utils.NewEnv(utils.GetLogger())), utils.GetLogger())
	api := r.Group("/api").Use(jwtMiddleware.HandleAuthWithRoles(utils.ROLE_USER))
	api.PUT("/user", userController.UpdateUser)
	token := ""
	name := "Dabidu"
	body := dto.UpdateUserDto{
		Name: &name,
	}
	data, _ := json.Marshal(body)
	req, _ := http.NewRequest("PUT", "/api/user", bytes.NewBuffer(data))
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	t.Log(w.Body.String())
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
