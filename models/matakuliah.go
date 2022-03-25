package models

type MataKuliah struct{
	Id	int `json:"id" gorm:"primary_key"`
	KodeMataKuliah string `json:"kodematakuliah"`
	NamaMataKuliah string `json:"namamatakuliah"`
	JumlahSKS int `json:"jumlahsks"`
	DosenPengampu int `json:"dosenpengampu"`
}