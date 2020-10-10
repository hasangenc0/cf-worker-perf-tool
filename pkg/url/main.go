package url

import (
	"net/url"
)

func IsValidUrl(fullUrl string) bool {
	_, err := url.ParseRequestURI(fullUrl)
	return err == nil
}
