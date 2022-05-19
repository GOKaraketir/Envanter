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
)

func Hello() {
	fmt.Println("Hello")
}

func Initialize() (inventory Inventory, err error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
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

func (this *Inventory) CreateProduct(barcode, name string, price Price, count int) (product Product, err error) {
	err = nil

	var products []Product
	if barcode != "" {
		this.Find(&products, "Barcode = ?", barcode)
		if len(products) != 0 {
			err = errors.New(ERR_BARCODE_EXIST)
		}
	} else {
		this.Find(&products, "Barcode = ? AND Name = ?", barcode, name)
		if len(products) != 0 {
			err = errors.New(ERR_NAME_EXIST_WITHOUT_BARCODE)
			return
		}
	}

	product = Product{
		Barcode: barcode,
		Name:    name,
		Price:   price,
		Stock: &Stock{
			Barcode: barcode,
			Name:    name,
			Product: &product,
			Count:   count,
		},
	}

	this.Create(&product)

	return
}

func (this *Inventory) GetProduct(barcode, name string) (product Product, err error) {
	err = nil
	var products []Product
	this.Preload("Stock").Find(&products, "Barcode = ? AND Name = ? ", barcode, name)
	if len(products) == 0 {
		err = errors.New(ERR_PROD_NOT_FOUND)
		return
	}
	product = products[0]
	return
}

func IsSameProduct(product1, product2 Product) bool {
	if product1.Barcode != product2.Barcode {
		return false
	} else if product1.Name != product2.Name {
		return false
	}
	return true
}

func (receiver *Inventory) DoneSell(sell *Sell) (err error) {
	for _, entry := range sell.Entries {
		product, err := receiver.GetProduct(entry.Product.Barcode, entry.Product.Name)
		if err != nil {
			return err
		}
		product.Stock.Count -= entry.Count
		receiver.Save(&product.Stock)
	}
	return nil
}
