package controller

import (
	"fmt"
	"github/com/sulkifli27/golang_api/dto"
	"github/com/sulkifli27/golang_api/helper"
	"github/com/sulkifli27/golang_api/service"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	Update(context *gin.Context)
	Profile(context *gin.Context)
}

type userController struct {
    userService service.UserService
    jwtService  service.JWTService
}

func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
    return &userController{
        userService: userService,
        jwtService: jwtService,
    }
}

func (c *userController) Update(context *gin.Context) {
    var userUpdateDTO dto.UserUpdateDTO
    errDTO := context.ShouldBindJSON(&userUpdateDTO)
    if errDTO != nil {
        res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
        context.AbortWithStatusJSON(http.StatusBadRequest, res)
        return
    }

    authHeader := context.GetHeader("Authorization")
    token, errToken := c.jwtService.ValidateToken(authHeader)
    if errToken != nil {
        panic(errToken.Error())
    }

    claims := token.Claims.(jwt.MapClaims)
    id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10 ,64)
    if err != nil {
        panic(err.Error())
    }
    userUpdateDTO.ID = id
    user := c.userService.Update(userUpdateDTO)
    res := helper.BuildResponse(true, "success",  user)
    context.JSON(http.StatusOK,res)
}

func (c *userController) Profile(context * gin.Context) {
    authHeader := context.GetHeader("Authorization")
    token, errToken := c.jwtService.ValidateToken(authHeader)
    if errToken != nil {
        panic(errToken.Error())
    }

    claims := token.Claims.(jwt.MapClaims)
    user := c.userService.Profile(fmt.Sprintf("%v", claims["user_id"]))
    res := helper.BuildResponse(true, "success", user)
    context.JSON(http.StatusOK , res)
}


