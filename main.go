package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webapimtkl/controller"
	"webapimtkl/models"
)

func main() {

	r := gin.Default()
	//Models
	db := models.SetUpModels()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message" : "web api mata kuliah",
		})
	})
		
	//GET All Data
	r.GET("/matakuliah", controller.GetData)
	//POST Data >> Create Data
	r.POST("/matakuliah", controller.CreateData)
	//Update Data >> Update Data
	r.PUT("/matakuliah/:kodematakuliah", controller.UpdateData)
    //Delete Data >> Delete data
	r.DELETE("/mahasiswa/:kodematakuliah", controller.DeleteData)
	r.Run()
}
