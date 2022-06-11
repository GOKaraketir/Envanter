package program

import (
	"fmt"
	"github.com/GOKaraketir/Envanter/backend"
	"github.com/GOKaraketir/Envanter/ui"
	"github.com/therecipe/qt/widgets"
	"unsafe"
)

const (
	SELL_BARCODE_COL = iota
	SELL_NAME_COL
	SELL_UNIT_PRICE_COL
	SELL_COUNT_COL
	SELL_TOTAL_PRICE_COL
)

var saleColNames = []string{
	"Barkod",
	"Ürün Adı",
	"Birim Fiyatı",
	"Adet",
	"Fiyat",
}
var saleColWidths = []int{
	125, //Barkod
	300, //Urun Adi
	75,  //Birim Fiyat
	75,  //Adet
	75,  //Toplam Fiyat
}

func saleCreateRow(entry backend.SellEntry) (items []*widgets.QTableWidgetItem) {
	items = append(items, widgets.NewQTableWidgetItem2(fmt.Sprint(entry.Barcode), 0))
	items = append(items, widgets.NewQTableWidgetItem2(fmt.Sprint(entry.Name), 0))
	items = append(items, widgets.NewQTableWidgetItem2(fmt.Sprintf("%.2f", entry.Price), 0))
	items = append(items, widgets.NewQTableWidgetItem2(fmt.Sprint(entry.Count), 0))
	items = append(items, widgets.NewQTableWidgetItem2(fmt.Sprintf("%.2f", entry.Price*backend.Price(entry.Count)), 0))
	return
}

func CreateSale(itf widgets.QWidget_ITF) unsafe.Pointer {
	sale := ui.NewSalePanel(itf)

	sell := backend.NewSell()

	table := sale.TableWidget

	SetLabels(table, saleColNames, saleColWidths)

	sale.SelectProductComboBox.SetInsertPolicy(widgets.QComboBox__NoInsert)

	products := Inventory.GetAllProducts()
	sale.SelectProductComboBox.InsertItems(0, []string{"Ürün Seçin"})
	sale.SelectProductComboBox.InsertItems(1, products.GetAllNames())

	findIndex := func(product *backend.Product) int {
		for i, p := range products {
			if p.ID == product.ID {
				return i
			}
		}
		return -1
	}

	sale.BarcodeLineEdit.ConnectTextChanged(func(text string) {

	})

	sale.SelectProductComboBox.ConnectCurrentIndexChanged(func(index int) {
		if index != 0 {
			sale.BarcodeLineEdit.SetText(products[index-1].Barcode)
		}
	})

	sale.BarcodeLineEdit.ConnectReturnPressed(func() {
		text := sale.BarcodeLineEdit.Text()
		product, err := Inventory.FindProductWithBarcode(text)
		if err != nil {
			sale.SelectProductComboBox.SetCurrentIndex(0)
		} else {
			index := findIndex(product)
			fmt.Println("index", index)
			if index == -1 {
				sale.SelectProductComboBox.SetCurrentIndex(0)
			} else {
				sale.SelectProductComboBox.SetCurrentIndex(index + 1)
			}
		}
	})

	sale.AddPushButton.ConnectClicked(func(checked bool) {
		index := sale.SelectProductComboBox.CurrentIndex()
		if index == 0 {
			ShowInfo("Hata", "Ürün Bulunamadı")
			return
		}
		product := products[index-1]

		sell.AddSellEntry(backend.CreateSellEntry(product, sale.CountSpinBox.Value()))
		refreshSaleList(table, sell, sale.TotalPiceLineEdit)

		sale.CountSpinBox.SetValue(1)
	})

	sale.RemoveSelectedPushButton.ConnectClicked(func(checked bool) {
		rows := table.SelectionModel().SelectedRows(0)
		if len(rows) == 0 {
			ShowInfo("HATA", "Seçili Ürün Yok")
			return
		}

		row := rows[0].Row()

		sell.Entries = backend.RemoveIndexWithOrder(sell.Entries, row)
		refreshSaleList(table, sell, sale.TotalPiceLineEdit)
	})

	sale.CommitSalePushButton.ConnectClicked(func(checked bool) {
		err := Inventory.CommitSell(sell)
		if err != nil {
			ShowInfo("HATA", err.Error())
			return
		} else {
			sale.Close()
		}

	})

	sale.CancelPushButton.ConnectClicked(func(checked bool) {
		sale.Close()
	})

	return sale.Pointer()
}

func refreshSaleList(table *widgets.QTableWidget, sell *backend.Sell, totalPriceWidget *widgets.QLineEdit) {
	table.SetRowCount(0)

	for row, entry := range sell.Entries {
		AddRow(table)
		rowElem := saleCreateRow(entry)
		for col, item := range rowElem {
			table.SetItem(row, col, item)
		}
	}

	sum := sell.Sum()
	totalPriceWidget.SetText(fmt.Sprintf("%.2f₺", sum))
}
