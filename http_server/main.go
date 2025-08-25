package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	// TODO: Setup HTTP server
}

// 简单的内存存储来跟踪玩家得分
var playerScores = map[string]int{
	"Pepper": 20,
	"Floyd":  40,
}

func playerServer(w http.ResponseWriter, r *http.Request) {
	playerName := strings.TrimPrefix(r.URL.Path, "/players/")
	
	switch r.Method {
	case http.MethodGet:
		score, exists := playerScores[playerName]
		if !exists {
			http.NotFound(w, r)
			return
		}
		fmt.Fprint(w, score)
		
	case http.MethodPost:
		playerScores[playerName]++
		
	default:
		http.NotFound(w, r)
	}
}

