package handles

import (
	"fmt"
	"net/http"
)

// Index Function to provide an entrance to the application.
func IndexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Charlotte: A Web SDK.")
}

// Page Returns a page when a URL for a page is requested.
func AboutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Charlotte: About the webs we weave.")
}