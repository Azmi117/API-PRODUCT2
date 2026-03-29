package models

import "gorm.io/gorm"

// 1. Buat struct untuk model

type Product struct {
	/*2. jika menggunakan GORM
	tuliskan juga gorm.model
	agar beberapa field dibuat otomatis */
	gorm.Model
	Name     string `json:"name"`
	Quantity int    `json:"qty"`
	Quality  string `json:"qly"`
}
