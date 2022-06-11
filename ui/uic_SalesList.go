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
	HorizontalLayout     *widgets.QHBoxLayout
	SalesListTableWidget *widgets.QTableWidget
	VerticalLayout       *widgets.QVBoxLayout
	UndoSalePushButton   *widgets.QPushButton
	VerticalSpacer       *widgets.QSpacerItem
	PushButton_2         *widgets.QPushButton
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
	w.Resize2(567, 291)
	w.HorizontalLayout = widgets.NewQHBoxLayout2(w)
	w.HorizontalLayout.SetObjectName("horizontalLayout")
	w.SalesListTableWidget = widgets.NewQTableWidget(w)
	w.SalesListTableWidget.SetObjectName("salesListTableWidget")
	w.SalesListTableWidget.SetSelectionBehavior(widgets.QAbstractItemView__SelectRows)
	w.HorizontalLayout.QLayout.AddWidget(w.SalesListTableWidget)
	w.VerticalLayout = widgets.NewQVBoxLayout()
	w.VerticalLayout.SetObjectName("verticalLayout")
	w.UndoSalePushButton = widgets.NewQPushButton(w)
	w.UndoSalePushButton.SetObjectName("undoSalePushButton")
	w.VerticalLayout.QLayout.AddWidget(w.UndoSalePushButton)
	w.VerticalSpacer = widgets.NewQSpacerItem(20, 40, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	w.VerticalLayout.AddItem(w.VerticalSpacer)
	w.PushButton_2 = widgets.NewQPushButton(w)
	w.PushButton_2.SetObjectName("pushButton_2")
	w.VerticalLayout.QLayout.AddWidget(w.PushButton_2)
	w.HorizontalLayout.AddLayout(w.VerticalLayout, 0)
	w.retranslateUi()
	core.QMetaObject_ConnectSlotsByName(w)

}
func (w *Form) retranslateUi() {
	w.SetWindowTitle(core.QCoreApplication_Translate("Form", "Form", "", 0))
	w.UndoSalePushButton.SetText(core.QCoreApplication_Translate("Form", "Sat\304\261\305\237\304\261 Ger\304\261 Al", "", 0))
	w.PushButton_2.SetText(core.QCoreApplication_Translate("Form", "PushButton", "", 0))

}
