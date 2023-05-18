package config

import (
	"fmt"
	"log"
	"os"

	"github.com/alfaa19/gin-restAPI-course/model"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading env file", err)
	}

	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DBNAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&model.Course{}, &model.Category{}, &model.User{})

	if err != nil {
		log.Fatal(err)
	}

	seedCourses(db)
	seedUser(db)
	DB = db
}

func seedCourses(db *gorm.DB) {

	var count int64
	result := db.Model(&model.Category{}).Count(&count)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	// Membuat data category

	categories := []model.Category{
		{
			Name: "Programming",
		},
		{
			Name: "Design",
		},
		{
			Name: "Data Science",
		},
		{
			Name: "Machine Learning",
		},
	}

	// Membuat data course
	courses := []model.Course{
		{Title: "Belajar Golang", Description: "Belajar dasar - dasar  bahasa pemrograman golang", Price: 19.99, Category: &model.Category{Name: categories[0].Name}},
		{Title: "UI UX Advance", Description: "Belajar UI UX Advance  ", Price: 29.99, Category: &model.Category{Name: categories[1].Name}},
		{Title: "Intro to Python", Description: "Learning Python for beginner ", Price: 9.99, Category: &model.Category{Name: categories[2].Name}},
		{Title: "Machine Learning Implementation", Description: "Implementing Machine Learning Algorithm", Price: 10000, Category: &model.Category{Name: categories[3].Name}},
		{Title: "Neural Network", Description: "Algoirthm of Machine Learning", CategoryID: 1, Price: 110000},
	}
	if count == 0 {

		// Menyimpan data course ke database
		for _, course := range courses {
			if err := db.Create(&course).Error; err != nil {
				log.Fatalf("Failed to create course: %v", err)
			}
		}
	}

	//
	log.Println("Seeder completed successfully")
}

func seedUser(db *gorm.DB) {
	var count int64
	result := db.Model(&model.User{}).Count(&count)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	//hash password
	pass, _ := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	pass1, _ := bcrypt.GenerateFromPassword([]byte("user"), bcrypt.DefaultCost)
	//membuat data user
	users := []model.User{
		{Username: "admin", Password: string(pass), Role: "admin"},
		{Username: "user", Password: string(pass1), Role: "user"},
	}

	if count == 0 {
		for _, user := range users {
			if err := db.Create(&user).Error; err != nil {
				log.Fatalf("Failed to create user: %v", err)
			}
		}
	}
	log.Println("Seeder completed successfully")
}
