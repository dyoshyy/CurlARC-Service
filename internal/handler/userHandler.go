package handler

import (
	"CurlARC/internal/handler/request"
	"CurlARC/internal/handler/response"
	"CurlARC/internal/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

// UserHandler handles requests related to users.
type UserHandler struct {
	userUsecase usecase.UserUsecase
}

// NewUserHandler creates a new UserHandler instance.
func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return UserHandler{userUsecase: userUsecase}
}

// SignUp handles user registration.
// @Summary Register a new user
// @Description Registers a new user with the provided ID token, name, and email
// @Tags Users
// @Accept json
// @Produce json
// @Param user body request.SignUpRequest true "User registration information"
// @Success 201 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Failure 409 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /signup [post]
func (h *UserHandler) SignUp() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req request.SignUpRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusBadRequest,
					Message: "invalid request",
				},
			})
		}

		signUpedUser, err := h.userUsecase.SignUp(c.Request().Context(), req.IdToken, req.Name, req.Email)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			})
		}

		return c.JSON(http.StatusCreated, response.SuccessResponse{
			Status: "success",
			Data: response.User{
				Id:    signUpedUser.GetId().Value(),
				Name:  signUpedUser.GetName(),
				Email: signUpedUser.GetEmail(),
			},
		})
	}
}

// SignIn handles user login.
// @Summary Log in a user
// @Description Logs in a user with the provided ID token and returns a JWT
// @Tags Users
// @Accept json
// @Produce json
// @Param user body request.SignInRequest true "User login information"
// @Success 200 {object} response.SuccessResponse{data=response.SignInResponse}
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /signin [post]
func (h *UserHandler) SignIn() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req request.SignInRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusBadRequest,
					Message: "invalid request",
				},
			})
		}

		user, cookie, err := h.userUsecase.SignIn(c.Request().Context(), req.IdToken)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			})
		}

		// Set the JWT token as a cookie
		c.SetCookie(cookie) // jwt

		return c.JSON(http.StatusOK, response.SuccessResponse{
			Status: "success",
			Data:   user,
		})
	}
}

// GetAllUsers retrieves all users.
// @Summary Get all users
// @Description Retrieves a list of all registered users
// @Tags Users
// @Produce json
// @Success 200 {object} response.SuccessResponse{data=[]model.User}
// @Failure 500 {object} response.ErrorResponse
// @Router /users [get]
func (h *UserHandler) GetAllUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := h.userUsecase.GetAllUsers(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			})
		}

		return c.JSON(http.StatusOK, response.SuccessResponse{
			Status: "success",
			Data:   users,
		})
	}
}

// GetUser retrieves information about a specific user.
// @Summary Get user information
// @Description Retrieves information about the currently authenticated user
// @Tags Users
// @Produce json
// @Success 200 {object} response.SuccessResponse{data=response.User}
// @Failure 500 {object} response.ErrorResponse
// @Router /users/me [get]
func (h *UserHandler) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Get("uid").(string)

		user, err := h.userUsecase.GetUser(c.Request().Context(), id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			})
		}

		res := response.User{
			Id:    user.GetId().Value(),
			Name:  user.GetName(),
			Email: user.GetEmail(),
		}

		return c.JSON(http.StatusOK, response.SuccessResponse{
			Status: "success",
			Data:   res,
		})
	}
}

// UpdateUser updates user information.
// @Summary Update user information
// @Description Updates the name and email of the currently authenticated user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body request.UpdateUserRequest true "Updated user information"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /users [PATCH]
func (h *UserHandler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req request.UpdateUserRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusBadRequest,
					Message: "invalid request",
				},
			})
		}
		id := c.Get("uid").(string)

		if _, err := h.userUsecase.UpdateUser(c.Request().Context(), id, req.Name, req.Email); err != nil {
			return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			})
		}
		return c.JSON(http.StatusOK, response.SuccessResponse{
			Status: "success",
			Data:   nil,
		})
	}
}

// DeleteUser deletes a specific user.
// @Summary Delete a user
// @Description Deletes a user with the provided ID
// @Tags Users
// @Accept json
// @Produce json
// @Param user body request.DeleteUserRequest true "User ID to delete"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /users [delete]
func (h *UserHandler) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req request.DeleteUserRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusBadRequest,
					Message: "invalid request",
				},
			})
		}

		if err := h.userUsecase.DeleteUser(c.Request().Context(), req.Id); err != nil {
			return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
				Status: "error",
				Error: response.ErrorDetail{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			})
		}
		return c.JSON(http.StatusOK, response.SuccessResponse{
			Status: "success",
			Data:   nil,
		})
	}
}
