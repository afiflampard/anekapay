package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	OrderDetails []OrderDetail `gorm:"foreignKey:IDProduct"`
	IDUser       uint          `gorm:"column:id_user"`
	User         User          `gorm:"foreignKey:IDUser"`
	NamaProduk   string        `gorm:"column:nama_produk; type:varchar(255); not null" json:"nama_produk"`
	HargaBeli    uint          `gorm:"column:harga_beli" json:"harga_beli"`
	HargaJual    uint          `gorm:"column:harga_jual" json:"harga_jual"`
	Stock        uint          `gorm:"column:stock" json:"stock"`
}
