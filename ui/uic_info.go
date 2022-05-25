package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type __info struct{}

func (*__info) init() {}

type Info struct {
	*__info
	*widgets.QWidget
	HorizontalLayout *widgets.QHBoxLayout
	TextEdit         *widgets.QTextEdit
}

func NewInfo(p widgets.QWidget_ITF) *Info {
	var par *widgets.QWidget
	if p != nil {
		par = p.QWidget_PTR()
	}
	w := &Info{QWidget: widgets.NewQWidget(par, 0)}
	w.setupUI()
	w.init()
	return w
}
func (w *Info) setupUI() {
	if w.ObjectName() == "" {
		w.SetObjectName("Info")
	}
	w.Resize2(492, 209)
	w.HorizontalLayout = widgets.NewQHBoxLayout2(w)
	w.HorizontalLayout.SetObjectName("horizontalLayout")
	w.TextEdit = widgets.NewQTextEdit(w)
	w.TextEdit.SetObjectName("textEdit")
	font := gui.NewQFont()
	font.SetPointSize(20)
	w.TextEdit.SetFont(font)
	w.TextEdit.SetReadOnly(true)
	w.HorizontalLayout.QLayout.AddWidget(w.TextEdit)
	w.retranslateUi()
	core.QMetaObject_ConnectSlotsByName(w)

}
func (w *Info) retranslateUi() {
	w.SetWindowTitle(core.QCoreApplication_Translate("Info", "Form", "", 0))

}
