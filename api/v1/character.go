package v1

import (
	characterController "github.com/ionnotion/statictools-server/controller/character"
	"github.com/labstack/echo/v4"
)

func InitCharacterRoute(e *echo.Echo) {
	characterGroup := e.Group("/character")
	characterGroup.POST("", characterController.CreateNewCharacter)

	characterGroup.GET("/my-characters", characterController.GetMyCharacter)
	characterGroup.GET("/:id", characterController.GetCharacterById)

	characterGroup.PUT("/:id", characterController.CreateNewCharacter)

	characterGroup.DELETE("/:id", characterController.DeleteCharacter)
}
