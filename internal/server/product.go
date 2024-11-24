package server

import (
	"errors"
	"net/http"

	"github.com/hareem7bilal/go-microservice/internal/dberrors"
	"github.com/hareem7bilal/go-microservice/internal/models"
	"github.com/labstack/echo"
)

func (s *EchoServer) GetAllProducts(ctx echo.Context) error {
	ProductId := ctx.QueryParam("ProductId")
	products, err := s.DB.GetAllProducts(ctx.Request().Context(), ProductId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, products)
}

func (s *EchoServer) AddProduct(ctx echo.Context) error {
	product := new(models.Product)
	if err := ctx.Bind(product); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	product, err := s.DB.AddProduct(ctx.Request().Context(), product)
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflictError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}

	}
	return ctx.JSON(http.StatusCreated, product)
}

func (s *EchoServer) GetProductByID(ctx echo.Context) error {
	ID := ctx.Param("id")
	product, err := s.DB.GetProductByID(ctx.Request().Context(), ID)
	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.JSON(http.StatusNotFound, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}

	}
	return ctx.JSON(http.StatusOK, product)
}

func (s *EchoServer) UpdateProduct(ctx echo.Context) error {
	ID := ctx.Param("id")
	product := new(models.Product)
	if err := ctx.Bind(product); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	if ID != product.ProductID {
		return ctx.JSON(http.StatusBadRequest, "id on path doesn't match id on body")
	}
	product, err := s.DB.UpdateProduct(ctx.Request().Context(), product)
	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.JSON(http.StatusNotFound, err)
		case *dberrors.ConflictError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}

	}
	return ctx.JSON(http.StatusOK, product)
}

func (s *EchoServer) DeleteProduct(ctx echo.Context) error {
	ID := ctx.Param("id")
	err := s.DB.DeleteProduct(ctx.Request().Context(), ID)
	if err != nil {
        switch {
        case errors.Is(err, &dberrors.NotFoundError{}):
            return ctx.JSON(http.StatusNotFound, err.Error())
        default:
            return ctx.JSON(http.StatusInternalServerError, err.Error())
        }
    }
	return ctx.NoContent(http.StatusResetContent)
}
