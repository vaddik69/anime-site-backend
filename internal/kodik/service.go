package kodik

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type Video struct {
	Title       string   `json:"title"`
	Quality     string   `json:"quality"`
	VoiceActing []string `json:"voice_acting"`
	URL         string   `json:"url"`
}

type SearchResponse struct {
	Results []struct {
		Title       string `json:"title"`
		Translation struct {
			Title string `json:"title"`
			ID    int    `json:"id"`
		} `json:"translation"`
		Link string `json:"link"`
	} `json:"results"`
	Total int `json:"total"`
}

type Service struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
}

func NewService(apiKey string) *Service {
	return &Service{
		apiKey:     apiKey,
		baseURL:    "https://kodikapi.com",
		httpClient: &http.Client{},
	}
}

func (s *Service) SearchAnime(ctx context.Context, title string) ([]Video, error) {
	query := url.Values{}
	query.Add("token", s.apiKey)
	query.Add("title", title)
	query.Add("types", "anime-serial,anime")
	query.Add("with_episodes", "true")

	resp, err := s.httpClient.Get(fmt.Sprintf("%s/search?%s", s.baseURL, query.Encode()))
	if err != nil {
		return nil, fmt.Errorf("kodik search failed: %w", err)
	}
	defer resp.Body.Close()

	var searchResp SearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&searchResp); err != nil {
		return nil, fmt.Errorf("failed to decode kodik response: %w", err)
	}

	var videos []Video
	for _, result := range searchResp.Results {
		videos = append(videos, Video{
			Title:       fmt.Sprintf("%s (%s)", result.Title, result.Translation.Title),
			VoiceActing: []string{result.Translation.Title},
			URL:         result.Link,
		})
	}

	return videos, nil
}

func (s *Service) GetVideoByShikimoriID(ctx context.Context, shikimoriID string) ([]Video, error) {
	query := url.Values{}
	query.Add("token", s.apiKey)
	query.Add("shikimori_id", shikimoriID)
	query.Add("with_episodes", "true")

	resp, err := s.httpClient.Get(fmt.Sprintf("%s/search?%s", s.baseURL, query.Encode()))
	if err != nil {
		return nil, fmt.Errorf("kodik search failed: %w", err)
	}
	defer resp.Body.Close()
	log.Print(resp)
	var searchResp SearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&searchResp); err != nil {
		return nil, fmt.Errorf("failed to decode kodik response: %w", err)
	}

	// Группируем по озвучкам
	voiceMap := make(map[int]Video)
	for _, result := range searchResp.Results {
		if _, exists := voiceMap[result.Translation.ID]; !exists {
			voiceMap[result.Translation.ID] = Video{
				Title:       result.Translation.Title,
				VoiceActing: []string{result.Translation.Title},
				URL:         result.Link,
			}
		}
	}

	var videos []Video
	for _, video := range voiceMap {
		videos = append(videos, video)
	}

	return videos, nil
}
