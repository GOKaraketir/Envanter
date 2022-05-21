package backend

import (
	"errors"
	"github.com/GOKaraketir/Envanter/backend/myDate"
	"gorm.io/gorm/clause"
	"time"
)

func (receiver *Sell) Sum() (total Price) {

	for _, entry := range receiver.Entries {
		total += Price(entry.Count) * entry.Price
	}

	return
}

func (receiver *Sell) AddSellEntry(newEntry SellEntry) {
	for i, sellEntry := range receiver.Entries {
		if sellEntry.Tag == newEntry.Tag {
			receiver.Entries[i].Count += newEntry.Count
			return
		}
	}
	receiver.Entries = append(receiver.Entries, newEntry)
}

func (receiver *Inventory) CommitSell(sell *Sell) (err error) {
	for _, entry := range sell.Entries {
		_, err := receiver.ReduceFromStock(entry.Tag, entry.Count)
		if err != nil {
			return err
		}
	}
	receiver.Save(sell)
	return nil
}

func (receiver *Inventory) UndoSell(ID int) (err error) {
	sell, err := receiver.GetSell(ID)
	if err != nil {
		return err
	}
	for _, entry := range sell.Entries {
		_, err := receiver.AddToStock(entry.Tag, entry.Count)
		if err != nil {
			return err
		}
		receiver.Delete(&entry)
	}
	receiver.Delete(sell)
	return nil
}

func (receiver *SellEntry) UpdatePrice(newPrice Price) {
	receiver.Price = newPrice
}

func (receiver *SellEntry) UpdateCount(newCount int) {
	receiver.Count = newCount
}

func NewSell() Sell {
	return Sell{
		Date: myDate.FromTime(time.Now()),
	}
}

func CreateSellEntry(product Product, count int) SellEntry {
	return SellEntry{
		Tag:   product.Tag,
		Price: product.Price,
		Count: count,
	}
}

func (receiver *Inventory) GetSell(ID int) (*Sell, error) {
	var sells []Sell
	receiver.Preload(clause.Associations).Find(&sells, ID)
	if len(sells) == 0 {
		err := errors.New(ERR_SELL_NOT_FOUND)
		return nil, err
	}

	return &sells[0], nil
}

func (receiver *Sell) GetSellEntry(tag Tag) (*SellEntry, error) {
	for i, entry := range receiver.Entries {
		if entry.Tag == tag {
			return &receiver.Entries[i], nil
		}
	}
	return nil, errors.New(ERR_SELL_ENTRY_NOT_FOUND)
}
