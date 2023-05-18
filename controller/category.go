package controller

import (
	"net/http"

	config "github.com/alfaa19/gin-restAPI-course/config/mysql"
	"github.com/alfaa19/gin-restAPI-course/helper"
	"github.com/alfaa19/gin-restAPI-course/model"
	"github.com/alfaa19/gin-restAPI-course/request"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

type CategoryEnroll struct {
	CategoryID   uint   `gorm:"column:category_id" json:"category_id"`
	CategoryName string `gorm:"column:name" json:"name"`
	TotalEnroll  uint   `gorm:"column:total_enroll" json:"total_enroll"`
}

func GetCategory(c *gin.Context) {
	var categories []model.Category

	if err := config.DB.Find(&categories).Error; err != nil {
		helper.ResponseErrorJSON(c, http.StatusInternalServerError, err)
	}

	helper.ResponseSuccessJSON(c, "", categories)
}

func GetPopularCategory(c *gin.Context) {
	var categoryEnrolls []CategoryEnroll

	if err := config.DB.Table("courses").
		Select("category_id,categories.name as name, SUM(enroll) as total_enroll").
		Joins("JOIN categories ON categories.id = courses.category_id").
		Group("category_id").
		Scan(&categoryEnrolls).
		Order("total_enroll DESC").
		Limit(5).
		Error; err != nil {
		helper.ResponseErrorJSON(c, http.StatusInternalServerError, err)
	}

	helper.ResponseSuccessJSON(c, "", categoryEnrolls)
}

func CreateCategory(c *gin.Context) {
	var req request.Category

	if err := c.ShouldBind(&req); err != nil {
		helper.ResponseErrorJSON(c, http.StatusBadRequest, err)
	}

	if err := req.Validate(); err != nil {
		helper.ResponseValidationErrorJson(c, "Error Validation", err.(validation.Errors))
	}

	var category = model.Category{
		Name: req.Name,
	}
	if err := config.DB.Create(&category).Error; err != nil {
		helper.ResponseErrorJSON(c, http.StatusUnprocessableEntity, err)
		return
	}

	helper.ResponseSuccessJSON(c, "", category)
}
