package controller

import (
	"github/com/sulkifli27/golang_api/dto"
	"github/com/sulkifli27/golang_api/helper"
	"github/com/sulkifli27/golang_api/model"
	"github/com/sulkifli27/golang_api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
    authService service.AuthService
    jwtService  service.JWTService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController{
    return &authController{
        authService: authService,
        jwtService: jwtService,
    }
}

func (c *authController) Login(ctx *gin.Context){
    var loginDTO dto.LoginDTO
    errDTO := ctx.ShouldBind(&loginDTO)
    if errDTO != nil {
        response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
        ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
        return
    }
    authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
    if value, ok := authResult.(model.User) ; ok {
        generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(value.ID, 10))
        value.Token = generatedToken
        response := helper.BuildResponse(true, "success", value)
        ctx.JSON(http.StatusOK, response)
        return
    }
    response := helper.BuildErrorResponse("Please check again your credential", "Invalid Credential", helper.EmptyObj{} )
    ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authController) Register(ctx *gin.Context){
    var registerDTO dto.RegisterDTO
    errDTO := ctx.ShouldBind(&registerDTO)
    if errDTO != nil {
        response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
        ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
        return
    }
    if  !c.authService.IsDuplicateEmail(registerDTO.Email) {
        response := helper.BuildErrorResponse("Failed to process request", "duplicat email", helper.EmptyObj{})
        ctx.JSON(http.StatusConflict, response)
    }else{
        createdUser := c.authService.CreateUser(registerDTO)
        token := c.jwtService.GenerateToken(strconv.FormatUint(createdUser.ID,10))
        createdUser.Token = token
        response := helper.BuildResponse(true, "success", createdUser)
        ctx.JSON(http.StatusCreated, response)
    }
}