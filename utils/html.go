package utils

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func GetQueryContentFromUrl(url string, query string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	// close io
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("connection failed: %s", url)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		return "", fmt.Errorf("error to convert content of <%s> to document query", url)
	}

	var value string = ""

	doc.Find(query).Each(func(i int, s *goquery.Selection) {
		value = s.Text()
	})

	value = ReplacePersianNumbersToEnglishNumbers(value, ",", "")

	parsedValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return "", fmt.Errorf("error to convert <%s> of <%s> in <%s> to float4", value, query, url)
	}

	formatValue := strconv.FormatFloat(parsedValue, 'f', -1, 64)

	return formatValue, nil
}
