package program

import (
	"fmt"
	"github.com/GOKaraketir/Envanter/backend"
	"github.com/GOKaraketir/Envanter/ui"
	"github.com/therecipe/qt/widgets"
	"time"
	"unsafe"
)

const (
	SALES_LIST_TIME = iota
	SALES_LIST_TOTAL_PRICE_COL
)

var salesListColNames = []string{
	"Tarih",
	"Toplam Fiyat",
}
var salesListColWidths = []int{
	200, //Tarih
	75,  //Toplam Fiyat
}

func SalesListCreateRow(sell backend.Sell) (items []*widgets.QTableWidgetItem) {
	items = append(items, widgets.NewQTableWidgetItem2(fmt.Sprint(sell.Time.Format(time.ANSIC)), 0))
	items = append(items, widgets.NewQTableWidgetItem2(fmt.Sprintf("%.2f₺", sell.Sum()), 0))
	return
}

func CreateSalesList(p widgets.QWidget_ITF) unsafe.Pointer {
	salesListView := ui.NewForm(p)

	SetLabels(salesListView.SalesListTableWidget, salesListColNames, salesListColWidths)

	list := refreshSalesList(salesListView.SalesListTableWidget)

	salesListView.UndoSalePushButton.ConnectClicked(func(checked bool) {
		rows := salesListView.SalesListTableWidget.SelectionModel().SelectedRows(0)
		if len(rows) == 0 {
			ShowInfo("Hata", "Seçili Satış Yok")
			return
		} else {
			row := rows[0].Row()
			err := Inventory.UndoSell(list[row].ID)
			if err != nil {
				ShowInfo("Hata", err.Error())
				return
			}
			list = refreshSalesList(salesListView.SalesListTableWidget)
		}
	})

	return salesListView.Pointer()
}

func refreshSalesList(table *widgets.QTableWidget) []backend.Sell {
	table.SetRowCount(0)
	sells := Inventory.GetAllSells()
	for row, sell := range sells {
		AddRow(table)
		rowElem := SalesListCreateRow(sell)
		for col, item := range rowElem {
			table.SetItem(row, col, item)
		}
	}
	return sells
}
