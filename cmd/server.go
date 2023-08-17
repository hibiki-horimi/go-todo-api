package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	echo "github.com/labstack/echo/v4"
	"github.com/spf13/cobra"

	"github.com/hibiki-horimi/go-todo-api/internal/config"
	"github.com/hibiki-horimi/go-todo-api/internal/database/postgres"
	"github.com/hibiki-horimi/go-todo-api/internal/middleware"
	"github.com/hibiki-horimi/go-todo-api/internal/route"
	"github.com/hibiki-horimi/go-todo-api/internal/server"
)

var serverCmd = &cobra.Command{
	Use:  "server",
	RunE: serverRun,
}

func serverRun(_ *cobra.Command, _ []string) error {
	e := echo.New()

	conf, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	loc, err := time.LoadLocation(conf.OS.TZ)
	if err != nil {
		log.Fatal(err)
	}
	time.Local = loc

	gdb, err := postgres.Connect(conf)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		db, errInDefer := gdb.DB()
		if errInDefer != nil {
			log.Fatal(errInDefer)
		}
		_ = db.Close()
	}()

	rdb := postgres.New()

	middleware.Setup(e, gdb, rdb)

	srv := server.New(rdb, conf)
	route.New(e, srv)

	go func() {
		if err = e.Start(fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port)); err != nil {
			log.Fatal(fmt.Errorf("shutting down the server"))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err = e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	return nil
}
