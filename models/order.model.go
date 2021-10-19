package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	TanggalBeli time.Time  `gorm:"column:tanggal_beli" json:"tanggal_beli"`
	BanyakOrder uint       `gorm:"column:banyak_order" json:"banyak_order"`
	IdStateUser uint       `gorm:"column:id_state_user" json:"id_state_user"`
	StateUser   User       `gorm:"foreignKey:IdStateUser"`
	IdUser      uint       `gorm:"column:user_id" json:"userId"`
	User        User       `gorm:"foreignKey:IdUser"`
	NoState     uint       `gorm:"column:state_no" json:"state_no"`
	OrderState  OrderState `gorm:"foreignKey:NoState"`
}

type OrderDetail struct {
	gorm.Model
	IDOrder   uint  `gorm:"column:order_id" json:"id_order"`
	IDProduct uint  `gorm:"column:product_id" json:"id_product"`
	Order     Order `gorm:"foreignKey:IDOrder"`
}

type OrderState struct {
	gorm.Model
	Name string `gorm:"column:state_name"`
}

type OrderKulak struct {
	TanggalBeli time.Time `json:"tanggal_beli"`
	BanyakOrder uint      `json:"banyak_order"`
	HargaJual   uint      `json:"harga_jual"`
	IdProduct   uint      `json:"id_product"`
	IdUserOrder uint      `json:"id_user_order"`
	Stock       uint      `json:"stock"`
}
