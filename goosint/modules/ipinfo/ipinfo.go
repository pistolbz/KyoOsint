package ipinfo

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type IPInformation struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func IPInfo(ip string) IPInformation {
	resIP := &IPInformation{}

	IPINFO_TOKEN, exists := os.LookupEnv("IPINFO_TOKEN")
	if exists {
		log.Print("IPINFO_TOKEN error")
	}

	url := fmt.Sprintf("https://ipinfo.io/%s/json?token=%s", ip, IPINFO_TOKEN)
	method := "GET"
	var err error

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	check(err)
	res, err := client.Do(req)
	check(err)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	check(err)
	err = json.Unmarshal(body, resIP)
	check(err)
	return *resIP
}
