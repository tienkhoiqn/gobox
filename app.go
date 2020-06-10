package gobox

import (
	"os"

	"github.com/tpphu/gobox/logger"
	"github.com/tpphu/gobox/provider"
	"github.com/tpphu/gobox/service"
	"github.com/tpphu/gobox/service/http"
	"github.com/urfave/cli/v2"
)

// App is main application of gobox
type App struct {
	cli.App
	httpService *http.Http
	Services    []service.Runable
	Provider    provider.Provider
	Log         *logger.Logger
}

type Option func(a *App)

func Name(name string) Option {
	return func(a *App) {
		a.Name = name
	}
}

func Description(desc string) Option {
	return func(a *App) {
		a.Description = desc
	}
}

func WithHTTPService(address string) Option {
	return func(a *App) {
		a.httpService = http.Default(http.Address(address))
	}
}

func (a *App) Init() {

}

func (a *App) Run() {
	a.App.Run(os.Args)
	if a.httpService != nil {
		a.httpService.Init()
		a.httpService.Run()
	}
}

func (a *App) AddService(s service.Runable) {
	a.Services = append(a.Services, s)
}

func (a *App) up(ctx *cli.Context) error {
	a.Log.Info("Application is up")
	return nil
}

func (a *App) down(ctx *cli.Context) error {
	a.Log.Info("Application is down")
	return nil
}

func (a *App) seed(ctx *cli.Context) error {
	a.Log.Info("Application is seeding data")
	return nil
}

func (a *App) migrate(ctx *cli.Context) error {
	a.Log.Info("Application is migrating schema")
	return nil
}

func NewApp(opts ...Option) *App {
	log := logger.New()
	log.Out = os.Stdout
	var app *App
	app = &App{
		App: cli.App{
			Name:                 "gobox",
			Usage:                "a simple gobox application",
			EnableBashCompletion: true,
			Commands: []*cli.Command{
				{
					Name:  "up",
					Usage: "Up application",
					Action: func(c *cli.Context) error {
						return app.up(c)
					},
				},
				{
					Name:  "down",
					Usage: "Down application",
					Action: func(c *cli.Context) error {
						return app.down(c)
					},
				},
				{
					Name:  "seed",
					Usage: "Seed data for application",
					Action: func(c *cli.Context) error {
						return app.seed(c)
					},
				},
				{
					Name:  "migrate",
					Usage: "Migrate data for application",
					Action: func(c *cli.Context) error {
						return app.migrate(c)
					},
				},
			},
		},
		Provider: provider.NewProvider(),
		Log:      log,
	}
	for _, opt := range opts {
		opt(app)
	}
	return app
}
