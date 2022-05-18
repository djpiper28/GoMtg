package main

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const JSON_URI = "https://mtgjson.com/api/v5/AllPrintings.json"
const MDFC = "modal_dfc"
const FLIP = "flip"
const TRANSFORM = "transform"
const ACCEPTED_FACE = ""

type Card struct {
	OracleId       string `json:"uuid"`
	CardName       string `json:"name"`
	OracleText     string `json:"text,omitempty"`
	Latouts        string `json:"layout,omitempty"`
	scryfallUri    string
	Colour         []string `json:"colors"`
	ColourIdentity []string `json:"colorIdentity"`
	Type           []string `json:"types"`
	Cmc            float64  `json:"convertedManaCost"`
	ManaCost       string   `json:"manaCost,omitempty"`
	Face           string   `json:"face,omitempty"`
}

type Set struct {
	Id    string `json:"uuid"`
	Name  string `json:"name"`
	Cards []Card `json:"cards"`
}

type AllPrintings struct {
	Sets map[string]Set `json:"data"`
}

func (card *Card) GetScryfallUri() string {
	nameenc := url.QueryEscape(card.CardName)
	return "https://scryfall.com/search?q=name%3D%2F%5E" + nameenc + "%24%2F&unique=cards&as=grid&order=name"
}

func FilterCardName(name string) string {
	ret := ""
	for c := range name {
		if c == 'û' {
			ret += "u"
		} else if c >= 'A' && c <= 'Z' {
			ret += string(c - 'A' + 'a')
		} else if c >= 'a' && c <= 'z' {
			ret += string(c)
		}
	}

	return ret
}

func GetCards() (AllPrintings, error) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{},
		},
	}

	resp, err := client.Get(JSON_URI)

	if err != nil {
		return AllPrintings{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return AllPrintings{}, err
	}
	defer resp.Body.Close()

	var data AllPrintings
	err = json.Unmarshal(body, &data)
	if err != nil {
		return AllPrintings{}, err
	}

	return data, nil
}

func main() {

}
