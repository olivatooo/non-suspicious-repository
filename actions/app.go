package actions

import (
	"account/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo-pop/v3/pop/popmw"
	"github.com/gobuffalo/envy"
	i18n "github.com/gobuffalo/mw-i18n/v2"
	paramlogger "github.com/gobuffalo/mw-paramlogger"
)

var ENV = envy.Get("GO_ENV", "development")

var (
	app *buffalo.App
	T   *i18n.Translator
)

func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:         ENV,
			SessionName: "_account_session",
		})

		app.Use(paramlogger.ParameterLogger)
		app.Use(popmw.Transaction(models.DB))

	}

	return app
}
