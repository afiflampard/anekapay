package services

import (
	"anekakios/models"

	"gorm.io/gorm"
)

func Kulak(db *gorm.DB, orderUser models.OrderKulak, idUser int) error {
	tx := db.Begin()
	var (
		user    models.User
		product models.Product
	)
	if err := tx.First(&user, idUser).Error; err != nil {
		return err
	}
	if err := tx.Where("id_user = ? AND id = ?", orderUser.IdUserOrder, orderUser.IdProduct).Find(&product).Error; err != nil {
		return err
	}
	order := &models.Order{
		TanggalBeli: orderUser.TanggalBeli,
		BanyakOrder: orderUser.BanyakOrder,
		IdStateUser: orderUser.IdUserOrder,
		IdUser:      uint(idUser),
		NoState:     2,
	}
	if err := tx.Create(order).Error; err != nil {
		tx.Rollback()
		return err
	}
	prductBeli := &models.Product{
		IDUser:     uint(idUser),
		NamaProduk: product.NamaProduk,
		HargaBeli:  product.HargaBeli,
		HargaJual:  product.HargaJual + orderUser.HargaJual,
	}
	if err := tx.Create(prductBeli).Error; err != nil {
		tx.Rollback()
		return err
	}
	orderDetail := &models.OrderDetail{
		IDOrder:   order.ID,
		IDProduct: prductBeli.ID,
	}
	if err := tx.Create(orderDetail).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
