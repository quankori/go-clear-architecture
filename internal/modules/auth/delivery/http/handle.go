package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/quankori/go-clear-architecture/internal/modules/auth"
	"github.com/quankori/go-clear-architecture/internal/modules/auth/dto"
	"github.com/quankori/go-clear-architecture/utils"
)

type authHandler struct {
	useCase auth.UseService
}

func NewAuthHandler(useCase auth.UseService) auth.Handler {
	return &authHandler{
		useCase: useCase,
	}
}

func (h *authHandler) SignUp() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := &dto.SignUpInput{}
		if err := utils.ReadRequest(c, input); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		user, err := h.useCase.SignUp(c.Request().Context(), input.Username, input.Password)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, dto.SignUpResponse{Id: user.ID, Username: user.Username})
	}
}

func (h *authHandler) SignIn() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := &dto.LoginInput{}
		if err := utils.ReadRequest(c, input); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		token, err := h.useCase.SignIn(c.Request().Context(), input.Username, input.Password)
		if err != nil {
			if err == auth.ErrUserNotFound {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}
			if err == auth.ErrWrongPassword {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, dto.LogInResponse{Token: token})
	}
}
