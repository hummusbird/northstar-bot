package util

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type ServerInfo struct {
	Players     int    `json:"playerCount"`
	MaxPlayers  int    `json:"maxPlayers"`
	HasPassword bool   `json:"hasPassword"`
	Region      string `json:"region"`
	Map         string `json:"map"`
	Name        string `json:"name"`
	Playlist    string `json:"playlist"`
}

func getAPIResp() ([]byte, error) {
	url := "https://northstar.tf/client/servers"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func ProcessAPIResp() []ServerInfo {
	data, err := getAPIResp()
	if err != nil {
		log.Println("[ERROR] Could not get data from API")
		log.Println(err)
		return nil
	}

	var respData []ServerInfo

	err = json.Unmarshal(data, &respData)
	if err != nil {
		log.Println("[ERROR] Could not unmarshal JSON")
		log.Println(err)
		return nil
	}

	return respData
}
