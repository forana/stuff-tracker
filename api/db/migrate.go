package db

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // go away golint
	"github.com/sirupsen/logrus"
)

func Migrate() error {
	conn, err := Open()
	if err != nil {
		return err
	}
	defer conn.Close()
	migrateDriver, err := postgres.WithInstance(conn, &postgres.Config{})
	if err != nil {
		return err
	}
	migrator, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		migrateDriver)
	if err != nil {
		return err
	}
	migrator.Log = &migrateLogger{}
	err = migrator.Up()
	if err == migrate.ErrNoChange {
		logrus.Infof("Migrations already up to date")
		return nil
	}
	return err
}

type migrateLogger struct{}

func (m *migrateLogger) Printf(format string, v ...interface{}) {
	logrus.Infof(format, v...)
}

func (m *migrateLogger) Verbose() bool {
	return false
}
