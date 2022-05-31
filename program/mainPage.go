package program

import (
	"github.com/GOKaraketir/Envanter/ui"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

func CreateMainPage(p widgets.QWidget_ITF) *ui.MainPage {
	mainPage := ui.NewMainPage(p)

	mainPage.AddProductPushButton.ConnectClicked(func(checked bool) {
		mainPage.AddProductPushButton.SetDisabled(true)
		addProduct := CreateAddProduct(nil)
		addProduct.Show()
		addProduct.ConnectCloseEvent(func(event *gui.QCloseEvent) {
			mainPage.AddProductPushButton.SetDisabled(false)
		})
	})

	mainPage.ListProducts.ConnectClicked(func(checked bool) {
		mainPage.ListProducts.SetDisabled(true)
		addProduct := CreateList(nil)
		addProduct.Show()
		addProduct.ConnectCloseEvent(func(event *gui.QCloseEvent) {
			mainPage.ListProducts.SetDisabled(false)
		})
	})

	return mainPage
}
