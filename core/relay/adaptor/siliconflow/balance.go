package siliconflow

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bytedance/sonic"
	"github.com/labring/aiproxy/core/model"
	"github.com/labring/aiproxy/core/relay/adaptor"
)

var _ adaptor.Balancer = (*Adaptor)(nil)

func (a *Adaptor) GetBalance(channel *model.Channel) (float64, error) {
	u := channel.BaseURL
	if u == "" {
		u = baseURL
	}

	url := u + "/user/info"

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		return 0, err
	}

	req.Header.Set("Authorization", "Bearer "+channel.Key)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("status code: %d", res.StatusCode)
	}

	response := UsageResponse{}

	err = sonic.ConfigDefault.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return 0, err
	}

	balance, err := strconv.ParseFloat(response.Data.Balance, 64)
	if err != nil {
		return 0, err
	}

	return balance, nil
}

type UsageResponse struct {
	Message string `json:"message"`
	Data    struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		Image         string `json:"image"`
		Email         string `json:"email"`
		Balance       string `json:"balance"`
		Status        string `json:"status"`
		Introduction  string `json:"introduction"`
		Role          string `json:"role"`
		ChargeBalance string `json:"chargeBalance"`
		TotalBalance  string `json:"totalBalance"`
		Category      string `json:"category"`
		IsAdmin       bool   `json:"isAdmin"`
	} `json:"data"`
	Code   int  `json:"code"`
	Status bool `json:"status"`
}
