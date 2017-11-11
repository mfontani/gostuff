package httpsha

import (
	"net/http"
	"net/url"
	"regexp"
)

// StripPrefixAndSHA returns a handler that serves HTTP requests by
// removing the prefix, then removing an optional leading SHA (40 hex digits)
// from the request URL's Path, and ultimately invoking the http.Handler h.
// Often useful to ensure assets contained in a public directory can be easily
// cache-busted without having to use querystrings.
// Similar to: http.StripPrefix
func StripPrefixAndSHA(prefix string, h http.Handler) http.Handler {
	re := regexp.MustCompile("^/?([a-fA-F0-9]{40})/")
	return http.StripPrefix(prefix, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		matched := re.MatchString(r.URL.Path)
		if matched {
			newURL := re.ReplaceAllString(r.URL.Path, "")
			r2 := new(http.Request)
			*r2 = *r
			r2.URL = new(url.URL)
			*r2.URL = *r.URL
			r2.URL.Path = newURL
			h.ServeHTTP(w, r2)
		} else {
			h.ServeHTTP(w, r)
		}
	}))
}
