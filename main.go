//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
package main

import (
	"os"
	"path"
	"path/filepath"

	"github.com/portapps/portapps/v3"
	"github.com/portapps/portapps/v3/pkg/log"
	"github.com/portapps/portapps/v3/pkg/utl"
)

type config struct {
	Cleanup bool `yaml:"cleanup" mapstructure:"cleanup"`
}

var (
	app *portapps.App
	cfg *config
)

func init() {
	var err error

	// Default config
	cfg = &config{
		Cleanup: false,
	}

	// Init app
	if app, err = portapps.NewWithCfg("insomnia-portable", "Insomnia", cfg); err != nil {
		log.Fatal().Err(err).Msg("Cannot initialize application. See log file for more info.")
	}
}

func main() {
	utl.CreateFolder(app.DataPath)
	electronAppPath := app.ElectronAppPath()

	app.Process = filepath.Join(electronAppPath, "Insomnia.exe")
	app.WorkingDir = electronAppPath
	app.Args = []string{
		"--user-data-dir=" + app.DataPath,
	}

	// Cleanup on exit
	if cfg.Cleanup {
		defer func() {
			utl.Cleanup([]string{
				path.Join(os.Getenv("APPDATA"), "Insomnia"),
			})
		}()
	}

	os.Setenv("INSOMNIA_DATA_PATH", app.DataPath)
	os.Setenv("INSOMNIA_DISABLE_AUTOMATIC_UPDATES", "true")

	defer app.Close()
	app.Launch(os.Args[1:])
}
