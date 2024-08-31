package injector

import (
	"CurlARC/internal/domain/repository"
	"CurlARC/internal/handler"
	"CurlARC/internal/infra"
	"CurlARC/internal/usecase"
)

// UserRepository (interface) に実装である SqlHandler を渡し生成する

func InjectUserRepository() repository.UserRepository {
	sqlHandler := InjectDB()
	return infra.NewUserRepository(sqlHandler)
}

func InjectUserUsecase() usecase.UserUsecase {
	userRepo := InjectUserRepository()
	authClient := infra.NewFireBaseAuthClient()
	return usecase.NewUserUsecase(userRepo, authClient)
}

func InjectUserHandler() handler.UserHandler {
	userUsecase := InjectUserUsecase()
	return handler.NewUserHandler(userUsecase)
}
