package sms

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/Ahmad940/assetly_server/pkg/config"
)

func SendSms(to string, message string) error {
	africaTalkSmsUrl := "https://api.africastalking.com/version1/messaging"

	// Encode the map to a form-urlencoded string.
	form := url.Values{}
	form.Add("username", config.GetEnv().AFRICA_TALK_USERNAME)
	form.Add("to", to)
	form.Add("message", message)

	// create http request
	req, err := http.NewRequest(http.MethodPost, africaTalkSmsUrl, strings.NewReader(form.Encode()))
	if err != nil {
		log.Println("Error while creating request, reason:", err)
		return err
	}

	// req headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("apiKey", config.GetEnv().AFRICA_TALK_API_KEY)

	// Perform http request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Errored when sending request to the server:", err)
		return err
	}

	// close the response body.
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing body, reason:", err)
			log.Fatal(err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error reading body %s", err)
		return err
	}

	// if the status is not okay, then it not correct
	if resp.StatusCode != http.StatusCreated {
		return errors.New(string(body))
	}

	fmt.Println("Response", string(body))

	return nil
}
