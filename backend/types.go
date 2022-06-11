package backend

import (
	"gorm.io/gorm"
	"time"
)

type Inventory struct {
	gorm.DB
}
type Price float64

type Tag struct {
	Barcode string
	Name    string
}

type Product struct {
	Tag     `gorm:"unique;"`
	ID      int
	Price   Price
	StockID int
	Stock   *Stock
}

type ProductList []Product

type Stock struct {
	ID        int
	ProductID int
	Product   *Product
	Count     int
}

type SellEntry struct {
	ID int
	Tag
	Price  Price
	Count  int
	SellID int
}

type Sell struct {
	ID      int `gorm:"PrimaryKey"`
	Time    time.Time
	Entries []SellEntry
}
