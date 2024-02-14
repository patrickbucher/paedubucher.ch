package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/patrickbucher/paedubucher.ch/gengo"
	yaml "gopkg.in/yaml.v3"
)

var staticAssets = []string{"style.css", "favicon.ico", "robots.txt", "key.gpg"}

func main() {
	if len(os.Args) < 2 {
		panic(fmt.Errorf("usage: %s BASE_DIR", os.Args[0]))
	}

	baseDir := os.Args[1]
	stat, err := os.Stat(baseDir)
	if err != nil {
		panic(fmt.Errorf("stat %s: %v", baseDir, err))
	}
	if !stat.IsDir() {
		panic(fmt.Errorf("%s is not a directory", baseDir))
	}

	articlesMdDir := joinPath(baseDir, "articles.md")
	if err != nil {
		panic(fmt.Errorf("stat %s: %v", articlesMdDir, err))
	}
	if !stat.IsDir() {
		panic(fmt.Errorf("%s is not a directory", baseDir))
	}

	publicDir, articlesDir, err := scaffold(baseDir)
	if err != nil {
		panic(fmt.Errorf("scaffolding: %v", err))
	}

	articles, err := parseArticles(articlesMdDir, articlesDir)
	if err != nil {
		panic(fmt.Errorf("parsing articles: %v", err))
	}

	indexTemplatePath := joinPath(baseDir, "gengo", "index.html")
	indexTemplate, err := template.New("index").Parse(mustRead(indexTemplatePath))
	if err != nil {
		panic(fmt.Errorf("parse index template: %v", err))
	}
	indexFilePath := joinPath(publicDir, "index.html")
	indexFile, err := os.Create(indexFilePath)
	if err != nil {
		panic(fmt.Errorf("open %s: %v", indexFilePath, err))
	}
	defer indexFile.Close()
	err = indexTemplate.Execute(indexFile, articles)
	if err != nil {
		panic(fmt.Errorf("execute index template: %v", err))
	}

	articleTemplatePath := joinPath(baseDir, "gengo", "article.html")
	articleTemplate, err := template.New("article").Parse(mustRead(articleTemplatePath))
	if err != nil {
		panic(fmt.Errorf("parse article template: %v", err))
	}
	for _, article := range articles {
		articleFilePath := joinPath(articlesDir, article.FileName("html"))
		articleFile, err := os.Create(articleFilePath)
		if err != nil {
			panic(fmt.Errorf("open %s: %v", indexFilePath, err))
		}
		defer articleFile.Close()
		err = articleTemplate.Execute(articleFile, article)
		if err != nil {
			panic(fmt.Errorf("execute article template on %s: %v", article.Title, err))
		}
	}
}

func parseArticles(sourceDir, targetDir string) ([]gengo.Article, error) {
	articles := make([]gengo.Article, 0)
	entries, _ := os.ReadDir(sourceDir)
	for _, entry := range entries {
		articleSrcPath := joinPath(sourceDir, entry.Name())
		article, err := parseArticle(articleSrcPath)
		if err != nil {
			return articles, err
		}
		articles = append(articles, *article)
	}
	sort.Sort(sort.Reverse(gengo.Articles(articles)))

	return articles, nil
}

func parseArticle(path string) (*gengo.Article, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open article %s: %v", path, err)
	}
	data, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("read from %s: %v", path, err)
	}
	content := string(data)
	yamlData, mdData, err := splitParts(content)
	if err != nil {
		return nil, fmt.Errorf("parse article %s: %v", path, err)
	}

	var meta gengo.MetaData
	err = yaml.Unmarshal(yamlData, &meta)
	if err != nil {
		return nil, err
	}

	p := parser.New()
	r := html.NewRenderer(html.RendererOptions{})
	htmlContent := string(markdown.ToHTML(mdData, p, r))

	return &gengo.Article{MetaData: meta, HTML: htmlContent}, nil
}

func splitParts(content string) ([]byte, []byte, error) {
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
	return yaml.Bytes(), md.Bytes(), nil
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
	os.Mkdir(articleDir, 0750)

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

func mustRead(path string) string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := bytes.NewBufferString("")
	_, err = io.Copy(buf, f)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
