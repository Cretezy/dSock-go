package dsock

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type DSockClient struct {
	url   string
	token string
}

func NewClient(url string, token string) *DSockClient {
	return &DSockClient{
		url:   url,
		token: token,
	}
}

func NewCloudClient(projectId string, key string) *DSockClient {
	return &DSockClient{
		url:   "https://api.dsock.cloud/" + projectId,
		token: key,
	}
}

type DSockError struct {
	Message  string
	Code     string
	Response []byte
}

func (err *DSockError) Error() string {
	return fmt.Sprintf("dSock error (%s): %s", err.Code, err.Message)
}

type Target struct {
	Id      string
	Channel string
	User    string
	/// Depends on [user]
	Session string
}

func addTargetToParams(target Target, params url.Values) {
	if target.User != "" {
		params.Add("user", target.User)
	}
	if target.Session != "" {
		params.Add("session", target.Session)
	}
	if target.Id != "" {
		params.Add("id", target.Id)
	}
	if target.Channel != "" {
		params.Add("channel", target.Channel)
	}
}

func (client *DSockClient) doApiRequest(method, path string, params url.Values, body io.Reader) ([]byte, error) {
	request, err := http.NewRequest(method, client.url+path+"?"+params.Encode(), body)
	if err != nil {
		return []byte{}, err
	}

	request.Header.Set("Authorization", "Bearer "+client.token)

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return responseBody, nil
}
