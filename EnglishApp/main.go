package main

import (
	"EnglishApp/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(db)
	fmt.Println("Hello")
	configGin(db)
}

func configGin(db *gorm.DB) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	v1 := r.Group("/v1")
	{
		bmeStudents := v1.Group("/bme")
		{
			bmeStudents.GET("/courses", controllers.GetAllBmeStudents(db))
		}

		users := v1.Group("/users")
		{
			users.GET("", controllers.FindAllUsers(db))
			users.POST("/create", controllers.CreateUser(db))
		}

		lessons := v1.Group("/lesson")
		{
			lessons.GET("", controllers.FindAllLessons(db))
			lessons.POST("/create", controllers.CreateLesson(db))
		}

		skills := v1.Group("/skills")
		{
			skills.GET("", controllers.FindAllSkills(db))
			skills.POST("/create", controllers.CreateSkill(db))
		}

		feedbacks := v1.Group("/feedbacks")
		{
			feedbacks.GET("", controllers.FindAllFeedbacks(db))
			feedbacks.POST("/create", controllers.CreateFeedback(db))
		}

		classCodes := v1.Group("/class-codes")
		{
			classCodes.GET("", controllers.FindAllClassCodes(db))
			classCodes.POST("/create", controllers.CreateClassCode(db))
		}

		classSkills := v1.Group("/class-skills")
		{
			classSkills.GET("", controllers.FindAllClassSkills(db))
			classSkills.POST("/create", controllers.CreateClassSkill(db))
		}

		classUsers := v1.Group("/class-users")
		{
			classUsers.GET("", controllers.FindAllClassUsers(db))
			classUsers.POST("/create", controllers.CreateClassUser(db))
		}

		userFeedbacks := v1.Group("/user-feedback")
		{
			userFeedbacks.GET("", controllers.FindAllUserFeedbacks(db))
			userFeedbacks.POST("/create", controllers.CreateUserFeedback(db))
		}

		lessonRoadmaps := v1.Group("/lesson-roadmap")
		{
			lessonRoadmaps.GET("", controllers.FindAllLessonRoadmaps(db))
			lessonRoadmaps.POST("/create", controllers.CreateLessonRoadmap(db))
		}
	}
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
