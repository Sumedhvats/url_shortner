package urlHandler

import "net/http"

func MapHandler(m map[string]string, mux *http.ServeMux) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := m[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		mux.ServeHTTP(w, r)
	}
}

func YAMLHandler(b []byte, fallback http.Handler) (http.HandlerFunc, error) {
	return nil, nil
}
