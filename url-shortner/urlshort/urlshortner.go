package urlshort

import "net/http"

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if val, ok := pathsToUrls[r.URL.Path]; ok {
			http.Redirect(w, r, val, http.StatusFound)
		}

		fallback.ServeHTTP(w, r)
	}
}
