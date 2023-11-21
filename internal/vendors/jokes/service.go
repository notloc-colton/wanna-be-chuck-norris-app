package jokes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"wanna-be-chuck-norris-app/pkg/httpclient"
)

type JokeResponse struct {
	Type  string `json:"type"`
	Value struct {
		ID         int      `json:"id"`
		Joke       string   `json:"joke"`
		Categories []string `json:"categories"`
	} `json:"value"`
}
type Service interface {
	GetJoke(firstName, lastName string) (*JokeResponse, error)
}

type service struct {
	baseURL    string
	httpClient httpclient.Client
}

func (s *service) GetJoke(firstName, lastName string) (*JokeResponse, error) {
	apiUrl, err := url.Parse(s.baseURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing base URL: %w", err)
	}

	query := apiUrl.Query()
	query.Set("limitTo", "nerdy")
	query.Set("firstName", firstName)
	query.Set("lastName", lastName)
	apiUrl.RawQuery = query.Encode()

	resp, err := s.httpClient.Get(apiUrl.String())
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var jokeResponse JokeResponse
	err = json.Unmarshal(body, &jokeResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	return &jokeResponse, nil
}

func NewService(baseURL string, httpClient httpclient.Client) *service {
	return &service{
		baseURL,
		httpClient,
	}
}
