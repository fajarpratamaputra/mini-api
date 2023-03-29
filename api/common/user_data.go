package common

import (
	"github.com/labstack/echo/v4"
)

func GetUserData(c echo.Context) map[string]interface{} {
	return c.Get("user_data").(map[string]interface{})
}
