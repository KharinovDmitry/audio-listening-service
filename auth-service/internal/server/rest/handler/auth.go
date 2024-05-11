package handler

import (
	"auth-service/internal/domain/service"
	"auth-service/internal/server/rest/dto"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	authService service.AuthService
	logger      service.LoggerService
}

func NewAuthHandler(authService service.AuthService, logger service.LoggerService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		logger:      logger,
	}
}

// SignUp godoc
// @Tags auth
// @Summary Регистрация аккаунта
// @Accept json
// @Produce json
// @Param input body dto.SignUpRequest true "Sign-up"
// @Success 200
// @Failure 404
// @Failure 500
// @Router /api/sign-up [post]
func (a *AuthHandler) SignUp(c *gin.Context) {
	var req dto.SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := a.authService.SignUp(c.Request.Context(), req.Login, req.Password, req.Role)
	if err != nil {
		if errors.Is(err, service.ErrUserAlreadyExists) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		a.logger.Error("In AuthHandler(SignUp): %s", err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

// SignIn godoc
// @Tags auth
// @Summary Регистрация аккаунта
// @Accept json
// @Produce json
// @Param input body dto.SignInRequest true "Sign-in"
// @Success 200 {object} dto.SignInResponse
// @Failure 404
// @Failure 500
// @Router /api/sign-in [post]
func (a *AuthHandler) SignIn(c *gin.Context) {
	var req dto.SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := a.authService.SignIn(c.Request.Context(), req.Login, req.Password)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		a.logger.Error("In AuthHandler(SignIn): %s", err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, dto.SignInResponse{Token: token})
}
