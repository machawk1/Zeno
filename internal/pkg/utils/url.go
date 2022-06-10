package utils

import (
	"errors"
	"net/url"

	"github.com/asaskevich/govalidator"
)

// MakeAbsolute turn all URLs in a slice of url.URL into absolute URLs, based
// on a given base *url.URL
func MakeAbsolute(base *url.URL, URLs []url.URL) []url.URL {
	for i, URL := range URLs {
		if URL.IsAbs() == false {
			URLs[i] = *base.ResolveReference(&URL)
		}
	}

	return URLs
}

func RemoveFragments(URLs []url.URL) []url.URL {
	for i := range URLs {
		URLs[i].Fragment = ""
	}

	return URLs
}

// DedupeURLs take a slice of *url.URL and dedupe it
func DedupeURLs(URLs []url.URL) []url.URL {
	keys := make(map[string]bool)
	list := []url.URL{}
	for _, entry := range URLs {
		if _, value := keys[entry.String()]; !value {
			keys[entry.String()] = true

			if entry.Scheme == "http" || entry.Scheme == "https" {
				list = append(list, entry)
			}
		}
	}
	return list
}

// ValidateURL validates a *url.URL
func ValidateURL(u *url.URL) error {
	valid := govalidator.IsURL(u.String())

	if u.Scheme != "http" && u.Scheme != "https" {
		valid = false
	}

	if valid == false {
		return errors.New("Not a valid URL")
	}

	return nil
}
