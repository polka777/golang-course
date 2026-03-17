package github

import (
	"collector/internal/domain"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	client  http.Client
	baseUrl string
}

func NewClient() (c Client) {
	return Client{
		client:  http.Client{Timeout: 6 * time.Second},
		baseUrl: "https://api.github.com",
	}
}

type RepositoryJson struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Stars       int    `json:"stargazers_count"`
	Forks       int    `json:"forks_count"`
	CreateAt    string `json:"created_at"`
}

func (c Client) GetRepositoryInfo(owner, name string) (domain.Repository, error) {
	url := fmt.Sprintf("%s/repos/%s/%s", c.baseUrl, owner, name)
	resp, err := c.client.Get(url)
	if err != nil {
		return domain.Repository{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return domain.Repository{}, fmt.Errorf("error with stutus code : %d", resp.StatusCode)
	}
	if resp.StatusCode == http.StatusNotFound {
		return domain.Repository{}, fmt.Errorf("repository not found: error %d", resp.StatusCode)
	}
	var repoJson RepositoryJson
	err = json.NewDecoder(resp.Body).Decode(&repoJson)
	if err != nil {
		return domain.Repository{}, fmt.Errorf("failed to decode response: %w", err)
	}
	return domain.Repository{
		Name:        repoJson.Name,
		Description: repoJson.Description,
		Stars:       repoJson.Stars,
		Forks:       repoJson.Forks,
		CreatedAt:   repoJson.CreateAt,
	}, nil

}
