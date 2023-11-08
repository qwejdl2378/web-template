package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qwejdl2378/web-template/httputil"
	"github.com/qwejdl2378/web-template/service"
	"github.com/qwejdl2378/web-template/service/dto"
	"github.com/qwejdl2378/web-template/store"
	"net/http"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	db := store.GetDBInstance()
	userService := service.NewUserService(db)
	return &UserController{userService: userService}
}

func (u *UserController) CreateUser(ctx *gin.Context) {
	var userDTO dto.UserDTO
	err := ctx.ShouldBindJSON(&userDTO)
	if err != nil {
		httputil.ThrowError(ctx, http.StatusBadRequest, fmt.Errorf("params error"))
		return
	}
	err = u.userService.CreateUser(&userDTO)
	if err != nil {
		httputil.ThrowError(ctx, http.StatusInternalServerError, fmt.Errorf("params error"))
		return
	}
	httputil.ReturnSuccessWithMsg(ctx, "create user success", nil)
	return
}
