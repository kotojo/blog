package main

import "testing"

func mockFileReader(filename string) ([]byte, error) {
	return []byte("# Test Title"), nil
}

func TestLoadBlogPost(t *testing.T) {
	bp, err := loadBlogPost("test", mockFileReader)
	if err != nil {
		t.Error("Failed to load blog post file")
	}
	if bp.Title != "test" {
		t.Errorf("loadBlogPost(%q) == %q, want %q", "test", bp.Title, "test")
	}

	if bp.Body != "<h1>Test Title</h1>\n" {
		t.Errorf("loadBlogPost(%q) == %q, want %q", "test", bp.Body, "<h1>Test Title</h1>")
	}
}
