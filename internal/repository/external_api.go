package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	ageUrl    = "https://api.agify.io"
	genderUrl = "https://api.genderize.io"
	nationUrl = "https://api.nationalize.io"
)

type ageResp struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}
type nationResp struct {
	Count   int    `json:"count"`
	Name    string `json:"name"`
	Country []struct {
		CountryID   string  `json:"country_id"`
		Probability float64 `json:"probability"`
	} `json:"country"`
}

type genderResp struct {
	Count       int     `json:"count"`
	Name        string  `json:"name"`
	Gender      string  `json:"gender"`
	Probability float64 `json:"probability"`
}

func GetAge(name string) (int, error) {
	req, err := http.Get(fmt.Sprintf("%s?name=%s", ageUrl, name))
	if err != nil {
		return 0, err
	}

	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return 0, err
	}
	res := ageResp{}
	if err := json.Unmarshal(body, &res); err != nil {
		return -1, err
	}

	if res.Count < 1 {
		return -1, err
	}
	return res.Age, nil
}

func GetNation(name string) (string, error) {
	req, err := http.Get(fmt.Sprintf("%s?name=%s", nationUrl, name))
	if err != nil {
		return "", err
	}

	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return "", err
	}
	res := nationResp{}
	if err := json.Unmarshal(body, &res); err != nil {
		return "", err
	}

	if res.Count < 1 {
		return "", err
	}
	return res.Country[0].CountryID, nil
}

func GetGender(name string) (string, error) {
	req, err := http.Get(fmt.Sprintf("%s?name=%s", genderUrl, name))
	if err != nil {
		return "", err
	}

	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return "", err
	}
	res := genderResp{}
	if err := json.Unmarshal(body, &res); err != nil {
		return "", err
	}

	if res.Count < 1 {
		return "", err
	}
	return res.Gender, nil
}
