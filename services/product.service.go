package services

import (
	"anekakios/models"

	"gorm.io/gorm"
)

func CreateProduct(db *gorm.DB, product models.Product, id int) (*models.Product, error) {
	var productTemp models.Product
	if err := db.Where("nama_produk = ? AND user_id = ?", product.NamaProduk, id).Find(&productTemp).Error; err != nil {
		addProduct := models.Product{
			IDUser:     uint(id),
			NamaProduk: product.NamaProduk,
			HargaBeli:  product.HargaBeli,
			HargaJual:  product.HargaJual,
			Stock:      product.Stock,
		}
		if err := db.Create(&addProduct).Error; err != nil {
			return nil, err
		}
		return &addProduct, nil
	}
	productTemp.HargaBeli = product.HargaBeli
	productTemp.HargaJual = product.HargaJual
	productTemp.Stock = product.Stock
	if err := db.Save(&productTemp).Error; err != nil {
		return nil, err
	}
	return &productTemp, nil
}

func GetProductID(db *gorm.DB, idUser int, idProduct int) (*models.Product, error) {
	var product models.Product
	if err := db.Debug().Where("id_user = ? AND id = ?", idUser, idProduct).Preload("User").Preload("User.Role").Find(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func GetAllProduct(db *gorm.DB, idUser int) ([]models.Product, error) {
	var products []models.Product
	if err := db.Where("id_user = ?", idUser).Preload("User").Preload("User.Role").Find(&products).Error; err != nil {
		return products, err
	}
	return products, nil
}

func DeleteProduct(db *gorm.DB, idUser, idProduct int) (map[string]string, error) {
	var product models.Product
	if err := db.Where("id_user = ? AND id = ?", idUser, idProduct).Delete(&product).Error; err != nil {
		return map[string]string{
			"message": "Product tidak ada",
		}, err
	}
	return map[string]string{
		"message": "Product telah terhapus",
	}, nil
}
func UpdateProduct(db *gorm.DB, product models.Product, idProduct int) (*models.Product, error) {
	var tempProduct models.Product
	if err := db.Where("id_user = ? AND id = ?", product.IDUser, idProduct).First(&tempProduct).Error; err != nil {
		return nil, err
	}
	tempProduct.HargaBeli = product.HargaBeli
	tempProduct.HargaJual = product.HargaJual
	tempProduct.Stock = product.Stock
	tempProduct.NamaProduk = product.NamaProduk
	db.Save(&tempProduct)
	return &tempProduct, nil
}
