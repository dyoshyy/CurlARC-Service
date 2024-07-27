package usecase

import (
	"CurlARC/internal/domain/model"
	"CurlARC/internal/domain/repository"
	"context"
	"errors"

	"firebase.google.com/go/v4/auth"
	"gorm.io/gorm"
)

type UserUsecase interface {
	// CRUD
	SignUp(ctx context.Context, idToken, name, email string) error
	GetUser(ctx context.Context, id string) (*model.User, error)
	UpdateUser(ctx context.Context, id, name, email string) error
	DeleteUser(ctx context.Context, id string) error

	AuthUser(ctx context.Context, id_token string) (*model.User, error)
	GetAllUsers(ctx context.Context) ([]*model.User, error)

	// Team関連
	// AcceptTeamInvitation(ctx context.Context, userID, teamID string) error
	// RejectTeamInvitation(ctx context.Context, userID, teamID string) error
}

type userUsecase struct {
	userRepo   repository.UserRepository
	authClient *auth.Client
}

func NewUserUsecase(userRepo repository.UserRepository, authCli *auth.Client) UserUsecase {
	return &userUsecase{userRepo: userRepo, authClient: authCli}
}

func (usecase *userUsecase) SignUp(ctx context.Context, idToken, name, email string) (err error) {
	// idTokenを検証
	token, err := usecase.authClient.VerifyIDToken(ctx, idToken)
	if err != nil {
		return repository.ErrUnauthorized
	}

	// ユーザーをdbに保存
	user := &model.User{
		Id:    token.UID,
		Name:  name,
		Email: email,
	}

	if _, err := usecase.userRepo.Save(user); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return repository.ErrEmailExists
		}
		usecase.authClient.DeleteUser(ctx, token.UID) // dbへの保存が失敗したらfirebase上のユーザーも削除
		return err
	}

	return nil
}

func (usecase *userUsecase) AuthUser(ctx context.Context, id_token string) (*model.User, error) {
	// Verify the ID token
	authToken, err := usecase.authClient.VerifyIDToken(context.Background(), id_token)
	if err != nil {
		return nil, repository.ErrUnauthorized
	}

	// Find the user by UID
	user, err := usecase.userRepo.FindById(authToken.UID)
	if err != nil {
		return nil, repository.ErrUserNotFound
	}

	return user, nil
}

func (usecase *userUsecase) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	return usecase.userRepo.FindAll()
}

func (usecase *userUsecase) GetUser(ctx context.Context, id string) (*model.User, error) {
	return usecase.userRepo.FindById(id)
}

func (usecase *userUsecase) UpdateUser(ctx context.Context, id, name, email string) error {

	// Firebase上のユーザー情報を更新
	params := (&auth.UserToUpdate{}).
		Email(email).
		DisplayName(name)

	_, err := usecase.authClient.UpdateUser(ctx, id, params)
	if err != nil {
		return err
	}

	// ユーザーをdbに保存
	user := &model.User{
		Id:    id,
		Name:  name,
		Email: email,
	}

	return usecase.userRepo.Update(user)
}

func (usecase *userUsecase) DeleteUser(ctx context.Context, id string) error {
	// Firebase上のユーザー情報を削除
	err := usecase.authClient.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return usecase.userRepo.Delete(id)
}
