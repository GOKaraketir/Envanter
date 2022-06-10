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

func testProductToProduct(product TestProduct) backend.Product {
	return backend.Product{
		Tag: backend.Tag{
			Barcode: product.Barcode,
			Name:    product.Name,
		},
		Price: product.Price,
	}
}

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

func TestFindWithBarcode(t *testing.T) {
	inventory := createDB(t)

	product, err := inventory.FindProductWithBarcode("1111111111111")
	if err != nil {
		t.Fatal(err)
	} else if product.Name != "Ürün 1" {
		t.Fatal("ERRR")
	}

	_, err = inventory.FindProductWithBarcode("123")
	if err == nil {
		t.Fatal("ErEr")
	}

}

func TestUpdate(t *testing.T) {
	inventory := createDB(t)
	var err error
	var prod backend.Product

	inventory.Preload("Stock").First(&prod)

	tag := prod.Tag

	product, err := inventory.GetProduct(tag)
	if err != nil {
		t.Fatal(err)
	}

	product.Price = 22
	product.Name = "ASDASD"

	newPro, err := inventory.UpdateProduct(tag, prod)
	if err != nil {
		t.Fatal(err)
	}

	_, err = inventory.UpdateStock(newPro.Tag, 444)
	if err != nil {
		t.Fatal(err)
	}
}
