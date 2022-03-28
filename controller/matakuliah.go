package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/go-playground/validator/v10"
	"net/http"
	"webapimtkl/models"
	"time"
	"fmt"
)

type MataKuliahInput struct {
	Id	int `json:"id" binding: "required,uuid" gorm:"primary_key"`
	KodeMataKuliah string `json:"kodematakuliah" binding: "required"`
	NamaMataKuliah string `json:"namamatakuliah" binding: "required,min=3"`
	JumlahSKS int `json:"jumlahsks" binding: "required"`
	DosenPengampu int `json:"dosenpengampu" binding: "required"`
}

//GET Data
func GetData(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var mtkl []models.MataKuliah
	db.Find(&mtkl)
	c.JSON(http.StatusOK, gin.H{
		"data" : mtkl,
		"time" : time.Now(),
	})
}

//POST Data >> Create Data
func CreateData(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	//validasi inputan
	var datamasuk MataKuliahInput
	err := c.ShouldBindJSON(&datamasuk); 
	if err != nil {
		errorMessages := []string {}
		for _,e := range err.(validator.ValidationErrors){
			switch  e.Tag() {
			case "required":
				report := fmt.Sprintf("%s is required", e.Field())
				errorMessages = append (errorMessages, report)
			case "min":
				report := fmt.Sprintf("%s must be more than 5 characters", e.Field())
				errorMessages = append (errorMessages, report)
			}
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}
//	prose input data
	mtkl := models.MataKuliah{
		Id:  datamasuk.Id,
		KodeMataKuliah: datamasuk.KodeMataKuliah,
		NamaMataKuliah: datamasuk.NamaMataKuliah,
		JumlahSKS: datamasuk.JumlahSKS,
		DosenPengampu: datamasuk.DosenPengampu,
	}

	db.Create(&mtkl)

//	menampilkan hasil
	c.JSON(http.StatusOK, gin.H{
		"Message" : "Sukses input data",
		"Data" : mtkl,
		"time" : time.Now(),
	})
}

//UPDATE Data >> Update Data
func UpdateData(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//cek data
	var mtkl models.MataKuliah
	if err := db.Where("namamatakuliah < 3", c.Param("namamatakuliah")).First(&mtkl).Error;
	err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Data mahasiswa tidak di temukan",
		})
		return
	}

	//validasi inputan
	var datamasuk MataKuliahInput
	err := c.ShouldBindJSON(&datamasuk) 
	if err != nil {
		errorMessages := []string{}
		for _,e := range err.(validator.ValidationErrors){
			switch  e.Tag() {
			case "required":
				report := fmt.Sprintf("%s is required", e.Field())
				errorMessages = append (errorMessages, report)
			case "min":
				report := fmt.Sprintf("%s must be more than 5 characters", e.Field())
				errorMessages = append (errorMessages, report)
			}
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}
	//	prose Ubah data
	db.Model(&mtkl).Update(&datamasuk)

	//	menampilkan hasil
	c.JSON(http.StatusOK, gin.H{
		"Message" : "Sukses ubah data",
		"Data" : mtkl,
		"time" : time.Now(),
	})
}

// Delete Data >> Hapus Data
func DeleteData(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//cek data
	var mtkl models.MataKuliah
	if err := db.Where("namamatakuliah < 3", c.Param("namamatakuliah")).First(&mtkl).Error;
		err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Data mahasiswa tidak di temukan",
		})
		return
	}
	//	prose hapus data
	db.Delete(&mtkl)

	//	menampilkan hasil
	c.JSON(http.StatusOK, gin.H{
		"Data" : true,
	})
}