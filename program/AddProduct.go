package program

import (
	"github.com/GOKaraketir/Envanter/backend"
	"github.com/GOKaraketir/Envanter/ui"
	"github.com/therecipe/qt/widgets"
)

func CreateAddProduct(p widgets.QWidget_ITF) *ui.AddProduct {
	addProduct := ui.NewAddProduct(p)

	addProduct.AddProductButton.ConnectClicked(func(bool) {
		_, err := Inventory.CreateProduct(backend.Tag{
			Barcode: addProduct.BarcodeLineEdit.Text(),
			Name:    addProduct.NameLineEdit.Text(),
		}, backend.Price(addProduct.PriceDoubleSpinBox.Value()),
			addProduct.CountSpinBox.Value())
		if err != nil {
			ShowInfo("Hata", err.Error())
		} else {
			addProduct.Close()
		}
	})

	addProduct.NoBarcodeCheckBox.ConnectClicked(func(checked bool) {
		addProduct.BarcodeLineEdit.SetDisabled(checked)
		if checked {
			addProduct.BarcodeLineEdit.SetPlaceholderText("Barkodsuz Ürün")
		} else {
			addProduct.BarcodeLineEdit.SetPlaceholderText("")
		}
	})

	return addProduct
}

func CreateDeleteProduct(tag backend.Tag, p widgets.QWidget_ITF) *ui.AddProduct {
	deleteProduct := ui.NewAddProduct(p)

	return deleteProduct
}
