package seon

import (
	"encoding/json"
	"fmt"
	"goosint/models"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func SeonEmailSearch(email string) models.SeonEmail {
	url := fmt.Sprintf("https://api.seon.io/SeonRestService/email-api/v2.2/%s", email)
	method := "GET"

	SEmail := &models.SeonEmail{}
	payload := strings.NewReader(``)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	check(err)

	SEON_APIKEY, exists := os.LookupEnv("SEON_APIKEY")
	if !exists {
		log.Print("SEON_APIKEY error")
	}

	req.Header.Add("X-API-KEY", SEON_APIKEY)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	check(err)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	check(err)
	err = json.Unmarshal(body, SEmail)
	check(err)
	return *SEmail
}

func SeonPhoneSearch(phone string) models.SeonPhone {
	SPhone := &models.SeonPhone{}
	url := fmt.Sprintf("https://api.seon.io/SeonRestService/phone-api/v1.4/%s", phone)
	method := "GET"

	payload := strings.NewReader(``)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	check(err)

	SEON_APIKEY, exists := os.LookupEnv("SEON_APIKEY")
	if !exists {
		log.Print("SEON_APIKEY error")
	}

	req.Header.Add("X-API-KEY", SEON_APIKEY)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	check(err)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	check(err)
	err = json.Unmarshal(body, SPhone)
	check(err)

	return *SPhone
}
