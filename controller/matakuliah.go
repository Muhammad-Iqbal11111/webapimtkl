package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"webapimtkl/models"
	"time"
)

type MataKuliahInput struct {
	Id	int `json:"id" gorm:"primary_key"`
	KodeMataKuliah string `json:"kodematakuliah"`
	NamaMataKuliah string `json:"namamatakuliah"`
	JumlahSKS int `json:"jumlahsks"`
	DosenPengampu int `json:"dosenpengampu"`
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
	if err := c.ShouldBindJSON(&datamasuk); err != nil {
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
	if err := c.ShouldBindJSON(&datamasuk); err != nil {
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