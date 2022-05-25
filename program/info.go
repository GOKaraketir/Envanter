package program

import (
	"github.com/GOKaraketir/Envanter/ui"
)

func ShowInfo(title, str string) *ui.Info {
	info := ui.NewInfo(nil)
	info.SetWindowTitle(title)
	info.TextEdit.SetText(str)
	info.Show()
	return info
}
