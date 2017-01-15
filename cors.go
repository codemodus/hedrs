package hedrs

import (
	"fmt"
	"net/http"
)

// CORS ...
func CORS(origs Allower, hdrs fmt.Stringer) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodOptions {
				next.ServeHTTP(w, r)
				return
			}

			o := r.Header.Get(Origin)

			if !origs.IsAllowed(o) {
				stts := http.StatusForbidden
				http.Error(w, http.StatusText(stts), stts)
				return
			}

			w.Header().Set(AccessControlAllowOrigin, o)
			w.Header().Set(AccessControlAllowHeaders, hdrs.String())

			next.ServeHTTP(w, r)
		})
	}
}