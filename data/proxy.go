package data

import (
	"fmt"
	"net/url"
)

func ExtractImageURL(image string) string {
	u, err := url.Parse(image)
	if err != nil {
		return ""
	}

	// assets.genius.com placeholders can't go through the image proxy.
	if u.Host == "assets.genius.com" {
		return "/static/default_cover_image.png"
	}

	return fmt.Sprintf("/images%s", u.Path)
}
