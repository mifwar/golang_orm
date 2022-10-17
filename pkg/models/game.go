package models

import (
	"gamestore/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Game struct {
	gorm.Model
	Name      string `json:"name"`
	Genre     string `json:"genre"`
	Publisher string `json:"publisher"`
}

func init() {
	// fmt.Println("init")
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Game{})
}

func (g *Game) CreateGame() *Game {
	db.NewRecord(g)
	db.Create(&g)
	return g
}

func GetAllGames() []Game {
	var Games []Game
	db.Find(&Games)
	return Games
}

func GetGameByID(id int) *Game {
	game := &Game{}
	_ = db.Where("ID = ?", id).Find(game)
	return game
}

func DeleteGame(id int) []Game {
	game := &Game{}
	remainingGames := []Game{}
	_ = db.Where("ID = ?", id).Delete(game)
	db.Find(&remainingGames)

	return remainingGames
}

func UpdateGame(id int, updatedGame *Game) bool {
	game := &Game{}
	games := []Game{}

	_ = db.Where("ID = ?", id).Find(game)

	if game.ID == 0 {
		return false
	}

	game.Name = updatedGame.Name
	game.Genre = updatedGame.Genre
	game.Publisher = updatedGame.Publisher

	db.Save(game)
	db.Find(&games)
	return true
}
