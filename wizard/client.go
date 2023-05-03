package wizard

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Client struct {
	client *http.Client
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout can`t be zero")
	}

	return &Client{
		client: &http.Client{
			Timeout: timeout,
			Transport: &loggingRoundTripper{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
		},
	}, nil
}

func (c Client) GetElixirs() ([]Elixir, error) {
	resp, err := c.client.Get("https://wizard-world-api.herokuapp.com/elixirs")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var r ElixirsResponse
	if err = json.Unmarshal(body, &r.Elixir); err != nil {
		return nil, err
	}
	return r.Elixir, nil
}

func (c Client) GetElixir(id string) (Elixir, error) {
	url := fmt.Sprintf("https://wizard-world-api.herokuapp.com/elixirs/%s", id)
	resp, err := c.client.Get(url)
	if err != nil {
		return Elixir{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Elixir{}, err
	}

	var r ElixirResponse
	if err = json.Unmarshal(body, &r.Elixir); err != nil {
		return Elixir{}, err
	}
	return r.Elixir, nil
}
