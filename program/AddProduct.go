package program

import (
	"github.com/GOKaraketir/Envanter/backend"
	"github.com/GOKaraketir/Envanter/ui"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func CreateAddProduct(p widgets.QWidget_ITF) *ui.AddProduct {
	addProduct := ui.NewAddProduct(p)

	addProduct.AddProductButton.ConnectClicked(func(bool) {

		barcode := addProduct.BarcodeLineEdit.Text()

		noBarcode := addProduct.NoBarcodeCheckBox.IsChecked()

		if barcode == "" && noBarcode {
			barcode = "0000000000000"
		}

		re := core.NewQRegularExpression()
		re.SetPattern("^[0-9]+$")
		match := re.Match(barcode, 0, core.QRegularExpression__NormalMatch, core.QRegularExpression__NoMatchOption)

		hasMatch := match.HasMatch()

		if !hasMatch {
			ShowInfo("Hata", "Hatalı Barkod, barkod sadece rakam içerebilir (0-9) ve en az 1 karakter uzunluğunda olmalıdır")
			return
		}

		_, err := Inventory.CreateProduct(backend.Tag{
			Barcode: barcode,
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
