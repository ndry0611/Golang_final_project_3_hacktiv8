package handler

import (
	"final_project_3/dto"
	"final_project_3/pkg/errs"
	"final_project_3/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) userHandler {
	return userHandler{userService: userService}
}

func (uh *userHandler) CreateUser(ctx *gin.Context) {
	var userPayload dto.NewUserRequest
	if err := ctx.ShouldBindJSON(&userPayload); err != nil {
		errBind := errs.NewUnprocessableEntityResponse("invalid request body")
		ctx.JSON(errBind.Status(), errBind)
		return
	}
	response, err := uh.userService.CreateUser(&userPayload)
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(http.StatusCreated, response)
}

func (uh *userHandler) Login(ctx *gin.Context) {
	var loginPayload dto.NewLoginRequest

	if err := ctx.ShouldBindJSON(&loginPayload); err != nil {
		errBindJson := errs.NewUnprocessableEntityResponse("invalid json request body")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	resp, err := uh.userService.Login(&loginPayload)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}