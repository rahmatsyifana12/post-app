package register

import (
	"net/http"
	"post-app/entities"

	"github.com/labstack/echo/v4"
)

type RegisterController struct {}

func (rc* RegisterController) CreateUser(c echo.Context) (err error) {
	u := new(entities.User)
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	rs := RegisterService{}
	rs.CreateUser(u)
	return err
}
