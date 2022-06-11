package program

import (
	"github.com/GOKaraketir/Envanter/backend"
	"github.com/therecipe/qt/widgets"
	"os"
	"path"
)

var Inventory backend.Inventory
var err error

const DBNAME = "invetoryDatabase.db"

func Run() {
	widgets.NewQApplication(len(os.Args), os.Args)

	var file string

	getwd, err := os.Getwd()
	if err != nil {
		return
	}
	file = path.Join(getwd, DBNAME)

	Inventory, err = backend.Initialize(file)
	productWindow := CreateMainPage(nil)

	if err != nil {
		ShowInfo("Hata", err.Error()+"File: "+file)
	} else {
		productWindow.Show()
		ShowInfo("Database Connected", "File: "+file)
	}

	widgets.QApplication_Exec()
}
