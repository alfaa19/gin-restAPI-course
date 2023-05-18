package controller

import (
	"context"
	"net/http"
	"strconv"

	config "github.com/alfaa19/gin-restAPI-course/config/mysql"
	"github.com/alfaa19/gin-restAPI-course/helper"
	"github.com/alfaa19/gin-restAPI-course/model"
	"github.com/alfaa19/gin-restAPI-course/request"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
)

func GetCourse(c *gin.Context) {
	keyword := c.Query("search")
	sortBy := c.Query("sort_by")

	var course []model.Course

	db := config.DB.Preload("Category")

	if keyword != "" {
		db = db.Where("title LIKE ?", "%"+keyword+"%")
	}

	if sortBy != "" {
		switch sortBy {
		case "lowest_price":
			db = db.Order("price")
		case "highest_price":
			db = db.Order("price DESC")
		case "free":
			db = db.Where("price = ?", 0)
		}
	}

	//get data course from database
	if err := db.Find(&course).Error; err != nil {
		helper.ResponseErrorJSON(c, http.StatusInternalServerError, err)
		return
	}

	//send success response
	helper.ResponseSuccessJSON(c, "", course)
}

func GetDetailCourse(c *gin.Context) {
	param := c.Param("id")

	id, err := strconv.Atoi(param)
	if err != nil {
		helper.ResponseErrorJSON(c, http.StatusBadRequest, err)
		return
	}

	var course model.Course

	if err := config.DB.First(&course, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			helper.ResponseErrorJSON(c, http.StatusNotFound, err)
			return
		}
		helper.ResponseErrorJSON(c, http.StatusInternalServerError, err)
		return
	}

	helper.ResponseSuccessJSON(c, "", course)
}

func CreateCourse(c *gin.Context) {
	var req request.Course

	//bind data from request
	if err := c.ShouldBind(&req); err != nil {
		helper.ResponseErrorJSON(c, http.StatusBadRequest, err)
		return
	}

	//validate request
	if err := req.Validate(); err != nil {
		helper.ResponseValidationErrorJson(c, "Error Validation", err.(validation.Errors))
		return
	}

	banner, err := c.FormFile("banner")

	if err != nil {
		helper.ResponseErrorJSON(c, http.StatusUnprocessableEntity, err)
		return
	}

	img, err := banner.Open()

	//create new cloudinary instance
	cld, _ := cloudinary.New()

	if err != nil {
		helper.ResponseErrorJSON(c, http.StatusInternalServerError, err)
		return
	}

	ctx := context.Background()

	//upload image banner
	upload, err := cld.Upload.Upload(ctx, img, uploader.UploadParams{})

	if err != nil {
		helper.ResponseErrorJSON(c, http.StatusBadRequest, err)

		return
	}

	var course = model.Course{
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		Banner:      upload.SecureURL,
		CategoryID:  req.CategoryID,
	}

	//insert into database
	if err := config.DB.Create(&course).Error; err != nil {
		helper.ResponseErrorJSON(c, http.StatusUnprocessableEntity, err)
		return
	}

	var createdCourse model.Course
	if err := config.DB.Preload("Category").First(&createdCourse, course.ID).Error; err != nil {
		helper.ResponseErrorJSON(c, http.StatusUnprocessableEntity, err)
		return
	}
	//send success response
	helper.ResponseSuccessJSON(c, "", createdCourse)

}

func GetTotalCourse(c *gin.Context) {
	var course []model.Course
	var count int64
	if err := config.DB.Preload("Category").Find(&course).Count(&count).Error; err != nil {
		helper.ResponseErrorJSON(c, http.StatusInternalServerError, err)
	}

	helper.ResponseSuccessJSON(c, "", gin.H{"total_courses": count})
}

func GetTotalFreeCourse(c *gin.Context) {
	var course []model.Course
	var count int64
	if err := config.DB.Preload("Category").Find(&course).Where("price = ?", 0).Count(&count).Error; err != nil {
		helper.ResponseErrorJSON(c, http.StatusInternalServerError, err)
	}

	helper.ResponseSuccessJSON(c, "", gin.H{"total_courses": count})
}

func UpdateCourse(c *gin.Context) {

	param := c.Param("id")

	id, err := strconv.Atoi(param)

	if err != nil {
		helper.ResponseErrorJSON(c, http.StatusInternalServerError, err)
		return
	}

	var existingCourse model.Course

	if err := config.DB.First(&existingCourse, id).Error; err != nil {
		helper.ResponseErrorJSON(c, http.StatusNotFound, err)
		return
	}

	var req request.Course

	//binding request data
	if err := c.ShouldBind(&req); err != nil {
		helper.ResponseErrorJSON(c, http.StatusBadRequest, err)
		return
	}

	// var newCourse = model.Course{}

	if title := c.PostForm("title"); title != "" {
		existingCourse.Title = req.Title
	}

	if desc := c.PostForm("desc"); desc != "" {
		existingCourse.Description = req.Description
	}

	if price := c.PostForm("price"); price != "" {
		existingCourse.Price = req.Price
	}

	if categoryID := c.PostForm("category_id"); categoryID != "" {
		existingCourse.CategoryID = req.CategoryID
	}

	banner, _ := c.FormFile("banner")

	if banner != nil {
		img, err := banner.Open()
		if err != nil {
			helper.ResponseErrorJSON(c, http.StatusInternalServerError, err)
			return
		}

		// Create a new cloudinary instance
		cld, _ := cloudinary.New()

		ctx := context.Background()

		// Upload image banner
		upload, err := cld.Upload.Upload(ctx, img, uploader.UploadParams{})
		if err != nil {
			helper.ResponseErrorJSON(c, http.StatusBadRequest, err)
			return
		}

		// Update the banner URL
		existingCourse.Banner = upload.SecureURL
	}

	// Update the course in the database
	if err := config.DB.Save(&existingCourse).Error; err != nil {
		helper.ResponseErrorJSON(c, http.StatusUnprocessableEntity, err)
		return
	}

	var updatedCourse model.Course
	if err := config.DB.Preload("Category").First(&updatedCourse, existingCourse.ID).Error; err != nil {
		helper.ResponseErrorJSON(c, http.StatusUnprocessableEntity, err)
		return
	}

	// Send success response
	helper.ResponseSuccessJSON(c, "", updatedCourse)
}

func DeleteCourse(c *gin.Context) {
	param := c.Param("id")

	id, err := strconv.Atoi(param)

	if err != nil {
		helper.ResponseErrorJSON(c, http.StatusInternalServerError, err)
		return
	}

	var course model.Course

	if err := config.DB.First(&course, id).Error; err != nil {
		helper.ResponseErrorJSON(c, http.StatusNotFound, err)
		return
	}

	if err := config.DB.Delete(&course).Error; err != nil {
		helper.ResponseErrorJSON(c, http.StatusNotFound, err)
		return
	}

	helper.ResponseSuccessJSON(c, "", "")

}
