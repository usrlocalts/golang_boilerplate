package console

import (
	"fmt"
	"strings"

	"golang_boilerplate/config"

	_ "github.com/mattes/migrate/driver/postgres"
	"github.com/mattes/migrate/migrate"
	pipep "github.com/mattes/migrate/pipe"
)

const dbMigrationsPath = "./migrations"

func RunDatabaseMigrations(config *config.Config) error {
	allErrors, ok := migrate.UpSync(config.DatabaseConfig().Url(), dbMigrationsPath)
	if !ok {
		return joinErrors(allErrors)
	}

	fmt.Println("Migration successful")

	return nil
}

func RollbackLatestMigration(config *config.Config) error {
	pipe := pipep.New()

	go migrate.Migrate(pipe, config.DatabaseConfig().Url(), dbMigrationsPath, -1)
	return joinErrors(pipep.ReadErrors(pipe))
}

func joinErrors(errors []error) error {
	var errorMsgs []string
	for _, err := range errors {
		errorMsgs = append(errorMsgs, err.Error())
	}

	return fmt.Errorf(strings.Join(errorMsgs, ","))
}
