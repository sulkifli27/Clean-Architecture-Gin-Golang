package controller

import (
	"fmt"
	"github/com/sulkifli27/golang_api/helper"
	"github/com/sulkifli27/golang_api/model"
	"github/com/sulkifli27/golang_api/service"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type BookController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type bookController struct {
    bookService service.BookService
    jwtService service.JWTService
}

func NewBookController(bookServ service.BookService, jwtServ service.JWTService) BookController {
    return &bookController{
        bookService: bookServ,
        jwtService: jwtServ,
    }
}

func (c *bookController) All(context *gin.Context){
    var books []model.Book = c.bookService.All()
    res := helper.BuildResponse(true, "success", books)
    context.JSON(http.StatusOK, res)
}

func (c *bookController) FindByID(context *gin.Context){
    id, err := strconv.ParseUint(context.Param("id"), 0, 0)
    if err != nil {
        res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
        context.AbortWithStatusJSON(http.StatusBadRequest, res)
        return
    }

    var book model.Book = c.bookService.FindByID(id)
    if (book == model.Book{}){
        res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
        context.AbortWithStatusJSON(http.StatusBadRequest, res)
        return
    }else {
        res := helper.BuildResponse(true, "success", book)
        context.JSON(http.StatusOK, res)
    }
}

func (c *bookController) getUserIDByToken(token string) string {
    aToken, err := c.jwtService.ValidateToken(token)
    if err != nil {
        panic(err.Error())
    }
    claims := aToken.Claims.(jwt.MapClaims)
    return fmt.Sprintf("%v", claims["user_id"])
}
