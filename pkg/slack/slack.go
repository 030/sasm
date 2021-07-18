package slack

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type Data struct {
	Blocks   []Blocks `json:"blocks"`
	Channel  string   `json:"channel"`
	Icon     string   `json:"icon_emoji"`
	Username string   `json:"username"`
}

type Blocks struct {
	Type string `json:"type"`
	// pointer required otherwise an empty text{} json will be included
	// which is not accepted by slack
	Text   *Text    `json:"text,omitempty"`
	Fields []Fields `json:"fields,omitempty"`
}

type Text struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
}

type Fields struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
}

const url = "https://hooks.slack.com/services/"

func (d *Data) PostMessage(slackToken string) error {
	payloadBytes, err := json.Marshal(d)
	if err != nil {
		return err
	}
	log.Info(string(payloadBytes))
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", url+slackToken, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodyString := string(bodyBytes)
	log.Info(bodyString)
	log.Info(resp.StatusCode)

	return nil
}
