package program

import (
	"github.com/GOKaraketir/Envanter/ui"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"unsafe"
)

func buttonConnectorWithDisable(button *widgets.QPushButton, creator func(itf widgets.QWidget_ITF) unsafe.Pointer) {
	button.ConnectClicked(func(checked bool) {
		button.SetDisabled(true)
		widget := widgets.QWidget{}
		widget.SetPointer(creator(nil))
		widget.Show()
		widget.ConnectCloseEvent(func(event *gui.QCloseEvent) {
			button.SetDisabled(false)
		})
	})
}

func CreateMainPage(p widgets.QWidget_ITF) *ui.MainPage {
	mainPage := ui.NewMainPage(p)

	buttonConnectorWithDisable(mainPage.AddProductPushButton, CreateAddProduct)
	buttonConnectorWithDisable(mainPage.ListProducts, CreateList)
	buttonConnectorWithDisable(mainPage.SalePanelPushButton, CreateSale)

	return mainPage
}
