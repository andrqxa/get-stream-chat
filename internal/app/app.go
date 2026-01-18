package app

import (
	"fmt"

	"get-stream-chat/internal/config"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

type App struct {
	cfg *config.Config
	pb  *pocketbase.PocketBase
}

func New(cfg *config.Config) (*App, error) {
	pb := pocketbase.New()

	// Register routes via OnServe hook (PocketBase provides the router there).
	pb.OnServe().BindFunc(func(se *core.ServeEvent) error {
		// Example health endpoint
		se.Router.GET("/health", func(e *core.RequestEvent) error {
			return e.String(200, "ok")
		})
		return se.Next()
	})

	return &App{
		cfg: cfg,
		pb:  pb,
	}, nil
}

func (a *App) Run() error {
	// PocketBase is started via "serve" command (same as: go run . serve).
	// We pass --http and --dir so you control port + data dir from config.
	a.pb.RootCmd.SetArgs([]string{
		"serve",
		"--http", fmt.Sprintf("0.0.0.0:%s", a.cfg.Port),
		"--dir", a.cfg.PBDataDir,
	})

	return a.pb.Start()
}
