package randomName

import (
	"encoding/json"
	"fmt"
	"io"
	"wanna-be-chuck-norris-app/internal/cache"
	"wanna-be-chuck-norris-app/pkg/httpclient"
)

type NameResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
type Service interface {
	GetName() (*NameResponse, error)
}

type service struct {
	baseURL    string
	cache      cache.SimpleCache[NameResponse]
	httpClient httpclient.Client
}

func (s *service) GetName() (*NameResponse, error) {
	if cached, ok := s.cache.Pop(); ok {
		go s.addToCache()
		return &cached, nil
	}
	return s.fetchName()
}
func (s *service) fetchName() (*NameResponse, error) {
	resp, err := s.httpClient.Get(s.baseURL)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var nameResponse NameResponse
	err = json.Unmarshal(body, &nameResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	return &nameResponse, nil
}

// TODO: More elegant populate cache solution
func (s *service) populateCache(numEntries int) {
	for i := 0; i < numEntries; i++ {
		go s.addToCache()
	}
}

// TODO: Handle what happens on failure to add to cache properly
func (s *service) addToCache() {
	if name, err := s.fetchName(); err == nil {
		s.cache.Add(*name)
	}
}

func NewService(baseURL string, httpClient httpclient.Client, initialCacheSize int) *service {
	newService := &service{
		baseURL:    baseURL,
		cache:      cache.NewSimpleCache[NameResponse](),
		httpClient: httpClient,
	}
	newService.populateCache(initialCacheSize)
	return newService
}
