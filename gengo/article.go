package gengo

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

type MetaData struct {
	Title    string `yaml:"title"`
	Subtitle string `yaml:"subtitle"`
	Author   string `yaml:"author"`
	RawDate  string `yaml:"date"`
	Language string `yaml:"lang"`
}

func (fm *MetaData) Date() time.Time {
	t, _ := time.Parse("2006-01-02T15:04:05-0700", fm.RawDate)
	return t
}

func (fm *MetaData) FormatDate() string {
	switch fm.Language {
	case "de":
		return fm.Date().Format("02.01.2006")
	default:
		return fm.Date().Format("2006-01-02")
	}
}

func (fm *MetaData) Name() string {
	return normalize(fm.Title)
}

func (fm *MetaData) FileName(extension string) string {
	return fmt.Sprintf("%s-%s.%s", fm.Date().Format("2006-01-02"), fm.Name(), extension)
}

func (fm *MetaData) Href(dir string) string {
	return strings.Join([]string{dir, fm.FileName("html")}, string(os.PathSeparator))
}

type Article struct {
	MetaData
	HTML string
}

func (a Article) PublishedFmt() string {
	switch a.Language {
	case "de":
		return a.Date().Format("01.02.2006 15:04")
	default:
		return a.Date().Format("2006-02-01 15:04")
	}
}

type Articles []Article

func (a Articles) Len() int           { return len(a) }
func (a Articles) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Articles) Less(i, j int) bool { return a[i].Date().Before(a[j].Date()) }

func normalize(bad string) string {
	good := strings.ReplaceAll(bad, " ", "-")
	good = strings.ReplaceAll(good, "ä", "ae")
	good = strings.ReplaceAll(good, "ö", "oe")
	good = strings.ReplaceAll(good, "ü", "ue")
	nonAlNum := regexp.MustCompilePOSIX("[^-_0-9a-zA-Z]+")
	good = nonAlNum.ReplaceAllString(good, "")
	return strings.ToLower(good)
}
