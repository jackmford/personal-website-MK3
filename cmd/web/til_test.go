package main

import (
	"strings"
	"testing"
	"time"
)

func TestParseTILAndSort(t *testing.T) {
	older, err := parseTIL([]byte("---\ntitle: Older\npublished: 2026-09-24T10:00:00Z\n---\n\nOlder note."))
	if err != nil {
		t.Fatal(err)
	}
	newer, err := parseTIL([]byte("---\ntitle: Newer\npublished: 2026-09-24T11:00:00Z\n---\n\nNewer note."))
	if err != nil {
		t.Fatal(err)
	}

	posts := sortedTILPosts(map[string]tilPost{"older.md": older, "nested/newer.md": newer})
	if len(posts) != 2 || posts[0].Title != "Newer" || !strings.Contains(string(posts[0].Content), "Newer note.") {
		t.Fatalf("unexpected posts: %#v", posts)
	}
	if !posts[0].Published.Equal(time.Date(2026, 9, 24, 11, 0, 0, 0, time.UTC)) {
		t.Fatalf("published = %s", posts[0].Published)
	}
}
