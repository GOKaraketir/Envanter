package program

import (
	"github.com/GOKaraketir/Envanter/backend"
	"github.com/therecipe/qt/widgets"
	"os"
)

var Inventory backend.Inventory
var err error

const DBNAME = "database.db"

func Run() {
	widgets.NewQApplication(len(os.Args), os.Args)

	Inventory, err = backend.Initialize(DBNAME)

	if err != nil {
		panic(err)
	}

	productWindow := CreateAddProduct(nil)

	productWindow.Show()

	widgets.QApplication_Exec()
}
