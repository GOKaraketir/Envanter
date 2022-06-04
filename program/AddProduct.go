package program

import (
	"fmt"
	"github.com/GOKaraketir/Envanter/backend"
	"github.com/GOKaraketir/Envanter/ui"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func setCheckBoxNoBarcode(addProduct *ui.AddProduct) func(bool) {
	return func(checked bool) {
		addProduct.BarcodeLineEdit.SetDisabled(checked)
		if checked {
			addProduct.BarcodeLineEdit.SetPlaceholderText("Barkodsuz Ürün")
		} else {
			addProduct.BarcodeLineEdit.SetPlaceholderText("")
		}
	}
}

func addOrUpdateFuncGenerator(addProduct *ui.AddProduct, oldProduct *backend.Product) func(bool) {

	addProduct.NoBarcodeCheckBox.ConnectClicked(setCheckBoxNoBarcode(addProduct))

	if oldProduct != nil {
		addProduct.SetWindowTitle(fmt.Sprintf("Ürün Güncelle ID:%d", oldProduct.ID))
		addProduct.AddProductButton.SetText(fmt.Sprintf("Ürün Güncelle: ID:%d", oldProduct.ID))

		addProduct.BarcodeLineEdit.SetText(oldProduct.Barcode)

		addProduct.NoBarcodeCheckBox.SetChecked(oldProduct.Barcode == backend.NO_BARCODE)
		setCheckBoxNoBarcode(addProduct)(oldProduct.Barcode == backend.NO_BARCODE)

		addProduct.NameLineEdit.SetText(oldProduct.Name)

		addProduct.PriceDoubleSpinBox.SetValue(float64(oldProduct.Price))
		addProduct.CountSpinBox.SetValue(oldProduct.Stock.Count)
	} else {
		addProduct.SetWindowTitle("Ürün Ekle")
		addProduct.AddProductButton.SetText("Ürün Ekle")
	}

	return func(bool) {

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

		fmt.Println("old product: ", oldProduct)

		if oldProduct != nil {
			copyOld := *oldProduct
			copyOld.Stock.Count = addProduct.CountSpinBox.Value()
			copyOld.Name = addProduct.NameLineEdit.Text()
			copyOld.Barcode = barcode
			copyOld.Price = backend.Price(addProduct.PriceDoubleSpinBox.Value())
			_, err := Inventory.UpdateProduct(oldProduct.Tag, copyOld)
			if err != nil {
				ShowInfo("Hata", err.Error())
			} else {
				addProduct.Close()
			}
		} else {
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
		}
	}
}

func CreateAddProduct(p widgets.QWidget_ITF) *ui.AddProduct {
	addProduct := ui.NewAddProduct(p)

	addProduct.AddProductButton.ConnectClicked(addOrUpdateFuncGenerator(addProduct, nil))

	return addProduct
}

func CreateUpdateProduct(tag backend.Tag, p widgets.QWidget_ITF) *ui.AddProduct {
	update := ui.NewAddProduct(p)

	product, err := Inventory.GetProduct(tag)
	if err != nil {
		ShowInfo("Hata", err.Error())
		return nil
	}

	update.AddProductButton.ConnectClicked(addOrUpdateFuncGenerator(update, product))

	return update
}
