# gostuff

This repository is likely to end up containing some `go` stuff I wrote, found
useful, and want to keep as a separate package for reuse.

Maybe others will find it useful, too. Maybe they won't.

```
$ go get -u github.com/mfontani/gostuff
```

## httpsha

Usage:

```
import "github.com/mfontani/gostuff/httpsha"
// ...
http.Handle("/static/", httpsha.StripPrefixAndSHA("/static/", http.FileServer(http.Dir("./public/"))))
```

Similar to:

```
http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public/"))))
```

... but also removes an optional `SHA` (i.e. 40 hex digits) after the `prefix`,
so that all requests like:

* `/static/foo.css`
* `/static/1234567890123456789012345678901234567890/foo.css`
* `/static/f00af595344582eb39d6f004bc4e01b6175cbf00/foo.css`

... all reply with the exact same "static" file.

This is useful when cache busting CSS or JS, without using a querystring.
