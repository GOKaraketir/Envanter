package main

import (
	"fmt"
	"github.com/GOKaraketir/Envanter/backend"
)

func main() {
	backend.Hello()

	inventory, err := backend.Initialize()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	_, err = inventory.CreateProduct("0", "First Product", 1, 0)
	if err != nil {
		panic(err)
	}
	_, err = inventory.CreateProduct("1", "Second Product", 0.1, 0)
	if err != nil {
		panic(err)
	}

	sell := backend.Sell{}

	product_1, err := inventory.GetProduct("0", "First Product")
	if err != nil {
		panic(err)
	}

	product_2, err := inventory.GetProduct("1", "Second Product")
	if err != nil {
		panic(err)
	}

	sellEntry_1 := backend.SellEntry{
		Product: product_1,
		Count:   5,
	}

	sell.AddSellEntry(sellEntry_1)

	sellEntry_2 := backend.SellEntry{
		Product: product_2,
		Count:   9,
	}

	sell.AddSellEntry(sellEntry_2)

	inventory.Create(&sell)

	var getSell backend.Sell

	inventory.Preload("Entries").First(&getSell)

	fmt.Printf("%v\n", getSell)

	fmt.Println("Sum: ", getSell.Sum())

	err = inventory.DoneSell(&getSell)
	if err != nil {
		panic(err)
	}

	//inventory.Delete()

}
