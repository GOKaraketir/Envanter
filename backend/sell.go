package backend

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

func (receiver *Inventory) DoneSell(sell *Sell) (err error) {
	for _, entry := range sell.Entries {
		product, err := receiver.GetProduct(entry.Tag)
		if err != nil {
			return err
		}
		product.Stock.Count -= entry.Count
		receiver.Save(&product.Stock)
	}
	return nil
}

func (receiver *SellEntry) UpdatePrice(newPrice Price) {
	receiver.Price = newPrice
}

func (receiver *SellEntry) UpdateCount(newCount int) {
	receiver.Count = newCount
}
