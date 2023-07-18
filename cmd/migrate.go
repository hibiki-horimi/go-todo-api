package cmd

import (
	"log"
	"time"

	"github.com/spf13/cobra"

	"github.com/hibiki-horimi/go-todo-api/internal/config"
	"github.com/hibiki-horimi/go-todo-api/internal/database/postgres"
)

var migrateCmd = &cobra.Command{
	Use:  "migrate",
	RunE: migrateRun,
}

func migrateRun(_ *cobra.Command, _ []string) error {
	conf, err := config.New()
	if err != nil {
		log.Fatal(err)
		return err
	}
	loc, err := time.LoadLocation(conf.OS.TZ)
	if err != nil {
		log.Fatal(err)
		return err
	}
	time.Local = loc

	con, err := postgres.Connect(conf)
	if err != nil {
		log.Fatal(err)
		return err
	}

	if err := postgres.Migrate(con); err != nil {
		log.Fatal(err)
		return err
	}

	log.Println("migrate done")
	return nil
}
