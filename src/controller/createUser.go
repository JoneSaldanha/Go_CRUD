package controller

import (
	"fmt"
	"log"

	"github.com/JoneSaldanha/Go_CRUD/src/configuration/rest_err"
	"github.com/JoneSaldanha/Go_CRUD/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal abject, error=%s", err.Error())
		errRest := rest_err.NewBadRequestError("Some fields are incorrect")

		c.JSON(errRest.Code, errRest)
		return
	}
	fmt.Println(userRequest)

}
