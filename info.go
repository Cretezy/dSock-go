package dsock

import (
	"encoding/json"
	"net/url"
)

type GetInfoOptions struct {
	Target
}

type InfoClaim struct {
	Id         string   `json:"id"`
	User       string   `json:"user"`
	Session    string   `json:"session"`
	Expiration int      `json:"expiration"`
	Channels   []string `json:"channels"`
}

type InfoConnection struct {
	Id       string   `json:"id"`
	User     string   `json:"user"`
	Session  string   `json:"session"`
	Channels []string `json:"channels"`
	Worker   string   `json:"string"`
	LastPing int      `json:"lastPing"`
}

type GetInfoResponse struct {
	Claims      []InfoClaim
	Connections []InfoConnection
}

type getInfoResponse struct {
	Claims      []InfoClaim      `json:"claims"`
	Connections []InfoConnection `json:"connections"`
	Success     bool             `json:"success"`
	Error       string           `json:"error"`
	ErrorCode   string           `json:"errorCode"`
}

func (client *DSockClient) GetInfo(options GetInfoOptions) (*GetInfoResponse, error) {
	params := url.Values{}

	addTargetToParams(options.Target, params)

	response, err := client.doApiRequest("GET", "/info", params, nil)

	var result getInfoResponse

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

	return &GetInfoResponse{
		Claims:      result.Claims,
		Connections: result.Connections,
	}, nil
}
