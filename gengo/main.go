package main

import (
	"bytes"
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
		fmt.Fprintf(os.Stderr, "stat %s: %v\n", baseDir, err)
		os.Exit(1)
	}
	if !stat.IsDir() {
		fmt.Fprintf(os.Stderr, "%s is not a directory\n", baseDir)
		os.Exit(1)
	}

	articlesMdDir := joinPath(baseDir, "articles.md")
	if err != nil {
		fmt.Fprintf(os.Stderr, "stat %s: %v\n", articlesMdDir, err)
		os.Exit(1)
	}
	if !stat.IsDir() {
		fmt.Fprintf(os.Stderr, "%s is not a directory\n", baseDir)
		os.Exit(1)
	}

	_, articlesDir, err := scaffold(baseDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "scaffolding: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(articlesMdDir, articlesDir)

	parseArticles(articlesMdDir, articlesDir)
}

func parseArticles(sourceDir, targetDir string) ([]*Article, error) {
	articles := make([]*Article, 0)
	entries, _ := os.ReadDir(sourceDir)
	for _, entry := range entries {
		articleSrcPath := joinPath(sourceDir, entry.Name())
		article, err := parseArticle(articleSrcPath)
		if err != nil {
			return articles, err
		}
		articles = append(articles, article)
	}

	return articles, nil
}

func parseArticle(path string) (*Article, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open article %s: %v", path, err)
	}
	data, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("read from %s: %v", path, err)
	}
	content := string(data)
	yaml, md, err := splitParts(content)
	if err != nil {
		return nil, fmt.Errorf("parse article %s: %v", path, err)
	}
	fmt.Println(yaml, md[:100])

	return nil, nil
}

func splitParts(content string) (string, string, error) {
	const sep = "---"
	var yaml, md bytes.Buffer
	inYaml, inMd := false, false
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		if strings.TrimSpace(line) == sep {
			if inYaml {
				inYaml = false
				inMd = true
			} else if yaml.Len() == 0 {
				inYaml = true
			}
		} else if inYaml {
			yaml.WriteString(line + "\n")
		} else if inMd {
			md.WriteString(line + "\n")
		}
	}
	return yaml.String(), md.String(), nil
}

func scaffold(baseDir string) (string, string, error) {
	publicDir := joinPath(baseDir, "public")
	stat, err := os.Stat(publicDir)
	if err == nil && stat.IsDir() {
		os.RemoveAll(publicDir)
	}
	os.Mkdir(publicDir, 0750)

	articleDir := joinPath(publicDir, "articles")
	stat, err = os.Stat(articleDir)
	if err == nil && stat.IsDir() {
		os.RemoveAll(articleDir)
	}
	os.Mkdir(publicDir, 0750)

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
