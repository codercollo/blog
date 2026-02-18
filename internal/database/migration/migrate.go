package migration

import (
	"fmt"
	"log"

	articleModels "github.com/codercollo/blog/internal/modules/article/models"
	userModels "github.com/codercollo/blog/internal/modules/user/models"
	"github.com/codercollo/blog/pkg/database"
)

func Migrate() {
	db := database.Connection()
	err := db.AutoMigrate(&userModels.User{}, &articleModels.Article{})
	if err != nil {
		log.Fatal("Cannot migrate")
		return
	}
	fmt.Println("Migration done!")
}
