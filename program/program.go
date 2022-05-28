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

	var dir, file string
	dir, err = os.UserHomeDir()
	if err != nil {
		file = path.Join(dir, DBNAME)
	} else {
		file = path.Join(os.TempDir(), DBNAME)
	}

	Inventory, err = backend.Initialize(file)

	if err != nil {
		ShowInfo("Hata", err.Error())
	}

	productWindow := CreateMainPage(nil)

	productWindow.Show()

	widgets.QApplication_Exec()
}
