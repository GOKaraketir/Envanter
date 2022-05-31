package program

import (
	"fmt"
	"github.com/GOKaraketir/Envanter/backend"
	"github.com/GOKaraketir/Envanter/ui"
	"github.com/therecipe/qt/widgets"
)

func addRow(table *widgets.QTableWidget) {
	table.InsertRow(table.RowCount())
}

var colNames = []string{
	"ID",
	"Barkod",
	"Ürün Adı",
	"Fiyat",
	"Stok Adedi",
}
var colWidths = []int{
	50,  //ID
	125, //Barkod
	300, //Urun Adi
	75,  //Fiyat
	75,  //Stok Adedi
}

func createRow(product backend.Product) (items []*widgets.QTableWidgetItem) {
	items = append(items, widgets.NewQTableWidgetItem2(fmt.Sprint(product.ID), 0))
	items = append(items, widgets.NewQTableWidgetItem2(fmt.Sprint(product.Barcode), 0))
	items = append(items, widgets.NewQTableWidgetItem2(fmt.Sprint(product.Name), 0))
	items = append(items, widgets.NewQTableWidgetItem2(fmt.Sprintf("%.2f", product.Price), 0))
	items = append(items, widgets.NewQTableWidgetItem2(fmt.Sprint(product.Stock.Count), 0))
	return
}

func setLabels(table *widgets.QTableWidget) {
	for i, name := range colNames {
		table.InsertColumn(i)
		table.SetHorizontalHeaderItem(i, widgets.NewQTableWidgetItem2(name, 0))
		table.SetColumnWidth(i, colWidths[i])
	}
}

//func addColumn(table *widgets.QTableWidget) {
//
//	fmt.Println("inserted: ", i+1)
//}

func CreateList(p widgets.QWidget_ITF) *ui.Form {
	asd := ui.NewForm(p)

	setLabels(asd.TableWidget)

	products := Inventory.GetAllProducts()
	for row, product := range products {
		addRow(asd.TableWidget)
		rowElem := createRow(product)
		for col, item := range rowElem {
			asd.TableWidget.SetItem(row, col, item)
		}
	}

	return asd
}
