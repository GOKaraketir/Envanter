package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type __form struct{}

func (*__form) init() {}

type Form struct {
	*__form
	*widgets.QWidget
	HorizontalLayout *widgets.QHBoxLayout
	TableWidget      *widgets.QTableWidget
	VerticalLayout   *widgets.QVBoxLayout
	PushButton       *widgets.QPushButton
	PushButton_2     *widgets.QPushButton
	PushButton_3     *widgets.QPushButton
}

func NewForm(p widgets.QWidget_ITF) *Form {
	var par *widgets.QWidget
	if p != nil {
		par = p.QWidget_PTR()
	}
	w := &Form{QWidget: widgets.NewQWidget(par, 0)}
	w.setupUI()
	w.init()
	return w
}
func (w *Form) setupUI() {
	if w.ObjectName() == "" {
		w.SetObjectName("Form")
	}
	w.Resize2(973, 497)
	w.HorizontalLayout = widgets.NewQHBoxLayout2(w)
	w.HorizontalLayout.SetObjectName("horizontalLayout")
	w.TableWidget = widgets.NewQTableWidget(w)
	w.TableWidget.SetObjectName("tableWidget")
	w.TableWidget.SetEditTriggers(widgets.QAbstractItemView__NoEditTriggers)
	w.TableWidget.SetSelectionMode(widgets.QAbstractItemView__SingleSelection)
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
	w.PushButton_3 = widgets.NewQPushButton(w)
	w.PushButton_3.SetObjectName("pushButton_3")
	w.VerticalLayout.QLayout.AddWidget(w.PushButton_3)
	w.HorizontalLayout.AddLayout(w.VerticalLayout, 0)
	w.retranslateUi()
	core.QMetaObject_ConnectSlotsByName(w)

}
func (w *Form) retranslateUi() {
	w.SetWindowTitle(core.QCoreApplication_Translate("Form", "Form", "", 0))
	w.PushButton.SetText(core.QCoreApplication_Translate("Form", "PushButton", "", 0))
	w.PushButton_2.SetText(core.QCoreApplication_Translate("Form", "PushButton", "", 0))
	w.PushButton_3.SetText(core.QCoreApplication_Translate("Form", "PushButton", "", 0))

}
