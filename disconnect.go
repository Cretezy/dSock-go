package dsock

import (
	"encoding/json"
	"net/url"
)

type DisconnectOptions struct {
	Target

	KeepClaims bool
}

type disconnectResponse struct {
	Success   bool   `json:"success"`
	Error     string `json:"error"`
	ErrorCode string `json:"errorCode"`
}

func (client *DSockClient) Disconnect(options DisconnectOptions) error {
	params := url.Values{}

	addTargetToParams(options.Target, params)

	if options.KeepClaims {
		params.Add("keepClaims", "true")
	}

	response, err := client.doApiRequest("POST", "/disconnect", params, nil)

	var result disconnectResponse

	err = json.Unmarshal(response, &result)
	if err != nil {
		return err
	}

	if !result.Success {
		return &DSockError{
			Message:  result.Error,
			Code:     result.ErrorCode,
			Response: response,
		}
	}
	return nil
}
