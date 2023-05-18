package controller

import (
	"net/http"
	"strconv"

	config "github.com/alfaa19/gin-restAPI-course/config/mysql"
	"github.com/alfaa19/gin-restAPI-course/helper"
	"github.com/alfaa19/gin-restAPI-course/model"
	"github.com/gin-gonic/gin"
)

func GetTotalUsers(c *gin.Context) {
	var users []model.User

	var count int64

	if err := config.DB.Where("role = ?", "user").Find(&users).Count(&count).Error; err != nil {
		helper.ResponseErrorJSON(c, http.StatusInternalServerError, err)
	}

	helper.ResponseSuccessJSON(c, "", gin.H{"total_users": count})
}

func DeleteUser(c *gin.Context) {
	param := c.Param("id")

	id, err := strconv.Atoi(param)

	if err != nil {
		helper.ResponseErrorJSON(c, http.StatusBadRequest, err)
	}

	var user model.User

	if err := config.DB.First(&user, id).Error; err != nil {
		helper.ResponseErrorJSON(c, http.StatusNotFound, err)
		return
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		helper.ResponseErrorJSON(c, http.StatusNotFound, err)
		return
	}

	helper.ResponseSuccessJSON(c, "", "")
}
