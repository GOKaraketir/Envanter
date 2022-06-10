package program

import (
	"fmt"
	"github.com/GOKaraketir/Envanter/backend"
	"github.com/GOKaraketir/Envanter/ui"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"unsafe"
)

func AddRow(table *widgets.QTableWidget) {
	table.InsertRow(table.RowCount())
}

const (
	LIST_ID_COL = iota
	LIST_BARCODE_COL
	LIST_NAME_COL
	LIST_PRICE_COL
	LIST_COUNT_COL
)

var listColnames = []string{
	"ID",
	"Barkod",
	"Ürün Adı",
	"Fiyat",
	"Stok Adedi",
}
var listColWidths = []int{
	50,  //ID
	125, //Barkod
	300, //Urun Adi
	75,  //Fiyat
	75,  //Stok Adedi
}

func listCreateRow(product backend.Product) (items []*widgets.QTableWidgetItem) {
	items = append(items, widgets.NewQTableWidgetItem2(fmt.Sprint(product.ID), 0))
	items = append(items, widgets.NewQTableWidgetItem2(fmt.Sprint(product.Barcode), 0))
	items = append(items, widgets.NewQTableWidgetItem2(fmt.Sprint(product.Name), 0))
	items = append(items, widgets.NewQTableWidgetItem2(fmt.Sprintf("%.2f", product.Price), 0))
	items = append(items, widgets.NewQTableWidgetItem2(fmt.Sprint(product.Stock.Count), 0))
	return
}

func SetLabels(table *widgets.QTableWidget, colN []string, colW []int) {
	for i, name := range colN {
		table.InsertColumn(i)
		table.SetHorizontalHeaderItem(i, widgets.NewQTableWidgetItem2(name, 0))
		table.SetColumnWidth(i, colW[i])
	}
}

func refreshProductList(table *widgets.QTableWidget) {
	table.SetRowCount(0)
	products := Inventory.GetAllProducts()
	for row, product := range products {
		AddRow(table)
		rowElem := listCreateRow(product)
		for col, item := range rowElem {
			table.SetItem(row, col, item)
		}
	}
}

func CreateList(p widgets.QWidget_ITF) unsafe.Pointer {
	listView := ui.NewList(p)

	SetLabels(listView.TableWidget, listColnames, listColWidths)

	refreshProductList(listView.TableWidget)

	listView.DeleteProductButtton.ConnectClicked(func(checked bool) {
		rows := listView.TableWidget.SelectionModel().SelectedRows(0)
		if len(rows) == 0 {
			ShowInfo("Hata", "Seçili Ürün Yok")
			return
		} else {
			for _, rowItem := range rows {
				row := rowItem.Row()
				barcode := listView.TableWidget.Item(row, LIST_BARCODE_COL).Text()
				name := listView.TableWidget.Item(row, LIST_NAME_COL).Text()
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
			refreshProductList(listView.TableWidget)
		}
	})

	listView.UpdateProductButton.ConnectClicked(func(checked bool) {
		rows := listView.TableWidget.SelectionModel().SelectedRows(0)
		if len(rows) == 0 {
			ShowInfo("Hata", "Seçili Ürün Yok")
			return
		} else if len(rows) > 1 {
			ShowInfo("Hata", "Birden Fazla Ürün Seçili, güncellemek için bir ürün seçin")
			return
		} else {
			rowItem := rows[0]
			row := rowItem.Row()
			barcode := listView.TableWidget.Item(row, LIST_BARCODE_COL).Text()
			name := listView.TableWidget.Item(row, LIST_NAME_COL).Text()

			updateView := widgets.QWidget{}
			updateView.SetPointer(CreateUpdateProduct(backend.Tag{
				Barcode: barcode,
				Name:    name,
			}, nil))

			updateView.Show()

			updateView.ConnectCloseEvent(func(event *gui.QCloseEvent) {
				refreshProductList(listView.TableWidget)
			})
		}
	})

	return listView.Pointer()
}
