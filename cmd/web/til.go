package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/url"
	"path"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/yuin/goldmark"
	"gopkg.in/yaml.v3"
)

const (
	tilTreeURL = "https://api.github.com/repos/jackmford/til/git/trees/main?recursive=1"
	tilRawURL  = "https://raw.githubusercontent.com/jackmford/til/main/"
	tilRefresh = 5 * time.Minute
)

type tilPost struct {
	Title     string
	Published time.Time
	Content   template.HTML
}

type tilFrontmatter struct {
	Title     string `yaml:"title"`
	Published string `yaml:"published"`
}

type tilTree struct {
	Tree []struct {
		Path string `json:"path"`
		SHA  string `json:"sha"`
		Type string `json:"type"`
	} `json:"tree"`
}

type tilFeed struct {
	client    *http.Client
	checkedAt time.Time
	posts     map[string]tilPost
	blobs     map[string]string
	mu        sync.Mutex
}

func newTILFeed() *tilFeed {
	return &tilFeed{client: http.DefaultClient}
}

func (f *tilFeed) get(ctx context.Context) ([]tilPost, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	if time.Since(f.checkedAt) < tilRefresh {
		return sortedTILPosts(f.posts), nil
	}
	f.checkedAt = time.Now()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, tilTreeURL, nil)
	if err != nil {
		return sortedTILPosts(f.posts), err
	}
	request.Header.Set("Accept", "application/vnd.github+json")
	response, err := f.client.Do(request)
	if err != nil {
		return sortedTILPosts(f.posts), err
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusConflict {
		f.posts = map[string]tilPost{}
		f.blobs = map[string]string{}
		return nil, nil
	}
	if response.StatusCode != http.StatusOK {
		return sortedTILPosts(f.posts), fmt.Errorf("GitHub TIL tree: %s", response.Status)
	}

	var tree tilTree
	if err := json.NewDecoder(response.Body).Decode(&tree); err != nil {
		return sortedTILPosts(f.posts), err
	}

	nextPosts := make(map[string]tilPost)
	nextBlobs := make(map[string]string)
	for _, entry := range tree.Tree {
		if entry.Type != "blob" || path.Ext(entry.Path) != ".md" || strings.EqualFold(path.Base(entry.Path), "README.md") {
			continue
		}

		nextBlobs[entry.Path] = entry.SHA
		if f.blobs[entry.Path] == entry.SHA {
			nextPosts[entry.Path] = f.posts[entry.Path]
			continue
		}

		post, err := f.fetchPost(ctx, entry.Path)
		if err != nil {
			return sortedTILPosts(f.posts), err
		}
		nextPosts[entry.Path] = post
	}

	f.posts = nextPosts
	f.blobs = nextBlobs
	return sortedTILPosts(f.posts), nil
}

func (f *tilFeed) fetchPost(ctx context.Context, filePath string) (tilPost, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, tilRawURL+escapePath(filePath), nil)
	if err != nil {
		return tilPost{}, err
	}
	response, err := f.client.Do(request)
	if err != nil {
		return tilPost{}, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return tilPost{}, fmt.Errorf("GitHub TIL file %s: %s", filePath, response.Status)
	}
	content, err := io.ReadAll(response.Body)
	if err != nil {
		return tilPost{}, err
	}
	return parseTIL(content)
}

func parseTIL(content []byte) (tilPost, error) {
	parts := bytes.SplitN(content, []byte("---\n"), 3)
	if len(parts) != 3 {
		return tilPost{}, fmt.Errorf("invalid TIL front matter")
	}

	var frontmatter tilFrontmatter
	if err := yaml.Unmarshal(parts[1], &frontmatter); err != nil {
		return tilPost{}, err
	}
	if frontmatter.Title == "" || frontmatter.Published == "" {
		return tilPost{}, fmt.Errorf("TIL title and published date are required")
	}
	published, err := time.Parse(time.RFC3339, frontmatter.Published)
	if err != nil {
		return tilPost{}, fmt.Errorf("invalid TIL published date: %w", err)
	}

	var rendered bytes.Buffer
	if err := goldmark.Convert(parts[2], &rendered); err != nil {
		return tilPost{}, err
	}
	return tilPost{Title: frontmatter.Title, Published: published, Content: template.HTML(rendered.String())}, nil
}

func sortedTILPosts(posts map[string]tilPost) []tilPost {
	result := make([]tilPost, 0, len(posts))
	for _, post := range posts {
		result = append(result, post)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Published.After(result[j].Published)
	})
	return result
}

func escapePath(filePath string) string {
	parts := strings.Split(filePath, "/")
	for i, part := range parts {
		parts[i] = url.PathEscape(part)
	}
	return strings.Join(parts, "/")
}
