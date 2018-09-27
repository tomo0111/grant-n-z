package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/tomoyane/grant-n-z/di"
	"github.com/tomoyane/grant-n-z/infra"
	"github.com/tomoyane/grant-n-z/handler"
)

func PostRole(c echo.Context) (err error) {
	token := c.Request().Header.Get("Authorization")
	errAuth := di.ProviderTokenService.VerifyToken(c, token)

	if errAuth != nil {
		return echo.NewHTTPError(errAuth.Code, errAuth)
	}

	role := new(entity.Role)
	if err := c.Bind(role); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, handler.BadRequest(""))
	}

	if err := c.Validate(role); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, handler.BadRequest(""))
	}

	roleData, errRole := di.ProviderRoleService.PostRoleData(role, token)
	if errRole != nil {
		return echo.NewHTTPError(errRole.Code, errRole)
	}

	c.Response().Header().Add("Location", infra.GetHostName() + "/v1/roles/" + role.Uuid.String())
	return c.JSON(http.StatusCreated, roleData)
}