package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type __addproduct struct{}

func (*__addproduct) init() {}

type AddProduct struct {
	*__addproduct
	*widgets.QWidget
	VerticalLayout     *widgets.QVBoxLayout
	HorizontalLayout_2 *widgets.QHBoxLayout
	Label_2            *widgets.QLabel
	HorizontalSpacer_2 *widgets.QSpacerItem
	NameLineEdit       *widgets.QLineEdit
	HorizontalLayout   *widgets.QHBoxLayout
	Label              *widgets.QLabel
	HorizontalSpacer   *widgets.QSpacerItem
	NoBarcodeCheckBox  *widgets.QCheckBox
	HorizontalSpacer_6 *widgets.QSpacerItem
	BarcodeLineEdit    *widgets.QLineEdit
	HorizontalLayout_4 *widgets.QHBoxLayout
	Label_4            *widgets.QLabel
	HorizontalSpacer_4 *widgets.QSpacerItem
	PriceDoubleSpinBox *widgets.QDoubleSpinBox
	HorizontalLayout_3 *widgets.QHBoxLayout
	Label_3            *widgets.QLabel
	HorizontalSpacer_3 *widgets.QSpacerItem
	CountSpinBox       *widgets.QSpinBox
	HorizontalLayout_5 *widgets.QHBoxLayout
	HorizontalSpacer_5 *widgets.QSpacerItem
	AddProductButton   *widgets.QPushButton
	HorizontalSpacer_7 *widgets.QSpacerItem
}

