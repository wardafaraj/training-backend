package routes

import (
	"training/backend/webserver/controllers"

	"github.com/labstack/echo/v4"
)

func LookUpApiRouters(app *echo.Echo) {
	/*
		|--------------------------------------------------------------------------
		| Lookup Routes
		|--------------------------------------------------------------------------
	*/
	trainingtz := app.Group("/trainingtz/api/v1")
	absenteeismType := trainingtz.Group("/absenteeism-type")
	{
		absenteeismType.POST("/list", controllers.ListAbsenteeismType)
		absenteeismType.POST("/show", controllers.GetAbsenteeismType)
		absenteeismType.POST("/create", controllers.CreateAbsenteeismType)
		absenteeismType.POST("/update", controllers.UpdateAbsenteeismType)
		absenteeismType.POST("/delete", controllers.SoftDeleteAbsenteeismType)
		absenteeismType.POST("/destroy", controllers.DeleteAbsenteeismType)
	}
}