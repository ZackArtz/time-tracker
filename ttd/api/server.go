package api

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/zackartz/ttd/models"
)

type Server struct {
	DB     *gorm.DB
	Router *fiber.App
}

func (s *Server) Initialize() {
	var err error

	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(fmt.Sprintf("%s/.local/share/tt", user.HomeDir)); os.IsNotExist(err) {
		err := os.Mkdir(fmt.Sprintf("%s/.local/share/tt", user.HomeDir), 0700)
		if err != nil {
			log.Printf("Error creating directory %v", err)
		} else {
			log.Printf("Creating directory ~/.local/share/tt")
		}
	}

	s.DB, err = gorm.Open("sqlite3", fmt.Sprintf("%s/.local/share/tt/data.db", user.HomeDir))
	if err != nil {
		log.Panicf("could not open database %v", err)
	}

	s.DB.Debug().AutoMigrate(&models.Timestamp{})

	s.initializeRoutes()

	log.Fatal(s.Router.Listen(6969))
}
