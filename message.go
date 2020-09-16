package dsock

import (
	"bytes"
	"encoding/json"
	"net/url"
)

type SendMessageOptions struct {
	Target
	/// `text` or `binary`
	Type    string
	Message []byte
}

type sendMessageResponse struct {
	Success   bool   `json:"success"`
	Error     string `json:"error"`
	ErrorCode string `json:"errorCode"`
}

func (client *DSockClient) SendMessage(options SendMessageOptions) error {
	params := url.Values{}

	addTargetToParams(options.Target, params)

	params.Add("type", options.Type)

	response, err := client.doApiRequest("POST", "/send", params, bytes.NewReader(options.Message))

	var result sendMessageResponse

	err = json.Unmarshal(response, &result)
	if err != nil {
		return err
	}

	if result.Error != "" {
		return &DSockError{
			Message: result.Error,
			Code:    result.ErrorCode,
		}
	}

	return nil
}
