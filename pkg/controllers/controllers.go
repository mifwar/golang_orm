package controllers

import (
	"encoding/json"
	"fmt"
	"gamestore/pkg/models"
	"gamestore/pkg/utils"
	"net/http"
)

func CreateGame(w http.ResponseWriter, r *http.Request) {
	createGame := &models.Game{}
	utils.ParseBody(r, createGame)
	game := createGame.CreateGame()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(game)
}

func GetGame(w http.ResponseWriter, r *http.Request) {
	allGames := models.GetAllGames()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allGames)
}

func GetGameByID(w http.ResponseWriter, r *http.Request) {
	id := utils.GetParamID(r, "id")

	game := models.GetGameByID(id)

	w.Header().Set("Content-Type", "application/json")

	if game.ID == 0 {
		res := map[string]string{"result": fmt.Sprintf("can't find game with ID = %d", id)}
		json.NewEncoder(w).Encode(res)
	} else {
		json.NewEncoder(w).Encode(game)
	}

}

func DeleteGame(w http.ResponseWriter, r *http.Request) {
	id := utils.GetParamID(r, "id")

	games := models.DeleteGame(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(games)
}

func UpdateGame(w http.ResponseWriter, r *http.Request) {
	id := utils.GetParamID(r, "id")

	updatedGame := &models.Game{}
	utils.ParseBody(r, updatedGame)

	status := models.UpdateGame(id, updatedGame)

	w.Header().Set("Content-Type", "application/json")

	if status {
		games := models.GetAllGames()
		json.NewEncoder(w).Encode(games)
	} else {
		res := map[string]string{"result": fmt.Sprintf("can't find game with ID = %d", id)}
		json.NewEncoder(w).Encode(res)
	}
}
