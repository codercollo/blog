package bootstrap

import (
	"github.com/codercollo/blog/internal/database/migration"
	"github.com/codercollo/blog/pkg/config"
	"github.com/codercollo/blog/pkg/database"
)

func Migrate() {
	config.Set()
	database.Connect()
	migration.Migrate()
}
