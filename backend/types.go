package backend

import (
	"gorm.io/gorm"
)

type Inventory struct {
	gorm.DB
}
type Price float64

type Product struct {
	Barcode string `gorm:"primarykey;unique"`
	Name    string `gorm:"primarykey"`
	Price   Price
	Stock   *Stock `gorm:"ForeignKey:Barcode,Name;References:Barcode,Name"`
}

type Stock struct {
	Barcode string   `gorm:"primarykey;unique"`
	Name    string   `gorm:"primarykey"`
	Product *Product `gorm:"ForeignKey:Barcode,Name;References:Barcode,Name"`
	Count   int
}
