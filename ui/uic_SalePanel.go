package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type __salepanel struct{}

func (*__salepanel) init() {}

type SalePanel struct {
	*__salepanel
	*widgets.QWidget
	VerticalLayout_2         *widgets.QVBoxLayout
	HorizontalLayout_5       *widgets.QHBoxLayout
	VerticalLayout           *widgets.QVBoxLayout
	SelectProductComboBox    *widgets.QComboBox
	HorizontalLayout         *widgets.QHBoxLayout
	Label                    *widgets.QLabel
	BarcodeLineEdit          *widgets.QLineEdit
	HorizontalLayout_2       *widgets.QHBoxLayout
	Label_2                  *widgets.QLabel
	CountSpinBox             *widgets.QSpinBox
	AddPushButton            *widgets.QPushButton
	VerticalSpacer_2         *widgets.QSpacerItem
	RemoveSelectedPushButton *widgets.QPushButton
	VerticalSpacer           *widgets.QSpacerItem
	HorizontalLayout_4       *widgets.QHBoxLayout
	Label_3                  *widgets.QLabel
	TotalPiceLineEdit        *widgets.QLineEdit
	TableWidget              *widgets.QTableWidget
	HorizontalLayout_3       *widgets.QHBoxLayout
	PushButton_2             *widgets.QPushButton
	HorizontalSpacer         *widgets.QSpacerItem
	PushButton_3             *widgets.QPushButton
}

