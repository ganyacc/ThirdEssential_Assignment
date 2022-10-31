package product

import (
	database "ThirdEssentials/db"
	"ThirdEssentials/services/payload"
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

func AddProduct(db *gorm.DB, userId int, request *payload.Product) (*database.Product, error) {

	product := database.Product{
		Name:      request.Name,
		ImageUrl:  request.ImageURL,
		Price:     request.Price,
		UserId:    userId,
		CreatedAt: time.Now(),
	}

	err := db.Create(&product).Error

	if err != nil {
		return nil, err
	}

	productActivity := database.UserProdActivity{
		ProductID: product.Id,
		ProdName:  product.Name,
		Image:     *request.ImageURL,
		Price:     request.Price,
		UserID:    userId,
		Type:      "AddProduct",
		CreatedAt: time.Now(),
	}

	err = db.Create(&productActivity).Error

	if err != nil {
		return nil, err
	}

	return &product, nil

}

func GetAllProducts(db *gorm.DB, userID int) ([]database.Product, error) {

	products := []database.Product{}

	err := db.Debug().Model(database.Product{}).Where(database.Product{
		UserId: userID,
	}).Find(&products).Error

	if err != nil {
		return nil, err
	}

	return products, nil
}

func UpdateProduct(db *gorm.DB, userID, prdId int, request *payload.Product) (*database.Product, error) {

	product, err := GetProductById(db, prdId, userID)
	if err != nil || product == nil {
		return nil, errors.New("no product found")
	}

	if product.UserId != userID {
		return nil, errors.New("not authorized to perform this action")
	}

	updateProd := database.Product{
		Id:        prdId,
		Name:      request.Name,
		ImageUrl:  request.ImageURL,
		Price:     request.Price,
		UpdatedAt: time.Now(),
	}

	err = db.Model(&database.Product{}).Update(&updateProd).Error
	if err != nil {
		return nil, err
	}

	productActivity := database.UserProdActivity{
		ProductID: product.Id,
		ProdName:  request.Name,
		Image:     *request.ImageURL,
		Price:     request.Price,
		UserID:    userID,
		Type:      "UpdateProduct",
		CreatedAt: time.Now(),
	}

	err = db.Create(&productActivity).Error

	if err != nil {
		return nil, err
	}

	return &updateProd, nil
}

func DeleteProductbyID(db *gorm.DB, productID, userId int) error {

	err := db.Debug().Model(&database.Product{}).Where(&database.Product{
		Id:     productID,
		UserId: userId,
	}).Delete(&database.Product{
		Id:     productID,
		UserId: userId,
	}).Error

	if err != nil {
		return err
	}

	productActivity := database.UserProdActivity{
		ProductID: productID,
		UserID:    userId,
		Type:      "DeleteProduct",
	}

	return db.Create(&productActivity).Error

}

func GetUserProductActivity(db *gorm.DB, userID int) ([]database.UserProdActivity, error) {

	products := []database.UserProdActivity{}

	err := db.Debug().Model(database.UserProdActivity{}).Where(database.UserProdActivity{
		UserID: userID,
	}).Find(&products).Error

	if err != nil {
		return nil, err
	}

	return products, nil
}

func GetProductById(db *gorm.DB, productID, userID int) (*database.Product, error) {

	var products database.Product

	err := db.Debug().Model(database.Product{}).Where(database.Product{
		Id:     productID,
		UserId: userID,
	}).Find(&products).Error

	if err != nil {
		return nil, err
	}

	return &products, nil
}
