package realize

import (
	"github.com/fatih/color"
	"sync"
	"fmt"
)

const(
	app_name = "Realize"
	app_version = "v1.0"
	app_email = "pracchia@hastega.it"
	app_description = "Run and build your applications on file changes."
	app_author = "Alessio Pracchia"
	app_file = "realize.config.yaml"
	watcher_ext = ".go"
	watcher_path = "/"
	watcher_ignore = "vendor"
)

var wg sync.WaitGroup
var green = color.New(color.FgGreen, color.Bold).SprintFunc()
var greenl = color.New(color.FgHiGreen).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()
var blue = color.New(color.FgBlue, color.Bold).SprintFunc()
var bluel = color.New(color.FgBlue).SprintFunc()

type App struct{
	Name,Version,Description,Author,Email string
}

func Init() *App{
	return &App{
		Name: app_name,
		Version: app_version,
		Description: app_description,
		Author: app_author,
		Email: app_email,
	}
}

func (app *App) Information(){
	fmt.Println(blue(app.Name)+" - "+blue(app.Version))
	fmt.Println(bluel(app.Description)+"\n")
}

