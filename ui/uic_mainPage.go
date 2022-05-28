package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type __mainpage struct{}

func (*__mainpage) init() {}

type MainPage struct {
	*__mainpage
	*widgets.QWidget
	VerticalLayout          *widgets.QVBoxLayout
	VerticalSpacer          *widgets.QSpacerItem
	HorizontalLayout        *widgets.QHBoxLayout
	SalePanelPushButton     *widgets.QPushButton
	AddProductPushButton    *widgets.QPushButton
	RemoveProductPushButton *widgets.QPushButton
	EditProductPushButton   *widgets.QPushButton
	PushButton_5            *widgets.QPushButton
	PushButton_4            *widgets.QPushButton
}

func NewMainPage(p widgets.QWidget_ITF) *MainPage {
	var par *widgets.QWidget
	if p != nil {
		par = p.QWidget_PTR()
	}
	w := &MainPage{QWidget: widgets.NewQWidget(par, 0)}
	w.setupUI()
	w.init()
	return w
}
func (w *MainPage) setupUI() {
	if w.ObjectName() == "" {
		w.SetObjectName("MainPage")
	}
	w.Resize2(996, 424)
	w.VerticalLayout = widgets.NewQVBoxLayout2(w)
	w.VerticalLayout.SetObjectName("verticalLayout")
	w.VerticalSpacer = widgets.NewQSpacerItem(20, 312, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	w.VerticalLayout.AddItem(w.VerticalSpacer)
	w.HorizontalLayout = widgets.NewQHBoxLayout()
	w.HorizontalLayout.SetObjectName("horizontalLayout")
	w.SalePanelPushButton = widgets.NewQPushButton(w)
	w.SalePanelPushButton.SetObjectName("salePanelPushButton")
	w.SalePanelPushButton.SetMinimumSize(core.NewQSize2(150, 75))
	w.HorizontalLayout.QLayout.AddWidget(w.SalePanelPushButton)
	w.AddProductPushButton = widgets.NewQPushButton(w)
	w.AddProductPushButton.SetObjectName("addProductPushButton")
	w.AddProductPushButton.SetMinimumSize(core.NewQSize2(150, 75))
	w.HorizontalLayout.QLayout.AddWidget(w.AddProductPushButton)
	w.RemoveProductPushButton = widgets.NewQPushButton(w)
	w.RemoveProductPushButton.SetObjectName("removeProductPushButton")
	w.RemoveProductPushButton.SetMinimumSize(core.NewQSize2(150, 75))
	w.HorizontalLayout.QLayout.AddWidget(w.RemoveProductPushButton)
	w.EditProductPushButton = widgets.NewQPushButton(w)
	w.EditProductPushButton.SetObjectName("editProductPushButton")
	w.EditProductPushButton.SetMinimumSize(core.NewQSize2(150, 75))
	w.HorizontalLayout.QLayout.AddWidget(w.EditProductPushButton)
	w.PushButton_5 = widgets.NewQPushButton(w)
	w.PushButton_5.SetObjectName("pushButton_5")
	w.PushButton_5.SetMinimumSize(core.NewQSize2(150, 75))
	w.HorizontalLayout.QLayout.AddWidget(w.PushButton_5)
	w.PushButton_4 = widgets.NewQPushButton(w)
	w.PushButton_4.SetObjectName("pushButton_4")
	w.PushButton_4.SetMinimumSize(core.NewQSize2(150, 75))
	w.HorizontalLayout.QLayout.AddWidget(w.PushButton_4)
	w.VerticalLayout.AddLayout(w.HorizontalLayout, 0)
	w.retranslateUi()
	core.QMetaObject_ConnectSlotsByName(w)

}
func (w *MainPage) retranslateUi() {
	w.SetWindowTitle(core.QCoreApplication_Translate("MainPage", "Form", "", 0))
	w.SalePanelPushButton.SetText(core.QCoreApplication_Translate("MainPage", "Sat\304\261\305\237 Paneli", "", 0))
	w.AddProductPushButton.SetText(core.QCoreApplication_Translate("MainPage", "\303\234r\303\274n Ekle", "", 0))
	w.RemoveProductPushButton.SetText(core.QCoreApplication_Translate("MainPage", "\303\234r\303\274n Sil", "", 0))
	w.EditProductPushButton.SetText(core.QCoreApplication_Translate("MainPage", "\303\234r\303\274n D\303\274zenle", "", 0))
	w.PushButton_5.SetText(core.QCoreApplication_Translate("MainPage", "PushButton", "", 0))
	w.PushButton_4.SetText(core.QCoreApplication_Translate("MainPage", "PushButton", "", 0))

}
