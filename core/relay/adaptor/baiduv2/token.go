package baiduv2

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/bytedance/sonic"
	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
)

var tokenCache = cache.New(time.Hour*23, time.Minute)

func GetBearerToken(ctx context.Context, apiKey string) (string, error) {
	parts := strings.Split(apiKey, "|")
	if len(parts) != 2 {
		return "", errors.New("invalid baidu apikey")
	}

	if val, ok := tokenCache.Get(apiKey); ok {
		token, ok := val.(string)
		if !ok {
			panic(fmt.Sprintf("invalid cache value type: %T", val))
		}

		return token, nil
	}

	tokenResponse, err := getBaiduAccessTokenHelper(ctx, apiKey)
	if err != nil {
		log.Errorf("get baiduv2 access token failed: %v", err)
		return "", errors.New("get baiduv2 access token failed")
	}

	tokenCache.Set(
		apiKey,
		tokenResponse.Token,
		time.Until(tokenResponse.ExpireTime.Add(-time.Minute*10)),
	)

	return tokenResponse.Token, nil
}

type TokenResponse struct {
	ExpireTime time.Time `json:"expireTime"`
	Token      string    `json:"token"`
}

func getBaiduAccessTokenHelper(ctx context.Context, apiKey string) (*TokenResponse, error) {
	ak, sk, err := getAKAndSK(apiKey)
	if err != nil {
		return nil, err
	}

	authorization := generateAuthorizationString(ak, sk)

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		"https://iam.bj.baidubce.com/v1/BCE-BEARER/token",
		nil,
	)
	if err != nil {
		return nil, err
	}

	query := req.URL.Query()
	query.Add("expireInSeconds", "86400")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Authorization", authorization)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("get token failed, status code: %d", res.StatusCode)
	}

	var tokenResponse TokenResponse

	err = sonic.ConfigDefault.NewDecoder(res.Body).Decode(&tokenResponse)
	if err != nil {
		return nil, err
	}

	return &tokenResponse, nil
}

func generateAuthorizationString(ak, sk string) string {
	httpMethod := http.MethodGet
	uri := "/v1/BCE-BEARER/token"
	queryString := "expireInSeconds=86400"
	hostHeader := "iam.bj.baidubce.com"
	canonicalRequest := fmt.Sprintf("%s\n%s\n%s\nhost:%s", httpMethod, uri, queryString, hostHeader)

	timestamp := time.Now().UTC().Format("2006-01-02T15:04:05Z")
	expirationPeriodInSeconds := 1800
	authStringPrefix := fmt.Sprintf(
		"bce-auth-v1/%s/%s/%d",
		ak,
		timestamp,
		expirationPeriodInSeconds,
	)

	signingKey := hmacSHA256(sk, authStringPrefix)

	signature := hmacSHA256(signingKey, canonicalRequest)

	signedHeaders := "host"
	authorization := fmt.Sprintf("%s/%s/%s", authStringPrefix, signedHeaders, signature)

	return authorization
}

func hmacSHA256(key, data string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
