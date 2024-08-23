package main
import (
	"os"
	"fmt"
	)


func main() {

	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	rawBaseURL := os.Args[1] 

	fmt.Printf("starting crawl of: %v", rawBaseURL)

	htmlBody, err := getHTML(rawBaseURL)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(htmlBody)
}
