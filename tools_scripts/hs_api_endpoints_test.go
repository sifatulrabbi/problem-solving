package tools_scripts

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"
)

func TestActivatePlan(t *testing.T) {
	accessToken, err := getAccessToken()
	if err != nil {
		t.Error(err)
		return
	}

	payload := url.Values{
		"action":            {"activate"}, // enhance_tier, reduce_tier, activate, refund, update
		"plan_id":           {"helloscribe_tier1"},
		"uuid":              {"3fafab3e-13b0-4527-80ab-8f402fcasdf34"},
		"activation_email":  {"islammasraful@gmail.com"},
		"invoice_item_uuid": {"invoice-3fafab3e-13b0-4527-80ab-8f402fcasdf34"},
	}
	req, err := http.NewRequest(http.MethodPost, getUrl("api/sumoling-notification"), strings.NewReader(payload.Encode()))
	if err != nil {
		t.Error(err)
		return
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
		return
	}
	result := map[string]any{}
	if err = json.NewDecoder(res.Body).Decode(&result); err != nil {
		t.Error(err)
		return
	}
	if res.StatusCode != http.StatusCreated {
		if msg, ok := result["message"]; !ok {
			fmt.Println(result)
			t.Error(errors.New("unable to full fill the request."))
		} else {
			if m, ok := msg.(string); !ok {
				t.Error(errors.New("unable to full fill the request."))
			} else {
				t.Error(fmt.Errorf(m))
			}
		}
	}

	fmt.Println(result)
}

func TestRefund(t *testing.T) {
	accessToken, err := getAccessToken()
	if err != nil {
		t.Error(err)
		return
	}

	payload := url.Values{
		"action":           {"refund"}, // enhance_tier, reduce_tier, activate, refund, update
		"uuid":             {"3fafab3e-13b0-4527-80ab-8f402fcasdf34"},
		"activation_email": {"islammasraful@gmail.com"},
	}
	req, err := http.NewRequest(http.MethodPost, getUrl("api/sumoling-notification"), strings.NewReader(payload.Encode()))
	if err != nil {
		t.Error(err)
		return
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
		return
	}
	result := map[string]any{}
	if err = json.NewDecoder(res.Body).Decode(&result); err != nil {
		t.Error(err)
		return
	}
	if res.StatusCode != http.StatusOK {
		if msg, ok := result["message"]; !ok {
			fmt.Println(result)
			t.Error(errors.New("unable to full fill the request."))
		} else {
			if m, ok := msg.(string); !ok {
				t.Error(errors.New("unable to full fill the request."))
			} else {
				t.Error(fmt.Errorf(m))
			}
		}
	}
	fmt.Println(result)
}

func TestDowngradeTier(t *testing.T) {
	accessToken, err := getAccessToken()
	if err != nil {
		t.Error(err)
		return
	}

	payload := url.Values{
		"action":           {"reduce_tier"}, // enhance_tier, reduce_tier, activate, refund, update
		"uuid":             {"3fafab3e-13b0-4527-80ab-8f402fcasdf34"},
		"activation_email": {"islammasraful@gmail.com"},
	}
	req, err := http.NewRequest(http.MethodPost, getUrl("api/sumoling-notification"), strings.NewReader(payload.Encode()))
	if err != nil {
		t.Error(err)
		return
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
		return
	}
	result := map[string]any{}
	if err = json.NewDecoder(res.Body).Decode(&result); err != nil {
		t.Error(err)
		return
	}
	if res.StatusCode != http.StatusOK {
		if msg, ok := result["message"]; !ok {
			fmt.Println(result)
			t.Error(errors.New("unable to full fill the request."))
		} else {
			if m, ok := msg.(string); !ok {
				t.Error(errors.New("unable to full fill the request."))
			} else {
				t.Error(fmt.Errorf(m))
			}
		}
	}
	fmt.Println(result)
}

func TestUpgradeTier(t *testing.T) {
	accessToken, err := getAccessToken()
	if err != nil {
		t.Error(err)
		return
	}

	payload := url.Values{
		"action":           {"enhance_tier"}, // enhance_tier, reduce_tier, activate, refund, update
		"uuid":             {"3fafab3e-13b0-4527-80ab-8f402fcasdf34"},
		"activation_email": {"islammasraful@gmail.com"},
	}
	req, err := http.NewRequest(http.MethodPost, getUrl("api/sumoling-notification"), strings.NewReader(payload.Encode()))
	if err != nil {
		t.Error(err)
		return
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
		return
	}
	result := map[string]any{}
	if err = json.NewDecoder(res.Body).Decode(&result); err != nil {
		t.Error(err)
		return
	}
	if res.StatusCode != http.StatusOK {
		if msg, ok := result["message"]; !ok {
			fmt.Println(result)
			t.Error(errors.New("unable to full fill the request."))
		} else {
			if m, ok := msg.(string); !ok {
				t.Error(errors.New("unable to full fill the request."))
			} else {
				t.Error(fmt.Errorf(m))
			}
		}
	}
	fmt.Println(result)
}

func getAccessToken() (string, error) {
	accessToken := ""

	if err := godotenv.Load(); err != nil {
		return accessToken, err
	}

	USERNAME := os.Getenv("APPSUMO_USERNAME")
	PASSWORD := os.Getenv("APPSUMO_PASSWORD")
	if USERNAME == "" || PASSWORD == "" {
		return accessToken, errors.New("not enough env vars.")
	}

	payload, err := json.Marshal(map[string]string{
		"username": USERNAME,
		"password": PASSWORD,
	})
	if err != nil {
		return accessToken, err
	}

	req, err := http.NewRequest(http.MethodPost, getUrl("api/sumoling-token"), bytes.NewBuffer(payload))
	if err != nil {
		return accessToken, err
	}

	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return accessToken, err
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	result := map[string]any{}
	if err = json.NewDecoder(res.Body).Decode(&result); err != nil {
		return accessToken, err
	}

	if token, ok := result["access"]; !ok {
		return accessToken, errors.New("not access token found in the response")
	} else {
		accessToken, ok = token.(string)
		if !ok {
			return accessToken, errors.New("the access token is not a valid string.")
		}
	}

	return accessToken, err
}

func getUrl(path string) string {
	return fmt.Sprintf("https://app.helloscribe.ai/%s", path)
}
