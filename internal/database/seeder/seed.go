package seeder

import (
	"fmt"
	"log"

	articleModel "github.com/codercollo/blog/internal/modules/article/models"
	userModel "github.com/codercollo/blog/internal/modules/user/models"

	"github.com/codercollo/blog/pkg/database"
	"golang.org/x/crypto/bcrypt"
)

func Seed() {
	db := database.Connection()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("secret"), 12)
	if err != nil {
		log.Fatal("hash password error")
		return
	}

	user := userModel.User{
		Name:     "Random name",
		Email:    "random@email.com",
		Password: string(hashedPassword),
	}
	db.Create(&user)

	log.Printf("User created successfully with email address %s \n", &user.Email)

	for i := 1; i <= 10; i++ {
		article := articleModel.Article{
			Title:   fmt.Sprintf("Random title %d", i),
			Content: fmt.Sprintf("Random content %d", i),
			UserID:  1,
		}

		db.Create(&article)
		log.Printf("Article create successfully with title %s \n", article.Title)

	}

	log.Println("Seeder done ..")
}
