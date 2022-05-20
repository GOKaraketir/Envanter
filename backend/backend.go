package backend

import (
	"errors"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	ERR_BARCODE_EXIST              string = "Barkod zaten mevcut"
	ERR_NAME_EXIST_WITHOUT_BARCODE        = "Aynı ürün mevcut"
	ERR_PROD_NOT_FOUND                    = "Ürün Bulunamadı"

	NO_BARCODE = "0000000000000"
)

func Hello() {
	fmt.Println("Hello")
}

func Initialize(fileName string) (inventory Inventory, err error) {
	db, err := gorm.Open(sqlite.Open(fileName), &gorm.Config{})
	if err != nil {
		return
	}

	if errMig := db.AutoMigrate(&Product{}, &Stock{}, &SellEntry{}, &Sell{}); errMig != nil {
		err = errMig
		return
	}

	inventory.DB = *db
	return
}

func (receiver *Inventory) CreateProduct(tag Tag, price Price, count int) (product Product, err error) {
	err = nil

	var products []Product
	if tag.Barcode != NO_BARCODE {
		receiver.Find(&products, Product{Tag: tag})
		if len(products) != 0 {
			err = errors.New(ERR_BARCODE_EXIST)
			return
		}
	} else {
		receiver.Find(&products, Product{Tag: tag})
		if len(products) != 0 {
			err = errors.New(ERR_NAME_EXIST_WITHOUT_BARCODE)
			return
		}
	}

	product = Product{
		Tag:   tag,
		Price: price,
		Stock: &Stock{
			Product: &product,
			Count:   count,
		},
	}

	receiver.Create(&product)

	return
}

func (receiver *Inventory) UpdateProduct(tag Tag, newProduct Product) (product Product, err error) {
	err = nil

	product, err = receiver.GetProduct(tag)
	if err != nil {
		return
	}

	newProduct.ID = product.ID

	receiver.Save(&newProduct)

	return
}

func (receiver *Inventory) DeleteProduct(tag Tag) (product Product, err error) {

	product, err = receiver.GetProduct(tag)

	receiver.Delete(&product)
	receiver.Delete(&product.Stock)

	return
}

func (receiver *Inventory) GetProduct(tag Tag) (product Product, err error) {
	err = nil
	var products []Product
	receiver.Preload("Stock").Find(&products, Product{Tag: tag})
	if len(products) == 0 {
		err = errors.New(ERR_PROD_NOT_FOUND)
		return
	}
	product = products[0]
	return
}

func (receiver *Inventory) GetStock(tag Tag) (Stock, error) {

	product, err := receiver.GetProduct(tag)
	if err != nil {
		return Stock{}, err
	}

	return *product.Stock, nil
}

func (receiver *Inventory) UpdateStock(tag Tag, count int) (Stock, error) {

	stock, err := receiver.GetStock(tag)
	if err != nil {
		return Stock{}, err
	}

	stock.Count = count

	receiver.Save(&stock)

	return stock, nil
}

func (receiver Inventory) AddToStock(tag Tag, count int) (Stock, error) {
	stock, err := receiver.GetStock(tag)
	if err != nil {
		return Stock{}, err
	}

	updatedStock, err := receiver.UpdateStock(tag, stock.Count+count)
	if err != nil {
		return Stock{}, err
	}

	return updatedStock, nil
}

func (receiver Inventory) ReduceFromStock(tag Tag, count int) (Stock, error) {
	return receiver.AddToStock(tag, -1*count)
}
