package whois

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goosint/models"
	"io"
	"log"
	"net/http"
	"os"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

func ReverseWhois(input string) models.DomainReverseList {

	url := "https://reverse-whois.whoisxmlapi.com/api/v2"
	method := "POST"

	WHOISXML_APIKEY, exists := os.LookupEnv("WHOISXML_APIKEY")
	if !exists {
		log.Print("WHOISXML_APIKEY error")
	}

	payload := map[string]interface{}{
		"apiKey":     WHOISXML_APIKEY,
		"searchType": "historic",
		"mode":       "purchase",
		"punycode":   true,
		"basicSearchTerms": map[string]interface{}{
			"include": []interface{}{
				input,
			},
		},
	}

	jsonPayload, err := json.Marshal(payload)
	check(err)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonPayload))
	check(err)

	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	check(err)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	check(err)

	resDomain := &models.DomainReverseList{}
	err = json.Unmarshal(body, resDomain)
	check(err)

	if resDomain.DomainsCount > 50 {
		resDomain = &models.DomainReverseList{
			DomainsCount: 50,
			DomainsList:  resDomain.DomainsList[:50],
		}
	}
	return *resDomain
}
