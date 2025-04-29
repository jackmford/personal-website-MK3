package blog

import (
    "bytes"
    "fmt"
    "io/fs"
    "jackmitchellfordyce.com/internal/models"
    "jackmitchellfordyce.com/ui"
    "sort"
    "strings"
    "time"

    "github.com/yuin/goldmark"
    "gopkg.in/yaml.v3"
)

type frontmatter struct {
    Title string    `yaml:"title"`
    Date  string    `yaml:"date"`
    Slug  string    `yaml:"slug"`
    Blurb string    `yaml:"blurb"`
}

var (
    postsCache models.BlogPosts
    markdown   = goldmark.New()
)

func init() {
    // Initialize the posts cache
    var err error
    postsCache, err = loadPosts()
    if err != nil {
        fmt.Printf("Error loading posts: %v\n", err)
        postsCache = models.BlogPosts{}
    }
}

func loadPosts() (models.BlogPosts, error) {
    var posts models.BlogPosts

    // Walk through the embedded content/blog directory
    err := fs.WalkDir(ui.Files, "content/blog", func(path string, d fs.DirEntry, err error) error {
        if err != nil {
            return err
        }

        // Skip directories and non-markdown files
        if d.IsDir() || !strings.HasSuffix(path, ".md") {
            return nil
        }

        // Read the file from embedded FS
        content, err := ui.Files.ReadFile(path)
        if err != nil {
            return fmt.Errorf("error reading %s: %v", path, err)
        }

        // Split frontmatter and content
        parts := bytes.SplitN(content, []byte("---\n"), 3)
        if len(parts) != 3 {
            return fmt.Errorf("invalid frontmatter in %s", path)
        }

        // Parse frontmatter
        var fm frontmatter
        if err := yaml.Unmarshal(parts[1], &fm); err != nil {
            return fmt.Errorf("error parsing frontmatter in %s: %v", path, err)
        }

        // Parse date
        date, err := time.Parse("2006-01-02", fm.Date)
        if err != nil {
            return fmt.Errorf("error parsing date in %s: %v", path, err)
        }

        // Convert markdown to HTML
        var buf bytes.Buffer
        if err := markdown.Convert(parts[2], &buf); err != nil {
            return fmt.Errorf("error converting markdown in %s: %v", path, err)
        }

        // Get preview (first paragraph)
        preview := fm.Blurb
        if len(preview) > 150 {
            preview = preview[:150] + "..."
        }

        // Create blog post
        post := models.BlogPost{
            Title:   fm.Title,
            Slug:    fm.Slug,
            Date:    date,
            Content: buf.String(),
            Preview: preview,
        }

        posts = append(posts, post)
        return nil
    })

    if err != nil {
        return nil, err
    }

    // Sort posts by date
    sort.Sort(posts)
    return posts, nil
}

// GetAllPosts returns all blog posts, sorted by date (newest first)
func GetAllPosts() models.BlogPosts {
    return postsCache
}

// GetPostBySlug returns a single blog post by its slug
func GetPostBySlug(slug string) (*models.BlogPost, bool) {
    for _, post := range postsCache {
        if post.Slug == slug {
            return &post, true
        }
    }
    return nil, false
} 
