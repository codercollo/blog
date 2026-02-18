package bootstrap

import (
	"github.com/codercollo/blog/internal/database/seeder"
	"github.com/codercollo/blog/pkg/config"
	"github.com/codercollo/blog/pkg/database"
)

func Seed() {
	config.Set()

	database.Connect()

	seeder.Seed()
}
