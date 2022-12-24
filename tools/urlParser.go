package tools

import (
	"log"
	"net/url"
)

func URLConnection(urlParam string) string {
	url, err := url.Parse(urlParam)
	if err != nil {
		log.Panic(err)
	}
	return url.String()
}
