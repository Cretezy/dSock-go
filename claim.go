package dsock

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"
)

type CreateClaimOptions struct {
	Id         string
	User       string
	Session    string
	Duration   int
	Expiration int
	Channels   []string
}

type Claim struct {
	Id         string   `json:"id"`
	User       string   `json:"user"`
	Session    string   `json:"session"`
	Expiration int      `json:"expiration"`
	Channels   []string `json:"channels"`
}

type createClaimResponse struct {
	Claim     Claim  `json:"claim"`
	Success   bool   `json:"success"`
	Error     string `json:"error"`
	ErrorCode string `json:"errorCode"`
}

func (client *DSockClient) CreateClaim(options CreateClaimOptions) (*Claim, error) {
	params := url.Values{}

	params.Add("user", options.User)

	if options.Session != "" {
		params.Add("session", options.Session)
	}
	if options.Id != "" {
		params.Add("id", options.Id)
	}
	if len(options.Channels) != 0 {
		params.Add("channels", strings.Join(options.Channels, ","))
	}
	if options.Duration != 0 {
		params.Add("duration", strconv.Itoa(options.Duration))
	}
	if options.Expiration != 0 {
		params.Add("expiration", strconv.Itoa(options.Expiration))
	}

	response, err := client.doApiRequest("POST", "/claim", params, nil)

	var result createClaimResponse

	println(string(response))

	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	if !result.Success {
		return nil, &DSockError{
			Message:  result.Error,
			Code:     result.ErrorCode,
			Response: response,
		}
	}

	return &result.Claim, nil
}
