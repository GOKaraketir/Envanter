package main

import (
	"github.com/GOKaraketir/Envanter/backend"
	"os"
	"testing"
)

type TestProduct struct {
	Barcode string
	Name    string
	Price   backend.Price
	Count   int
}

var testProducts = [...]TestProduct{
	{
		Barcode: "1111111111111",
		Name:    "Ürün 1",
		Price:   10,
		Count:   13,
	},
	{
		Barcode: "2222222222222",
		Name:    "Ürün 2",
		Price:   0.01,
		Count:   10,
	},
}

const testDBname = "TestTest.db"

func createDB(t *testing.T) backend.Inventory {
	_ = os.Remove(testDBname)

	inventory, err := backend.Initialize(testDBname)
	if err != nil {
		t.Fatal(err)
	}

	for _, product := range testProducts {
		_, err := inventory.CreateProduct(backend.Tag{
			Barcode: product.Barcode,
			Name:    product.Name,
		}, product.Price, product.Count)
		if err != nil {
			t.Fatal(err)
		}
	}

	return inventory
}

func TestDelete(t *testing.T) {
	inventory := createDB(t)

	for _, product := range testProducts {
		product, err := inventory.GetProduct(backend.Tag{
			Barcode: product.Barcode,
			Name:    product.Name,
		})
		if err != nil {
			t.Fatal(err)
		}

		_, err = inventory.DeleteProduct(product.Tag)
		if err != nil {
			t.Fatal(err)
		}
	}

}

func TestUpdate(t *testing.T) {
	inventory := createDB(t)
	var err error
	var prod backend.Product

	inventory.Preload("Stock").First(&prod)

	tag := prod.Tag

	prod.Price = 22
	prod.Name = "ASDASD"

	_, err = inventory.UpdateProduct(tag, prod)
	if err != nil {
		t.Fatal(err)
	}

	_, err = inventory.UpdateStock(prod.Tag, 22)
	if err != nil {
		t.Fatal(err)
	}
}

func TestStock(t *testing.T) {

	inventory := createDB(t)

	tag0 := backend.Tag{
		Barcode: testProducts[0].Barcode,
		Name:    testProducts[0].Name,
	}
	stock, err := inventory.AddToStock(tag0, 20)
	if err != nil {
		t.Fatal(err)
	}
	getStock, err := inventory.GetStock(tag0)
	if err != nil {
		t.Fatal(err)
	}

	if getStock != stock {
		t.Fatal("Not Equal")
	}

	if getStock.Count != (testProducts[0].Count + 20) {
		t.Fatal("Not equal increased")
	}
}
