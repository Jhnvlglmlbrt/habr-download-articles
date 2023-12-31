package main

import (
	"bufio"
	"fmt"
	"os"

	doc "habr-download-articles/doc"

	"github.com/gocolly/colly"
)

func main() {
	var URL string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Введите URL: ")
	scanner.Scan()

	URL = scanner.Text()

	c := colly.NewCollector(
		colly.AllowedDomains("habr.com"),
	)

	var headerText string
	var articleText string

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36")
		// fmt.Printf("Посещение - %s\n\n", r.URL)
	})

	c.OnHTML(".tm-title.tm-title_h1", func(h *colly.HTMLElement) {
		headerText = h.Text
	})

	c.OnHTML(".tm-article-body", func(h *colly.HTMLElement) {
		articleText = h.Text
	})

	c.OnScraped(func(r *colly.Response) {
		if headerText != "" && articleText != "" {
			fmt.Print("Вы уверены, что хотите перезаписать файлы? (Yes/no): ")
			scanner.Scan()
			answer := scanner.Text()

			if answer == "Yes" {
				doc.SaveToTextFile(headerText, articleText)
				fmt.Println("Данные успешно сохранены в .docx файл.")
			} else {
				fmt.Println("Операция отменена.")
			}
		} else {
			fmt.Println("He удалось извлечь необходимые данные.")
		}
	})

	c.Visit(URL)
}
