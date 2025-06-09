package telegram

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

func New(host string, token string) Client {
	return Client{
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{},
	}
}

func newBasePath(token string) string {
	return "bot" + token

}

func (c *Client) SendMessage(chat_id int, text string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chat_id))
	q.Add("text", text)
	_, err := c.doRequest("sendMessage", q)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Updates(offset int, limit int) ([]Updates, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(offset))

	data, err := c.doRequest("getUpdates", q)
	if err != nil {
		return nil, err
	}
	var res UpdatesResponce
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res.Result, nil
}

func (c *Client) doRequest(method string, query url.Values) ([]byte, error) {
	const errorStr = "cant do request"

	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method), //c.basePath + method,
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("cant do request: %w", err)
	}
	req.URL.RawQuery = query.Encode()
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("cant do request: %w", err)
	}

	body, err := io.ReadAll(resp.Body)

	return body, nil
}
