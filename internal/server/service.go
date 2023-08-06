package server

import (
	"net/http"

	"github.com/cwrenhold/linkedin-learning-build-a-microservice-with-go/internal/dberrors"
	"github.com/cwrenhold/linkedin-learning-build-a-microservice-with-go/internal/models"
	"github.com/labstack/echo/v4"
)

func (s *EchoServer) GetAllServices(ctx echo.Context) error {
	services, err := s.DB.GetAllServices(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, services)
}

func (s *EchoServer) GetServiceById(ctx echo.Context) error {
	ID := ctx.Param("id")

	service, err := s.DB.GetServiceById(ctx.Request().Context(), ID)

	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.JSON(http.StatusNotFound, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}

	return ctx.JSON(http.StatusOK, service)
}

func (s *EchoServer) AddService(ctx echo.Context) error {
	service := new(models.Service)

	if err := ctx.Bind(service); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}

	service, err := s.DB.AddService(ctx.Request().Context(), service)

	if err != nil {
		switch err.(type) {
		case *dberrors.ConflictError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}

	return ctx.JSON(http.StatusCreated, service)
}
