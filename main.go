//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
//go:generate goversioninfo -icon=res/papp.ico
package main

import (
	"os"

	. "github.com/portapps/portapps"
)

func init() {
	Papp.ID = "insomnia-portable"
	Papp.Name = "Insomnia"
	Init()
}

func main() {
	Papp.AppPath = AppPathJoin("app")
	Papp.DataPath = CreateFolder(AppPathJoin("data"))

	electronBinPath := PathJoin(Papp.AppPath, FindElectronAppFolder("app-", Papp.AppPath))

	Papp.Process = PathJoin(electronBinPath, "Insomnia.exe")
	Papp.Args = nil
	Papp.WorkingDir = electronBinPath

	Launch(os.Args[1:])
}
