package mteams

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/hugo-lorenzo-mato/go-notify/conf"
	"github.com/hugo-lorenzo-mato/go-notify/mteams/cards"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var tr *http.Transport = &http.Transport{
	TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
}

type API interface {
	Send(webhookURL string, webhookMessage []byte) error
}

type teamsClient struct {
	httpClient *http.Client
}

func NewClient() (API, error) {
	client := teamsClient{
		httpClient: &http.Client{
			Transport: tr,
			Timeout:   conf.TIMEOUT_MTEAMS * time.Second,
		},
	}
	return &client, nil
}

func CreateNewMessageCard() (cards.DefaultMessageCard, error) {
	log.Debug("        CreateNewMessageCard - Started")
	// Create an instance of the Box struct.
	var card cards.DefaultMessageCard
	err := json.Unmarshal([]byte(cards.DefaultMessageCardExample), &card)
	if err != nil {
		log.Debug("        CreateNewMessageCard - Finished")
		return cards.DefaultMessageCard{}, errors.Wrapf(err, "Error in json.Marshal - CreateNewMessageCard function")
	}
	fmt.Printf("First %s", card)
	log.Debug("CreateNewMessageCard - Finished")
	return card, nil
}

func MarshalMessageAlert(card cards.DefaultMessageCard) ([]byte, error) {
	jsonCard, err := json.Marshal(card)
	if err != nil {
		return []byte{}, errors.Wrapf(err, "Error in json.Marshal - CreateNewAlert function")
	}
	log.Debug(string(jsonCard))
	return jsonCard, nil
}

func (c teamsClient) Send(webhookURL string, webhookMessage []byte) error {
	webhookMessageBuffer := bytes.NewBuffer(webhookMessage)
	req, _ := http.NewRequest("POST", webhookURL, webhookMessageBuffer)
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	res, err := c.httpClient.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	if res.StatusCode >= 299 {
		err = errors.New("error on notification: " + res.Status)
		log.Println(err)
		return err
	}

	return nil
}

func IsValidWebhookURL(webhookURL string) bool {
	if len(webhookURL) == 0 {
		return false
	}
	_, err := url.Parse(webhookURL)
	if err != nil {
		return false
	}
	hasPrefix := strings.HasPrefix(webhookURL, "https://outlook.office.com/webhook/")
	if hasPrefix != true {
		return false
	}
	return true
}
