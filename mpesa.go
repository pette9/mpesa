package mpesa

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

//Client ...
type Client struct{}

func authenticateRedirectPolicyFunc(req *http.Request, via []*http.Request) error {
	req.SetBasicAuth(viper.GetString("mpesa_consumer_key"), viper.GetString("mpesa_consumer_secret"))
	req.Header.Set("cache-control", "no-cache")
	return nil
}

//Authenticate ...
func (mpesa *Client) Authenticate() (map[string]string, error) {
	var mpesaToken map[string]string

	client := &http.Client{
		CheckRedirect: authenticateRedirectPolicyFunc,
	}

	req, err := http.NewRequest("GET", viper.GetString("mpesa_authenticate_url"), nil)
	if err != nil {
		return mpesaToken, err
	}

	req.SetBasicAuth(viper.GetString("mpesa_consumer_key"), viper.GetString("mpesa_consumer_secret"))
	req.Header.Set("cache-control", "no-cache")

	resp, err := client.Do(req)
	if err != nil {
		return mpesaToken, err
	}
	if resp.StatusCode > 299 {
		errBody, _ := ioutil.ReadAll(resp.Body)
		logrus.Println(string(errBody))
		return mpesaToken, fmt.Errorf("The mpesa api error code for authenticate: %v, error text: %s", resp.StatusCode, resp.Status)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &mpesaToken)
	return mpesaToken, nil
}

//MakeRequest ...
func (mpesa *Client) MakeRequest(token string, request string, params interface{}) (map[string]interface{}, error) {
	requestURL := ""
	var response map[string]interface{}
	switch requestName := request; requestName {
	default:
		return response, errors.New("invalid request")
	case "b2b":
		requestURL = viper.GetString("mpesa_b2b_url")
	case "b2c":
		requestURL = viper.GetString("mpesa_b2c_url")
	case "c2b-register-url":
		requestURL = viper.GetString("mpesa_c2b_register_url")
	case "c2b-simulate":
		requestURL = viper.GetString("mpesa_c2b_simulate_url")
	case "lipa-na-mpesa-online-payment":
		requestURL = viper.GetString("lipa_na_mpesa_online_payment_url")
	case "lipa-na-mpesa-online-query":
		requestURL = viper.GetString("lipa_na_mpesa_online_query_url")
	case "reversal":
		requestURL = viper.GetString("mpesa_reversal_url")
	case "transaction-status":
		requestURL = viper.GetString("mpesa_transaction_status_url")
	}
	client := &http.Client{}
	body, err := json.Marshal(params)
	if err != nil {
		return response, err
	}
	req, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(body))
	if err != nil {
		return response, err
	}
	req.Header.Set("authorization", "Bearer "+token)
	req.Header.Set("content-type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return response, err
	}
	if resp.StatusCode > 299 {
		var errorResponse map[string]interface{}
		errBody, _ := ioutil.ReadAll(resp.Body)
		logrus.Println(string(errBody))
		json.Unmarshal(errBody, &errorResponse)
		return response, fmt.Errorf("The mpesa api error code: %v, error text: %s, error message: %s", resp.StatusCode, resp.Status, errorResponse["errorMessage"].(string))
	}
	body, _ = ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &response)
	return response, nil
}
