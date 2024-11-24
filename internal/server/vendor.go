package server

import (
	"errors"
	"net/http"

	"github.com/hareem7bilal/go-microservice/internal/dberrors"
	"github.com/hareem7bilal/go-microservice/internal/models"
	"github.com/labstack/echo"
)

func (s *EchoServer) GetAllVendors(ctx echo.Context) error {
	vendors, err := s.DB.GetAllVendors(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, vendors)
}

func (s *EchoServer) AddVendor(ctx echo.Context) error {
	vendor := new(models.Vendor)
	if err := ctx.Bind(vendor); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	vendor, err := s.DB.AddVendor(ctx.Request().Context(), vendor)
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflictError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}

	}
	return ctx.JSON(http.StatusCreated, vendor)
}

func (s *EchoServer) GetVendorByID(ctx echo.Context) error {
	ID := ctx.Param("id")
	vendor, err := s.DB.GetVendorByID(ctx.Request().Context(), ID)
	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.JSON(http.StatusNotFound, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}

	}
	return ctx.JSON(http.StatusOK, vendor)
}

func (s *EchoServer) UpdateVendor(ctx echo.Context) error {
	ID := ctx.Param("id")
	vendor := new(models.Vendor)
	if err := ctx.Bind(vendor); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	if ID != vendor.VendorID {
		return ctx.JSON(http.StatusBadRequest, "id on path doesn't match id on body")
	}
	vendor, err := s.DB.UpdateVendor(ctx.Request().Context(), vendor)
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
	return ctx.JSON(http.StatusOK, vendor)
}

func (s *EchoServer) DeleteVendor(ctx echo.Context) error {
	ID := ctx.Param("id")
	err := s.DB.DeleteVendor(ctx.Request().Context(), ID)
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
