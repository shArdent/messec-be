package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInterface interface {
    GetAll(c *gin.Context)
}

type UserController struct {
	DB any
}

func NewPersonController(db any) *UserController {
	return &UserController{DB: db}
}

func (pc *UserController) GetAll (c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "status" : "ok",
        "message" : "hello world",
    })
}

