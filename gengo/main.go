package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

var staticAssets = []string{"style.css", "favicon.ico", "robots.txt", "key.gpg"}

type Article struct {
	Title     string
	Subtitle  string
	Author    string
	Published time.Time
	Language  string
	Content   string
	// TODO: markdown, name, href
}

func (a Article) PublishedFmt() string {
	switch a.Language {
	case "de":
		return a.Published.Format("01.02.2006 15:04")
	default:
		return a.Published.Format("2006-02-01 15:04")
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s PAGE_DIR\n", os.Args[0])
		os.Exit(1)
	}

	baseDir := os.Args[1]
	stat, err := os.Stat(baseDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s does not exist\n", baseDir)
		os.Exit(1)
	}
	if !stat.IsDir() {
		fmt.Fprintf(os.Stderr, "%s is not a directory\n", baseDir)
		os.Exit(1)
	}

	_, _, err = scaffold(baseDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "scaffolding: %v\n", err)
		os.Exit(1)
	}
}

func scaffold(baseDir string) (string, string, error) {
	publicDir := joinPath(baseDir, "public")
	stat, err := os.Stat(publicDir)
	if err == nil && stat.IsDir() {
		os.RemoveAll(publicDir)
	}
	os.Mkdir(publicDir, 0750)

	articleDir := joinPath(baseDir, "articles.md")
	stat, err = os.Stat(articleDir)
	if err != nil {
		return "", "", fmt.Errorf("stat %s: %v", articleDir, err)
	}
	if !stat.IsDir() {
		return "", "", fmt.Errorf("%s is not a directory", articleDir)
	}

	staticDir := strings.Join([]string{baseDir, "static"}, string(os.PathSeparator))
	for _, entry := range staticAssets {
		source := joinPath(staticDir, entry)
		target := joinPath(publicDir, entry)
		if err := copyFile(source, target); err != nil {
			return "", "", err
		}
	}
	return publicDir, articleDir, nil
}

func joinPath(entries ...string) string {
	return strings.Join(entries, string(os.PathSeparator))
}

func copyFile(from, to string) error {
	src, err := os.Open(from)
	if err != nil {
		return fmt.Errorf("open %s: %v", from, err)
	}
	defer src.Close()
	dst, err := os.Create(to)
	if err != nil {
		return fmt.Errorf("open %s: %v", to, err)
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	return err
}
