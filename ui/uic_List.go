package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type __list struct{}

func (*__list) init() {}

type List struct {
	*__list
	*widgets.QWidget
	HorizontalLayout     *widgets.QHBoxLayout
	TableWidget          *widgets.QTableWidget
	VerticalLayout       *widgets.QVBoxLayout
	PushButton           *widgets.QPushButton
	PushButton_2         *widgets.QPushButton
	DeleteProductButtton *widgets.QPushButton
}

func NewList(p widgets.QWidget_ITF) *List {
	var par *widgets.QWidget
	if p != nil {
		par = p.QWidget_PTR()
	}
	w := &List{QWidget: widgets.NewQWidget(par, 0)}
	w.setupUI()
	w.init()
	return w
}
func (w *List) setupUI() {
	if w.ObjectName() == "" {
		w.SetObjectName("List")
	}
	w.Resize2(973, 497)
	w.HorizontalLayout = widgets.NewQHBoxLayout2(w)
	w.HorizontalLayout.SetObjectName("horizontalLayout")
	w.TableWidget = widgets.NewQTableWidget(w)
	w.TableWidget.SetObjectName("tableWidget")
	w.TableWidget.SetEditTriggers(widgets.QAbstractItemView__NoEditTriggers)
	w.TableWidget.SetSelectionMode(widgets.QAbstractItemView__MultiSelection)
	w.TableWidget.SetSelectionBehavior(widgets.QAbstractItemView__SelectRows)
	w.HorizontalLayout.QLayout.AddWidget(w.TableWidget)
	w.VerticalLayout = widgets.NewQVBoxLayout()
	w.VerticalLayout.SetObjectName("verticalLayout")
	w.PushButton = widgets.NewQPushButton(w)
	w.PushButton.SetObjectName("pushButton")
	w.VerticalLayout.QLayout.AddWidget(w.PushButton)
	w.PushButton_2 = widgets.NewQPushButton(w)
	w.PushButton_2.SetObjectName("pushButton_2")
	w.VerticalLayout.QLayout.AddWidget(w.PushButton_2)
	w.DeleteProductButtton = widgets.NewQPushButton(w)
	w.DeleteProductButtton.SetObjectName("deleteProductButtton")
	w.VerticalLayout.QLayout.AddWidget(w.DeleteProductButtton)
	w.HorizontalLayout.AddLayout(w.VerticalLayout, 0)
	w.retranslateUi()
	core.QMetaObject_ConnectSlotsByName(w)

}
func (w *List) retranslateUi() {
	w.SetWindowTitle(core.QCoreApplication_Translate("List", "Form", "", 0))
	w.PushButton.SetText(core.QCoreApplication_Translate("List", "PushButton", "", 0))
	w.PushButton_2.SetText(core.QCoreApplication_Translate("List", "PushButton", "", 0))
	w.DeleteProductButtton.SetText(core.QCoreApplication_Translate("List", "\303\234r\303\274n Sil", "", 0))

}