func NewAddProduct(p widgets.QWidget_ITF) *AddProduct {
	var par *widgets.QWidget
	if p != nil {
		par = p.QWidget_PTR()
	}
	w := &AddProduct{QWidget: widgets.NewQWidget(par, 0)}
	w.setupUI()
	w.init()
	return w
}
func (w *AddProduct) setupUI() {
	if w.ObjectName() == "" {
		w.SetObjectName("AddProduct")
	}
	w.Resize2(586, 333)
	w.VerticalLayout = widgets.NewQVBoxLayout2(w)
	w.VerticalLayout.SetObjectName("verticalLayout")
	w.HorizontalLayout_2 = widgets.NewQHBoxLayout()
	w.HorizontalLayout_2.SetObjectName("horizontalLayout_2")
	w.Label_2 = widgets.NewQLabel(w, 0)
	w.Label_2.SetObjectName("label_2")
	w.HorizontalLayout_2.QLayout.AddWidget(w.Label_2)
	w.HorizontalSpacer_2 = widgets.NewQSpacerItem(40, 20, widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Minimum)
	w.HorizontalLayout_2.AddItem(w.HorizontalSpacer_2)
	w.NameLineEdit = widgets.NewQLineEdit(w)
	w.NameLineEdit.SetObjectName("nameLineEdit")
	w.NameLineEdit.SetMinimumSize(core.NewQSize2(400, 0))
	w.HorizontalLayout_2.QLayout.AddWidget(w.NameLineEdit)
	w.VerticalLayout.AddLayout(w.HorizontalLayout_2, 0)
	w.HorizontalLayout = widgets.NewQHBoxLayout()
	w.HorizontalLayout.SetObjectName("horizontalLayout")
	w.Label = widgets.NewQLabel(w, 0)
	w.Label.SetObjectName("label")
	w.HorizontalLayout.QLayout.AddWidget(w.Label)
	w.HorizontalSpacer = widgets.NewQSpacerItem(40, 20, widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Minimum)
	w.HorizontalLayout.AddItem(w.HorizontalSpacer)
	w.NoBarcodeCheckBox = widgets.NewQCheckBox(w)
	w.NoBarcodeCheckBox.SetObjectName("noBarcodeCheckBox")
	w.HorizontalLayout.QLayout.AddWidget(w.NoBarcodeCheckBox)
	w.HorizontalSpacer_6 = widgets.NewQSpacerItem(40, 20, widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Minimum)
	w.HorizontalLayout.AddItem(w.HorizontalSpacer_6)
	w.BarcodeLineEdit = widgets.NewQLineEdit(w)
	w.BarcodeLineEdit.SetObjectName("barcodeLineEdit")
	w.BarcodeLineEdit.SetMinimumSize(core.NewQSize2(200, 0))
	w.BarcodeLineEdit.SetMaximumSize(core.NewQSize2(200, 16777215))
	w.BarcodeLineEdit.SetInputMethodHints(core.Qt__ImhNone)
	w.HorizontalLayout.QLayout.AddWidget(w.BarcodeLineEdit)
	w.VerticalLayout.AddLayout(w.HorizontalLayout, 0)
	w.HorizontalLayout_4 = widgets.NewQHBoxLayout()
	w.HorizontalLayout_4.SetObjectName("horizontalLayout_4")
	w.Label_4 = widgets.NewQLabel(w, 0)
	w.Label_4.SetObjectName("label_4")
	w.HorizontalLayout_4.QLayout.AddWidget(w.Label_4)
	w.HorizontalSpacer_4 = widgets.NewQSpacerItem(40, 20, widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Minimum)
	w.HorizontalLayout_4.AddItem(w.HorizontalSpacer_4)
	w.PriceDoubleSpinBox = widgets.NewQDoubleSpinBox(w)
	w.PriceDoubleSpinBox.SetObjectName("priceDoubleSpinBox")
	w.PriceDoubleSpinBox.SetMinimumSize(core.NewQSize2(150, 0))
	w.PriceDoubleSpinBox.SetMaximumSize(core.NewQSize2(150, 16777215))
	w.PriceDoubleSpinBox.SetAlignment(core.Qt__AlignCenter)
	w.PriceDoubleSpinBox.SetMinimum(0.000000000000000)
	w.PriceDoubleSpinBox.SetMaximum(9999.989999999999782)
	w.HorizontalLayout_4.QLayout.AddWidget(w.PriceDoubleSpinBox)
	w.VerticalLayout.AddLayout(w.HorizontalLayout_4, 0)
	w.HorizontalLayout_3 = widgets.NewQHBoxLayout()
	w.HorizontalLayout_3.SetObjectName("horizontalLayout_3")
	w.Label_3 = widgets.NewQLabel(w, 0)
	w.Label_3.SetObjectName("label_3")
	w.HorizontalLayout_3.QLayout.AddWidget(w.Label_3)
	w.HorizontalSpacer_3 = widgets.NewQSpacerItem(40, 20, widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Minimum)
	w.HorizontalLayout_3.AddItem(w.HorizontalSpacer_3)
	w.CountSpinBox = widgets.NewQSpinBox(w)
	w.CountSpinBox.SetObjectName("countSpinBox")
	w.CountSpinBox.SetMinimumSize(core.NewQSize2(150, 0))
	w.CountSpinBox.SetMaximumSize(core.NewQSize2(150, 16777215))
	w.CountSpinBox.SetAlignment(core.Qt__AlignCenter)
	w.CountSpinBox.SetMinimum(-99999)
	w.CountSpinBox.SetMaximum(99999)
	w.HorizontalLayout_3.QLayout.AddWidget(w.CountSpinBox)
	w.VerticalLayout.AddLayout(w.HorizontalLayout_3, 0)
	w.HorizontalLayout_5 = widgets.NewQHBoxLayout()
	w.HorizontalLayout_5.SetObjectName("horizontalLayout_5")
	w.HorizontalSpacer_5 = widgets.NewQSpacerItem(40, 20, widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Minimum)
	w.HorizontalLayout_5.AddItem(w.HorizontalSpacer_5)
	w.AddProductButton = widgets.NewQPushButton(w)
	w.AddProductButton.SetObjectName("addProductButton")
	w.AddProductButton.SetMinimumSize(core.NewQSize2(200, 0))
	w.HorizontalLayout_5.QLayout.AddWidget(w.AddProductButton)
	w.HorizontalSpacer_7 = widgets.NewQSpacerItem(40, 20, widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Minimum)
	w.HorizontalLayout_5.AddItem(w.HorizontalSpacer_7)
	w.VerticalLayout.AddLayout(w.HorizontalLayout_5, 0)
	widgets.QWidget_SetTabOrder(w.NameLineEdit, w.BarcodeLineEdit)
	widgets.QWidget_SetTabOrder(w.BarcodeLineEdit, w.PriceDoubleSpinBox)
	widgets.QWidget_SetTabOrder(w.PriceDoubleSpinBox, w.CountSpinBox)
	widgets.QWidget_SetTabOrder(w.CountSpinBox, w.AddProductButton)
	widgets.QWidget_SetTabOrder(w.AddProductButton, w.NoBarcodeCheckBox)
	w.retranslateUi()
	core.QMetaObject_ConnectSlotsByName(w)

}
func (w *AddProduct) retranslateUi() {
	w.SetWindowTitle(core.QCoreApplication_Translate("AddProduct", "\303\234r\303\274n Ekle", "", 0))
	w.Label_2.SetText(core.QCoreApplication_Translate("AddProduct", "\303\234r\303\274n Ad\304\261:", "", 0))
	w.Label.SetText(core.QCoreApplication_Translate("AddProduct", "Barkod Numaras\304\261:", "", 0))
	w.NoBarcodeCheckBox.SetText(core.QCoreApplication_Translate("AddProduct", "Barkodsuz \303\234r\303\274n", "", 0))
	w.BarcodeLineEdit.SetInputMask(core.QCoreApplication_Translate("AddProduct", "99999999999999999999", "", 0))
	w.BarcodeLineEdit.SetText("")
	w.Label_4.SetText(core.QCoreApplication_Translate("AddProduct", "\303\234r\303\274n Fiyat\304\261:", "", 0))
	w.PriceDoubleSpinBox.SetSuffix(core.QCoreApplication_Translate("AddProduct", "\342\202\272", "", 0))
	w.Label_3.SetText(core.QCoreApplication_Translate("AddProduct", "\303\234r\303\274n Adedi:", "", 0))
	w.AddProductButton.SetText(core.QCoreApplication_Translate("AddProduct", "\303\234r\303\274n Ekle", "", 0))

}
