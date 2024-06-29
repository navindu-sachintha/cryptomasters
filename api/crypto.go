package api

import (
	"encoding/json"
	"fem/go/crypto/models"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*models.Rate, error) {
	if len(currency) != 3 {
		return nil, fmt.Errorf("3 Characters required; %d recived", len(currency))
	}
	upperCurrency := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiUrl, upperCurrency))

	if err != nil {
		return nil, err
	}

	var response CEXResponses
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(bodyBytes, &response)
		if err != nil {
			return nil, err
		}

	} else {
		return nil, fmt.Errorf("status code: %v", res.StatusCode)
	}
	rate := models.Rate{Currency: currency, Price: response.Bid}
	return &rate, nil
}
