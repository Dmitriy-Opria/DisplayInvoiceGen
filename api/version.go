package api

import (
	"fmt"
	"net/http"

	"github.com/InVisionApp/rye"
)

func (a *Api) versionHandler(rw http.ResponseWriter, r *http.Request) *rye.Response {
	fmt.Fprintf(rw, "invoice_gen: %v", a.Version)
	return nil
}
