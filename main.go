package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webapimtkl/controller"
	"webapimtkl/models"
)

func main() {

	r := gin.Default()
	v1 := r.Group("/v1")
	//Models
	db := models.SetUpModels()
	v1.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	v1.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message" : "web api mata kuliah",
		})
	})
		
	//GET All Data
	v1.GET("/matakuliah", controller.GetData)
	//POST Data >> Create Data
	v1.POST("/matakuliah", controller.CreateData)
	//Update Data >> Update Data
	v1.PUT("/matakuliah/:kodematakuliah", controller.UpdateData)
    //Delete Data >> Delete data
	v1.DELETE("/mahasiswa/:kodematakuliah", controller.DeleteData)
	r.Run()
}
