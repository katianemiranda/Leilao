package user_controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	rest_err "github.com/katianemiranda/leilao/configuration/rest_err"
	"github.com/katianemiranda/leilao/internal/usecase/user_usecase"
)

type userController struct {
	userUseCase user_usecase.UseruseCase
}

func NewUserController(userUseCase user_usecase.UseruseCase) *userController {
	return &userController{
		userUseCase: userUseCase,
	}
}

func (u *userController) FindUserById(c *gin.Context) {
	userId := c.Param("userId")

	//localhost:8080/user?userId=123456789

	if err := uuid.Validate(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid fields", rest_err.Causes{
			Field:   "userId",
			Message: "userId must be a valid UUID",
		})
		c.JSON(errRest.Code, errRest)
		return
	}

	userData, err := u.userUseCase.FindUserById(context.Background(), userId)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, userData)

}
