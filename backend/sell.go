package backend

import "github.com/GOKaraketir/Envanter/backend/myDate"

type SellEntry struct {
	Product Product `gorm:"embedded"`
	Count   int
	SellID  int
}

type Sell struct {
	ID      int           `gorm:"PrimaryKey"`
	Date    myDate.MyDate `gorm:"embedded"`
	Entries []SellEntry
}

func (receiver *Sell) Sum() (total Price) {

	for _, entry := range receiver.Entries {
		total += Price(entry.Count) * entry.Product.Price
	}

	return
}

func (receiver *Sell) AddSellEntry(newEntry SellEntry) {
	for i, sellEntry := range receiver.Entries {
		if IsSameProduct(sellEntry.Product, newEntry.Product) {
			receiver.Entries[i].Count += newEntry.Count
			return
		}
	}
	receiver.Entries = append(receiver.Entries, newEntry)
}
