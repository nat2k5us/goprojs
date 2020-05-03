package orders

import (
	"crypto/tls"
	"encoding/json"

	//"strconv"
	"strings"

	//"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	
	"github.com/chrisftw/ezconf"
	log "github.com/sirupsen/logrus"
)

// oAuthResponse holds decoded JSON response from Box
type oAuthResponse struct {
	AccessToken  string   `json:"access_token"`
	Expires      int      `json:"expires_in"`
	RestrictedTo []string `json:"restricted_to"`
	TokenType    string   `json:"token_type"`
}

func GetToken() string {
	var decodedResponse *oAuthResponse
	var msg string
	// TODO: This is insecure; use only in dev environments.
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	body := strings.NewReader(`grant_type=http://auth0.com/oauth/grant-type/password-realm&username=omgmkob@tradestation.com&password=Trade123&audience=https://api.tradestation.io&scope=openid+profile+MarketData+ReadAccount+Trade+HotLists+News+Matrix+Crypto&client_id=GbHvpF6vhzMA4xoLrv9t9UBDdl2fTUrI&client_secret=kJclVpoM76vmXiabXQsp-aakqwPL1hii5B9hC8kGrEjw1pwPPcejPi15a4cP9NFj&realm=CryptoDEVAuth0DBTemp`)
	req, err := http.NewRequest("POST", "https://tradestation-qa.auth0.com/oauth/token", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	fmt.Println(resp.Body)
	err = json.NewDecoder(resp.Body).Decode(&decodedResponse)

	if err != nil {
		msg = "Error decoding OAuthResponse"
	} else {
		// We only need the Access Token
		msg = decodedResponse.AccessToken
	}

	return msg
}

func GetOrders(account string, token string) string {

	req, err := http.NewRequest("GET", "http://127.0.0.1:8001/brokerage/accounts/"+account+"/orders", nil)
	if err != nil {
		log.Error("error occured making the request", err)
	}
	tokenBearer := "Bearer " + token
	userID := ezconf.Get("settings", "UserId")
	req.Header.Set("Authorization", tokenBearer)
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Postman-Token", "7b0e630f-00f3-ca93-df46-5b3ccfb77c4a")
	req.Header.Set("X-Ts-Auth-User-Id", userID)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error("error occured making the request", err)
	}
	defer resp.Body.Close()
	var bodyString string
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString = string(bodyBytes)
		//fmt.Println(bodyString)
	} else {
		fmt.Printf("Error Status: %s Code: %d", resp.Status, req.Response.StatusCode)
	}
	return bodyString
}

func GetOrdersParallel(token string) {
	accts := ezconf.Get("settings", "accounts")
	accts = strings.Trim(accts, "\"")
	accounts := strings.Split(accts, ",")
	startAcc := time.Now()
	accountCount := 0
	for _, acct := range accounts {
		go func(acct string) {
			GetOrders(acct, token)
			accountCount++
		}(acct)
	}
	for accountCount < len(accounts) {
		time.Sleep(1 * time.Millisecond)
	}
	elaspedAcc := time.Since(startAcc)
	fmt.Println("Accounts time: ", elaspedAcc)
}
