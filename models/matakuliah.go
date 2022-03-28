package models

type MataKuliah struct{
	Id	int `json:"id" binding: "required,uuid" gorm:"primary_key"`
	KodeMataKuliah string `json:"kodematakuliah" binding: "required"`
	NamaMataKuliah string `json:"namamatakuliah" binding: "required,min=3"`
	JumlahSKS int `json:"jumlahsks" binding: "required"`
	DosenPengampu int `json:"dosenpengampu" binding: "required"`
}