package main

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main() {

	// проверка что параметр был введен при запуске программы
	if len(os.Args) != 1 {
		_, _ = fmt.Fprintln(os.Stderr, "usage: [url]")
		os.Exit(1)
	}

	// скачивание сайта
	err := wget("web", os.Args[1])
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

// wget скачивает сайт целиком
func wget(dir, url string) error {

	// создание каталога для хранения файлов сайта
	if err := os.MkdirAll(dir, 0777); err != nil {
		return fmt.Errorf("не удалось создать новую дерикторию: %v", err)
	}

	// переход в только что созданый каталог
	if err := os.Chdir(dir); err != nil {
		return fmt.Errorf("не удалось перейти в дерикторию: %v", err)
	}

	// отправка запроса к сайту
	r, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("не удалось выполнить запрос: %v", err)
	}
	defer func() { _ = r.Body.Close() }()

	if r.StatusCode != 200 {
		return fmt.Errorf("status code: %d", r.StatusCode)
	}

	// чтение тела респонса
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("не удалось прочитать респонс: %v", err)
	}

	// создание файла и запись в него данных
	if err = createFile(filepath.Dir(r.Request.URL.Path), "index.html", data); err != nil {
		return err
	}

	// поиск всех ссылок со страницы
	links := getLinks(bytes.NewReader(data))

	for _, link := range links {

		// прообразование к нормальному виду ссылки
		normLink := normalizeLink(link, r.Request.URL.Scheme, r.Request.URL.Host)

		// скачивание данных по ссылке
		if err := downloadLink(normLink); err != nil {
			fmt.Printf("%s: %s\n", link, err.Error())
		}
	}

	return nil
}

// createFile создает файл в папке
func createFile(path, filename string, data []byte) error {

	currentDir, _ := os.Getwd()
	dir := currentDir + path

	if err := os.MkdirAll(dir, 0777); err != nil {
		return fmt.Errorf("не удалось создать папку: %v", err)
	}

	file, err := os.Create(dir + "/" + filename)
	if err != nil {
		return fmt.Errorf("не удалось создать файл: %v", err)
	}

	if _, err = fmt.Fprintln(file, string(data)); err != nil {
		return fmt.Errorf("не удалось записать данные в файл: %v", err)
	}

	return nil
}

// normalizeLink преобразовывает ссылку, добавляя к ней схему и хост, если они отсутствуют
func normalizeLink(link, scheme, host string) string {
	if !strings.HasPrefix(link, "http://") && !strings.HasPrefix(link, "https://") {
		link = scheme + "://" + host + link
	}
	return link
}

// extractLinks возвращает массив ссылок из html.Token
func extractLinks(token html.Token) []string {
	var links []string

	switch token.Data {
	case "link":
		for _, attr := range token.Attr {
			if attr.Key == "href" {
				links = append(links, attr.Val)
			}
		}
	case "script":
		for _, attr := range token.Attr {
			if attr.Key == "src" {
				links = append(links, attr.Val)
			}
		}
	}
	return links
}

// getLinks возвращает массив ссылок из html страницы
func getLinks(r io.Reader) []string {
	var links []string

	tokenizer := html.NewTokenizer(r)

	for {
		token := tokenizer.Next()

		switch token {
		case html.ErrorToken:
			return links
		case html.StartTagToken, html.EndTagToken:
			links = append(links, extractLinks(tokenizer.Token())...)
		}
	}
}

// downloadLink скачивает файл по ссылке
func downloadLink(url string) error {
	r, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("не получилось выполнить запрос: %v", err)
	}
	defer func() { _ = r.Body.Close() }()

	if r.StatusCode != 200 {
		return fmt.Errorf("status code: %d", r.StatusCode)
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("не получилось прочитать респонс: %v", err)
	}

	if err = createFile(
		path.Dir(r.Request.URL.Path),
		path.Base(r.Request.URL.Path),
		data,
	); err != nil {
		return err
	}

	return nil
}
