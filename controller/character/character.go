package characterController

import "github.com/labstack/echo/v4"

func GetMyCharacter(ctx echo.Context) error {
	return ctx.JSON(200, "ok")
}

func GetCharacterById(ctx echo.Context) error {
	return ctx.JSON(200, "ok")
}

func CreateNewCharacter(ctx echo.Context) error {
	return ctx.JSON(200, "ok")
}

func EditCharacter(ctx echo.Context) error {
	return ctx.JSON(200, "ok")
}

func DeleteCharacter(ctx echo.Context) error {
	return ctx.JSON(200, "ok")
}
