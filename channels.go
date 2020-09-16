package dsock

import (
	"encoding/json"
	"net/url"
)

type ChannelOptions struct {
	Target
	Channel      string
	IgnoreClaims bool
}

type channelResponse struct {
	Success   bool   `json:"success"`
	Error     string `json:"error"`
	ErrorCode string `json:"errorCode"`
}

func (client *DSockClient) SubscribeChannel(options ChannelOptions) error {
	params := url.Values{}

	addTargetToParams(options.Target, params)

	if options.IgnoreClaims {
		params.Set("ignoreClaims", "true")
	}

	response, err := client.doApiRequest("POST", "/channel/subscribe/"+options.Channel, params, nil)

	var result channelResponse

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

func (client *DSockClient) UnsubscribeChannel(options ChannelOptions) error {
	params := url.Values{}

	addTargetToParams(options.Target, params)

	if options.IgnoreClaims {
		params.Set("ignoreClaims", "true")
	}

	response, err := client.doApiRequest("POST", "/channel/unsubscribe/"+options.Channel, params, nil)

	var result channelResponse

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
