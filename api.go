package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"path"
	"time"
)

const LIVE_CLIENT_DATA_URL string = "https://127.0.0.1:2999/liveclientdata"

type API struct {
	httpClient *http.Client
}

func NewAPI() *API {
	// Create a default HTTP Client
	caCert := []byte(riotGamesPem)
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	httpClient := &http.Client{
		Timeout: time.Millisecond * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
			},
		},
	}
	apiInstance := API{
		httpClient: httpClient,
	}
	return &apiInstance
}

// HTTPClient sets the HTTP Client used by the LiveClientData API
func (a *API) HTTPClient(c *http.Client) *API {
	a.httpClient = c
	return a
}

func (a *API) Get(e string, params url.Values) (*http.Response, error) {
	requestURL, err := url.Parse(LIVE_CLIENT_DATA_URL)
	requestURL.Path = path.Join(requestURL.Path, e)
	requestURL.RawQuery = params.Encode()
	if err != nil {
		return nil, err
	}

	resp, err := a.httpClient.Get(requestURL.String())

	// If the status code is not 200, we decode as error JSON to get error message
	if resp.StatusCode != 200 {
		defer resp.Body.Close()
		apiError := APIError{}
		json.NewDecoder(resp.Body).Decode(&apiError)
		return resp, errors.New(apiError.Message)
	}
	return resp, err
}

func (a *API) GetAllGameData() (AllGameData, error) {
	data := AllGameData{}
	resp, err := a.Get("allgamedata", nil)
	if err != nil {
		return data, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetActivePlayer gets all data about the active player.
func (a *API) GetActivePlayer() (ActivePlayer, error) {
	data := ActivePlayer{}
	resp, err := a.Get("activeplayer", nil)
	if err != nil {
		return data, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetActivePlayerName returns the player name.
func (a *API) GetActivePlayerName() (string, error) {
	var data string
	resp, err := a.Get("activeplayername", nil)
	if err != nil {
		return data, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetActivePlayerAbilities gets Abilities for the active player.
func (a *API) GetActivePlayerAbilities() (Abilities, error) {
	data := Abilities{}
	resp, err := a.Get("activeplayerabilities", nil)
	if err != nil {
		return data, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetActivePlayerRunes retrieves the full list of runes for the active player.
func (a *API) GetActivePlayerRunes() (FullRunes, error) {
	data := FullRunes{}
	resp, err := a.Get("activeplayerrunes", nil)
	if err != nil {
		return data, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetAllPlayers retrieves the list of heroes in the game and their stats.
func (a *API) GetAllPlayers() ([]Player, error) {
	data := []Player{}
	resp, err := a.Get("playerlist", nil)
	if err != nil {
		return data, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetPlayerScores retrieves the list of the current scores for the player.
func (a *API) GetPlayerScores(playerName string) (Scores, error) {
	data := Scores{}
	params := url.Values{}
	params.Add("summonerName", playerName)
	resp, err := a.Get("playerscores", params)
	if err != nil {
		return data, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetPlayerSummonerSpells retrieves the list of the summoner spells for the player.
func (a *API) GetPlayerSummonerSpells(playerName string) (SummonerSpells, error) {
	data := SummonerSpells{}
	params := url.Values{}
	params.Add("summonerName", playerName)
	resp, err := a.Get("playersummonerspells", params)
	if err != nil {
		return data, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetPlayerMainRunes retrieves the basic runes of any player.
func (a *API) GetPlayerMainRunes(playerName string) (Runes, error) {
	data := Runes{}
	params := url.Values{}
	params.Add("summonerName", playerName)
	resp, err := a.Get("playermainrunes", params)
	if err != nil {
		return data, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetPlayerItems retrieves the list of items for the player.
func (a *API) GetPlayerItems(playerName string) ([]Item, error) {
	data := []Item{}
	params := url.Values{}
	params.Add("summonerName", playerName)
	resp, err := a.Get("playeritems", params)
	if err != nil {
		return data, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetEvents gets a list of events that have occurred in the game.
func (a *API) GetEvents() (Events, error) {
	data := Events{}
	resp, err := a.Get("eventdata", nil)
	if err != nil {
		return data, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetGameData gets basic data about the game.
func (a *API) GetGameData() (GameData, error) {
	data := GameData{}
	resp, err := a.Get("gamestats", nil)
	if err != nil {
		return data, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}
