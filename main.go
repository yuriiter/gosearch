package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type PackageInfo struct {
	Path       string
	Synopsis   string
	Version    string
	Published  string
	ImportedBy string
	License    string
	URL        string
}

const (
	ColorReset  = "\033[0m"
	ColorBold   = "\033[1m"
	ColorCyan   = "\033[36m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorGray   = "\033[90m"
)

func main() {
	limit := flag.Int("limit", 10, "Max number of results to display")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: gosearch [flags] <query>")
		fmt.Println("Flags:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	query := strings.Join(args, " ")

	baseURL := "https://pkg.go.dev/search"
	params := url.Values{}
	params.Add("q", query)
	params.Add("limit", fmt.Sprintf("%d", *limit))
	params.Add("m", "package")

	searchURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	fmt.Printf("%sSearching pkg.go.dev for: %s%s\n", ColorGray, query, ColorReset)
	fmt.Printf("%sRequest URL: %s%s\n\n", ColorGray, searchURL, ColorReset)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; gopkgsearch/1.0)")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("Failed to fetch data. Status Code: %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var results []PackageInfo

	doc.Find(".SearchSnippet").Each(func(i int, s *goquery.Selection) {
		rawPath := s.Find(".SearchSnippet-header-path").Text()
		cleanPath := strings.Trim(rawPath, "()")

		if cleanPath == "" {
			href, _ := s.Find("h2 a").Attr("href")
			cleanPath = strings.TrimPrefix(href, "/")
		}

		synopsis := strings.TrimSpace(s.Find(".SearchSnippet-synopsis").Text())

		infoLabel := s.Find(".SearchSnippet-infoLabel")

		importedBy := strings.TrimSpace(infoLabel.Find("a[aria-label='Go to Imported By'] strong").Text())

		license := strings.TrimSpace(infoLabel.Find("span[data-test-id='snippet-license']").Text())

		publishedDate := strings.TrimSpace(infoLabel.Find("span[data-test-id='snippet-published'] strong").Text())

		version := ""
		infoLabel.Find("span.go-textSubtle").Each(func(i int, sub *goquery.Selection) {
			if strings.Contains(sub.Text(), "published on") {
				version = sub.Find("strong").First().Text()
			}
		})

		results = append(results, PackageInfo{
			Path:       cleanPath,
			Synopsis:   synopsis,
			Version:    version,
			Published:  publishedDate,
			ImportedBy: importedBy,
			License:    license,
			URL:        "https://pkg.go.dev/" + cleanPath,
		})
	})

	if len(results) == 0 {
		fmt.Println("No results found.")
		return
	}

	for _, r := range results {
		printResult(r)
	}
}

func printResult(p PackageInfo) {
	fmt.Printf("%s%s%s", ColorBold, ColorCyan, p.Path)
	if p.Version != "" {
		fmt.Printf(" %s(%s)%s", ColorGreen, p.Version, ColorReset)
	} else {
		fmt.Print(ColorReset)
	}
	fmt.Println()

	var meta []string
	if p.ImportedBy != "" {
		meta = append(meta, fmt.Sprintf("Imports: %s%s%s", ColorYellow, p.ImportedBy, ColorReset))
	}
	if p.License != "" {
		meta = append(meta, fmt.Sprintf("License: %s%s%s", ColorYellow, p.License, ColorReset))
	}
	if p.Published != "" {
		meta = append(meta, fmt.Sprintf("Updated: %s%s%s", ColorYellow, p.Published, ColorReset))
	}

	if len(meta) > 0 {
		fmt.Printf("  %s\n", strings.Join(meta, " | "))
	}

	if p.Synopsis != "" {
		fmt.Printf("  %s\n", p.Synopsis)
	} else {
		fmt.Printf("  %s(No description available)%s\n", ColorGray, ColorReset)
	}

	fmt.Println()
}
