package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Repo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	HTMLURL     string `json:"html_url"`
	Language    string `json:"language"`
	Fork        bool   `json:"fork"`
	Stars       int    `json:"stargazers_count"`
}

func GetRepos(c *gin.Context) {
	username := "reldyal"

	url := fmt.Sprintf("https://api.github.com/users/%s/repos?per_page=100", username)

	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal fetch GitHub"})
		return
	}
	defer resp.Body.Close()

	var repos []Repo
	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal parse response"})
		return
	}

	// Filter fork
	var filtered []Repo
	for _, r := range repos {
		if !r.Fork {
			filtered = append(filtered, r)
		}
	}

	c.JSON(http.StatusOK, filtered)
}