package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func createRequest(method string, url string) *http.Request {
	 return httptest.NewRequest(method, url, nil)
}

func TestGETPlayers(t *testing.T) {

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := createRequest("GET", "/players/Pepper")
		response := httptest.NewRecorder()
		playerServer(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	// TODO(human): Add test for Floyd's score
	t.Run("returns Floyd's score", func(t *testing.T) {
		request := createRequest("GET", "/players/Floyd")
		response := httptest.NewRecorder()

		playerServer(response, request)

		got := response.Body.String()
		want := "40"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestPostWins(t *testing.T) {
	t.Run("it records wins on POST", func(t *testing.T) {
		player := "NewPlayer"
		
		// 发送POST请求记录一次胜利
		postRequest := createRequest("POST", "/players/"+player)
		postResponse := httptest.NewRecorder()
		playerServer(postResponse, postRequest)
		
		// 发送GET请求验证得分增加了
		getRequest := createRequest("GET", "/players/"+player)
		getResponse := httptest.NewRecorder()
		playerServer(getResponse, getRequest)
		
		got := getResponse.Body.String()
		want := "1"
		
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
