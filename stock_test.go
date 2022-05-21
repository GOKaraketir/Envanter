package main

import (
	"testing"
)

func TestStock(t *testing.T) {

	inventory := createDB(t)

	tag0 := testProductToProduct(testProducts[0]).Tag

	stock, err := inventory.AddToStock(tag0, 20)

	if err != nil {
		t.Fatal(err)
	}
	getStock, err := inventory.GetStock(tag0)
	if err != nil {
		t.Fatal(err)
	}

	if getStock.Product.Tag != stock.Product.Tag {
		t.Fatal("Not Equal")
	}

	if getStock.Count != stock.Count {
		t.Fatal("Not Equal")
	}

	if getStock.Count != (testProducts[0].Count + 20) {
		t.Fatal("Not equal increased")
	}
}
