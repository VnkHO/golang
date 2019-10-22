package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://www.google.fr/",
		"https://golang.org/",
		"https://www.twitch.tv/",
		"https://stackoverflow.com/",
	}

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down !")
		return
	}
	fmt.Println(link, "is up!")
}
