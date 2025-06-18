package services

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const (
	AgifyURL       = "https://api.agify.io?name="
	GenderizeURL   = "https://api.genderize.io?name="
	NationalizeURL = "https://api.nationalize.io?name="
)

type AgifyResponse struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

type GenderizeResponse struct {
	Count       int     `json:"count"`
	Name        string  `json:"name"`
	Gender      string  `json:"gender"`
	Probability float64 `json:"probability"`
}

type NationalizeResponse struct {
	Count   int    `json:"count"`
	Name    string `json:"name"`
	Country []struct {
		CountryID   string  `json:"country_id"`
		Probability float64 `json:"probability"`
	} `json:"country"`
}

func GetAge(name string) (int, error) {
	resp, err := http.Get(AgifyURL + name)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var agifyResponse AgifyResponse
	if err := json.Unmarshal(body, &agifyResponse); err != nil {
		return 0, err
	}

	return agifyResponse.Age, nil
}

func GetGender(name string) (string, error) {
	resp, err := http.Get(GenderizeURL + name)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var genderizeResponse GenderizeResponse
	if err := json.Unmarshal(body, &genderizeResponse); err != nil {
		return "", err
	}

	return genderizeResponse.Gender, nil
}

func GetNationality(name string) (string, error) {
	resp, err := http.Get(NationalizeURL + name)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var nationalizeResponse NationalizeResponse
	if err := json.Unmarshal(body, &nationalizeResponse); err != nil {
		return "", err
	}

	if len(nationalizeResponse.Country) == 0 {
		return "", errors.New("no country data available")
	}

	return nationalizeResponse.Country[0].CountryID, nil
}
