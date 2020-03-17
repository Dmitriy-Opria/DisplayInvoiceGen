package api

import (
	"fmt"
	"github.com/InVisionApp/rye"
	"net/http"
)

func (a *Api) versionHandler(rw http.ResponseWriter, r *http.Request) *rye.Response {
	fmt.Fprintf(rw, "invoice_gen: %v", a.Version)
	return nil
}