func NewSalePanel(p widgets.QWidget_ITF) *SalePanel {
	var par *widgets.QWidget
	if p != nil {
		par = p.QWidget_PTR()
	}
	w := &SalePanel{QWidget: widgets.NewQWidget(par, 0)}
	w.setupUI()
	w.init()
	return w
}
func (w *SalePanel) setupUI() {
	if w.ObjectName() == "" {
		w.SetObjectName("SalePanel")
	}
	w.Resize2(1200, 429)
	w.VerticalLayout_2 = widgets.NewQVBoxLayout2(w)
	w.VerticalLayout_2.SetObjectName("verticalLayout_2")
	w.HorizontalLayout_5 = widgets.NewQHBoxLayout()
	w.HorizontalLayout_5.SetObjectName("horizontalLayout_5")
	w.VerticalLayout = widgets.NewQVBoxLayout()
	w.VerticalLayout.SetObjectName("verticalLayout")
	w.SelectProductComboBox = widgets.NewQComboBox(w)
	w.SelectProductComboBox.SetObjectName("selectProductComboBox")
	w.SelectProductComboBox.SetMaximumSize(core.NewQSize2(16777215, 16777215))
	w.VerticalLayout.QLayout.AddWidget(w.SelectProductComboBox)
	w.HorizontalLayout = widgets.NewQHBoxLayout()
	w.HorizontalLayout.SetObjectName("horizontalLayout")
	w.Label = widgets.NewQLabel(w, 0)
	w.Label.SetObjectName("label")
	w.HorizontalLayout.QLayout.AddWidget(w.Label)
	w.BarcodeLineEdit = widgets.NewQLineEdit(w)
	w.BarcodeLineEdit.SetObjectName("barcodeLineEdit")
	w.HorizontalLayout.QLayout.AddWidget(w.BarcodeLineEdit)
	w.VerticalLayout.AddLayout(w.HorizontalLayout, 0)
	w.HorizontalLayout_2 = widgets.NewQHBoxLayout()
	w.HorizontalLayout_2.SetObjectName("horizontalLayout_2")
	w.Label_2 = widgets.NewQLabel(w, 0)
	w.Label_2.SetObjectName("label_2")
	w.HorizontalLayout_2.QLayout.AddWidget(w.Label_2)
	w.CountSpinBox = widgets.NewQSpinBox(w)
	w.CountSpinBox.SetObjectName("countSpinBox")
	w.CountSpinBox.SetMaximum(999999999)
	w.CountSpinBox.SetValue(1)
	w.HorizontalLayout_2.QLayout.AddWidget(w.CountSpinBox)
	w.VerticalLayout.AddLayout(w.HorizontalLayout_2, 0)
	w.AddPushButton = widgets.NewQPushButton(w)
	w.AddPushButton.SetObjectName("AddPushButton")
	w.VerticalLayout.QLayout.AddWidget(w.AddPushButton)
	w.VerticalSpacer_2 = widgets.NewQSpacerItem(20, 40, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	w.VerticalLayout.AddItem(w.VerticalSpacer_2)
	w.RemoveSelectedPushButton = widgets.NewQPushButton(w)
	w.RemoveSelectedPushButton.SetObjectName("removeSelectedPushButton")
	w.VerticalLayout.QLayout.AddWidget(w.RemoveSelectedPushButton)
	w.VerticalSpacer = widgets.NewQSpacerItem(20, 40, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	w.VerticalLayout.AddItem(w.VerticalSpacer)
	w.HorizontalLayout_4 = widgets.NewQHBoxLayout()
	w.HorizontalLayout_4.SetObjectName("horizontalLayout_4")
	w.Label_3 = widgets.NewQLabel(w, 0)
	w.Label_3.SetObjectName("label_3")
	w.HorizontalLayout_4.QLayout.AddWidget(w.Label_3)
	w.TotalPiceLineEdit = widgets.NewQLineEdit(w)
	w.TotalPiceLineEdit.SetObjectName("totalPiceLineEdit")
	w.TotalPiceLineEdit.SetReadOnly(true)
	w.HorizontalLayout_4.QLayout.AddWidget(w.TotalPiceLineEdit)
	w.VerticalLayout.AddLayout(w.HorizontalLayout_4, 0)
	w.HorizontalLayout_5.AddLayout(w.VerticalLayout, 0)
	w.TableWidget = widgets.NewQTableWidget(w)
	w.TableWidget.SetObjectName("tableWidget")
	w.TableWidget.SetMinimumSize(core.NewQSize2(750, 0))
	w.TableWidget.SetSelectionMode(widgets.QAbstractItemView__SingleSelection)
	w.TableWidget.SetSelectionBehavior(widgets.QAbstractItemView__SelectRows)
	w.HorizontalLayout_5.QLayout.AddWidget(w.TableWidget)
	w.VerticalLayout_2.AddLayout(w.HorizontalLayout_5, 0)
	w.HorizontalLayout_3 = widgets.NewQHBoxLayout()
	w.HorizontalLayout_3.SetObjectName("horizontalLayout_3")
	w.PushButton_2 = widgets.NewQPushButton(w)
	w.PushButton_2.SetObjectName("pushButton_2")
	w.HorizontalLayout_3.QLayout.AddWidget(w.PushButton_2)
	w.HorizontalSpacer = widgets.NewQSpacerItem(40, 20, widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Minimum)
	w.HorizontalLayout_3.AddItem(w.HorizontalSpacer)
	w.PushButton_3 = widgets.NewQPushButton(w)
	w.PushButton_3.SetObjectName("pushButton_3")
	w.HorizontalLayout_3.QLayout.AddWidget(w.PushButton_3)
	w.VerticalLayout_2.AddLayout(w.HorizontalLayout_3, 0)
	w.retranslateUi()
	core.QMetaObject_ConnectSlotsByName(w)

}
func (w *SalePanel) retranslateUi() {
	w.SetWindowTitle(core.QCoreApplication_Translate("SalePanel", "Form", "", 0))
	w.Label.SetText(core.QCoreApplication_Translate("SalePanel", "Barkod", "", 0))
	w.Label_2.SetText(core.QCoreApplication_Translate("SalePanel", "Adet", "", 0))
	w.AddPushButton.SetText(core.QCoreApplication_Translate("SalePanel", "Ekle", "", 0))
	w.RemoveSelectedPushButton.SetText(core.QCoreApplication_Translate("SalePanel", "Se\303\247ili \303\234r\303\274n\303\274 \303\207\304\261kar", "", 0))
	w.Label_3.SetText(core.QCoreApplication_Translate("SalePanel", "Toplam Fiyat: ", "", 0))
	w.PushButton_2.SetText(core.QCoreApplication_Translate("SalePanel", "Sat\304\261\305\237", "", 0))
	w.PushButton_3.SetText(core.QCoreApplication_Translate("SalePanel", "\304\260ptal Et", "", 0))

}
