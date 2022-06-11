package main

import (
	"github.com/GOKaraketir/Envanter/backend"
	"testing"
)

func TestSell(t *testing.T) {
	inventory := createDB(t)

	aSell := backend.NewSell()
	bSell := backend.NewSell()

	aSell.AddSellEntry(backend.CreateSellEntry(testProductToProduct(testProducts[0]), 5))
	bSell.AddSellEntry(backend.CreateSellEntry(testProductToProduct(testProducts[0]), 5))

	err := inventory.CommitSell(aSell)
	if err != nil {
		t.Fatal(err)
	}

	stock, err := inventory.GetStock(testProductToProduct(testProducts[0]).Tag)
	if err != nil {
		t.Fatal(err)
	}
	if testProducts[0].Count-5 != stock.Count {
		t.Fatal(stock.Product.Stock)
	}

	err = inventory.CommitSell(bSell)
	if err != nil {
		t.Fatal(err)
	}

	afterSell, err := inventory.GetSell(bSell.ID)
	if err != nil {
		t.Fatal(err)
	}
	if afterSell.Entries[0].SellID != 2 {
		t.Fatal("Entiries does not mach")
	}

	stock, err = inventory.GetStock(testProductToProduct(testProducts[0]).Tag)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Before undo", stock)

	err = inventory.UndoSell(1)
	if err != nil {
		t.Fatal(err)
	}

	stock, err = inventory.GetStock(testProductToProduct(testProducts[0]).Tag)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("After undo", stock)

	_, err = inventory.GetSell(1)
	t.Log("Satis Silindi mi: ", err)
	if err == nil {
		t.Fatal("Undo Not Worked")
	}

	err = inventory.UndoSell(2)
	if err != nil {
		t.Fatal(err)
	}

	_, err = inventory.GetSell(2)
	t.Log("Satis Silindi mi: ", err)
	if err == nil {
		t.Fatal("Undo Not Worked")
	}

	stock, err = inventory.GetStock(testProductToProduct(testProducts[0]).Tag)
	if stock.Count != testProducts[0].Count {
		t.Fatal("Undo Error")
	}
}

func TestSellEntryUpdate(t *testing.T) {
	inventory := createDB(t)

	aSell := backend.NewSell()

	aSell.AddSellEntry(backend.CreateSellEntry(testProductToProduct(testProducts[0]), 5))

	entry, err := aSell.GetSellEntry(testProductToProduct(testProducts[0]).Tag)
	if err != nil {
		t.Fatal(err)
	}

	entry.UpdateCount(6)

	err = inventory.CommitSell(aSell)
	if err != nil {
		t.Fatal(err)
	}

	stock, err := inventory.GetStock(testProductToProduct(testProducts[0]).Tag)
	if err != nil {
		t.Fatal(err)
	}
	if testProducts[0].Count-6 != stock.Count {
		t.Fatal(stock.Product.Stock)
	}

	aSell.ID++
	err = inventory.CommitSell(aSell)
	if err != nil {
		t.Fatal(err)
	}

	afterSell, err := inventory.GetSell(aSell.ID)
	if err != nil {
		t.Fatal(err)
	}
	if afterSell.Entries[0].SellID != 2 {
		t.Fatal("Entiries does not mach")
	}

}
