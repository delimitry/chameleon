package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// Article is a strct with article title and content
type Article struct {
	Category
	File  string
	Title string
	Raw   string
}

func (article *Article) getRaw() (string, error) {
	link, err := article.makeLink(rawLinkT)
	if err != nil {
		return "", err
	}
	link += "/" + article.File
	res, err := http.Get(link)
	if err != nil {
		return "", err
	}
	content, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (article *Article) getTitle() string {
	title := strings.Split(article.Raw, "\n")[0]
	if strings.Index(title, "# ") != 0 {
		return article.File
	}
	return title[2:]
}
