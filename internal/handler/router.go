package handler

import (
	"CurlARC/internal/middleware"

	"github.com/labstack/echo"
)

func InitRouting(
	e *echo.Echo,
	userHandler UserHandler,
	teamHandler TeamHandler,
	recordHandler RecordHandler,
) {

	e.POST("/signup", userHandler.SignUp())
	e.POST("/signin", userHandler.SignIn())

	// デバッグ用
	debug := e.Group("/debug")
	debug.GET("/users", userHandler.GetAllUsers())
	debug.POST("/teams", teamHandler.CreateTeam())
	debug.GET("/teams", teamHandler.GetAllTeams())
	debug.GET("/teams/:teamId", teamHandler.GetMembers())
	debug.POST("/teams/:teamId/:targetId", teamHandler.InviteUser())
	debug.PATCH("/teams/:teamId/:userId", teamHandler.AcceptInvitation())
	debug.DELETE("/teams/:teamId/:userId", teamHandler.RemoveMember())

	// 認証が必要なルートにミドルウェアを適用
	authGroup := e.Group("/auth")

	// user集約
	authGroup.Use(middleware.JWTMiddleware)
	authGroup.GET("/me", userHandler.GetUser())
	authGroup.PATCH("/me", userHandler.UpdateUser())
	authGroup.DELETE("/me", userHandler.DeleteUser())

	// team集約
	authGroup.POST("/teams", teamHandler.CreateTeam())
	authGroup.GET("/teams", teamHandler.GetAllTeams())

	authGroup.GET("/teams/:teamId", teamHandler.GetMembers())
	authGroup.PATCH("/teams/:teamId", teamHandler.UpdateTeam())
	authGroup.DELETE("/teams/:teamId", teamHandler.DeleteTeam())

	authGroup.POST("/teams/:teamId/:userId", teamHandler.InviteUser())
	authGroup.PATCH("/teams/:teamId/:userId", teamHandler.AcceptInvitation())
	authGroup.DELETE("/teams/:teamId/:userId", teamHandler.RemoveMember())

	// record集約
	authGroup.POST("/record/:teamId/:userId", recordHandler.CreateRecord())
	authGroup.GET("/record/:teamId", recordHandler.GetRecordByTeamId())
	authGroup.PATCH("/record/:recordId/:userId", recordHandler.UpdateRecord())
	authGroup.DELETE("/record/:recordId", recordHandler.DeleteRecord())

	authGroup.PATCH("/record/:recordId/:userId", recordHandler.SetVisibility())
}
