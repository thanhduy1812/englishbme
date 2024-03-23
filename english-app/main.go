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

var (
	ServiceName = "english-app"
	port        = "8080"
)

func main() {

	//dsn := os.Getenv("DB_CONN_STR_SV")
	dsn := "root:123456@tcp(103.161.173.59:3308)/english_uat?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(db)
	configGin(db)

	//http.HandleFunc("/ping", func(w http.ResponseWriter, req *http.Request) {
	//	_, err := w.Write([]byte(fmt.Sprintf("ping ok %s", ServiceName)))
	//	if err != nil {
	//		return
	//	}
	//})
	fmt.Println("start service with name: ", ServiceName)
	fmt.Println("start service with port: ", port)
	//log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
	// Iterate over the network interfaces
	hostIP := os.Getenv("HOST_IP")
	if hostIP == "" {
		log.Fatal("HOST_IP environment variable not set")
	}

	log.Printf("Host IP: %s", hostIP)
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
			bmeStudents.GET("/courses", controllers.GetAllBmeCourse(db))
			bmeStudents.GET("/courses/find-by-key", controllers.FindBmeCourseByKey(db))
			bmeStudents.GET("/courses/search-by-column", controllers.SearchBmeCourseByColumn(db))
			bmeStudents.PUT("/courses/update/:id", controllers.UpdateBmeCourse(db))
			bmeStudents.POST("/courses/create", controllers.CreateBmeCourse(db))
			bmeStudents.DELETE("/courses/:id", controllers.DeleteBmeCourse(db))

			bmeStudents.GET("/courses-other", controllers.GetAllBmeCourseHocBu(db))
			bmeStudents.GET("/courses-other/find-by-key", controllers.FindBmeCourseHocBuByKey(db))

			bmeStudents.GET("/students", controllers.GetAllBmeStudents(db))
			bmeStudents.GET("/students/find-by-key", controllers.FindBmeStudentByKey(db))
			bmeStudents.GET("/students/search-by-column", controllers.SearchBmeStudentByColumn(db))
		}

		users := v1.Group("/users")
		{
			users.GET("", controllers.FindAllUsers(db))
			users.POST("/create/users", controllers.CreateUsers(db))
			users.POST("/create", controllers.CreateUser(db))
			users.GET("/find-by-key", controllers.FindUserByKey(db))
			users.GET("/search-by-column", controllers.SearchUserByColumn(db))
			users.PUT("/update/:id", controllers.UpdateUser(db))
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
			userFeedbacks.GET("/find-by-key", controllers.FindUserFeedbackByKey(db))
			userFeedbacks.POST("/find-by-lessons", controllers.FindUserFeedbackByLessonRoadmapIds(db))
		}

		lessonRoadmaps := v1.Group("/lesson-roadmap")
		{
			lessonRoadmaps.GET("", controllers.FindAllLessonRoadmaps(db))
			lessonRoadmaps.GET("/find-by-key", controllers.FindLessonRoadMapByKey(db))
			lessonRoadmaps.POST("/create", controllers.CreateLessonRoadmap(db))
		}
	}
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
