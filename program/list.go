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

const (
	ID_COL = iota
	BORCODE_COL
	NAME_COL
	PRICE_COL
	COUNT_COL
)

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

func refreshProductLıst(table *widgets.QTableWidget) {
	table.SetRowCount(0)
	products := Inventory.GetAllProducts()
	for row, product := range products {
		addRow(table)
		rowElem := createRow(product)
		for col, item := range rowElem {
			table.SetItem(row, col, item)
		}
	}
}

func CreateList(p widgets.QWidget_ITF) *ui.List {
	listView := ui.NewList(p)

	setLabels(listView.TableWidget)

	refreshProductLıst(listView.TableWidget)

	listView.DeleteProductButtton.ConnectClicked(func(checked bool) {
		rows := listView.TableWidget.SelectionModel().SelectedRows(0)
		if len(rows) == 0 {
			ShowInfo("Hata", "Seçili Ürün Yok")
			return
		} else {
			for _, rowItem := range rows {
				row := rowItem.Row()
				barcode := listView.TableWidget.Item(row, BORCODE_COL).Text()
				name := listView.TableWidget.Item(row, NAME_COL).Text()
				_, err = Inventory.DeleteProduct(backend.Tag{
					Barcode: barcode,
					Name:    name,
				})
				if err != nil {
					ShowInfo("Hata", barcode+" "+name+" "+err.Error())
				} else {
					ShowInfo("Ürün Silindi", barcode+" "+name)
				}
			}
			refreshProductLıst(listView.TableWidget)
		}
	})

	return listView
}
