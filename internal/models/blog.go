package models

import "time"

type BlogPost struct {
    Title       string
    Slug        string    // URL-friendly version of the title
    Date        time.Time
    Content     string
    Preview     string    // First paragraph or excerpt for the preview
}

// BlogPosts is a slice of BlogPost that implements sort.Interface
type BlogPosts []BlogPost

func (bp BlogPosts) Len() int           { return len(bp) }
func (bp BlogPosts) Swap(i, j int)      { bp[i], bp[j] = bp[j], bp[i] }
func (bp BlogPosts) Less(i, j int) bool { return bp[i].Date.After(bp[j].Date) } 
