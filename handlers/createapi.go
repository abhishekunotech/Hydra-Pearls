package handlers

import (
    "fmt"
 _ "github.com/go-sql-driver/mysql"
    "net/http"
    "github.com/Unotechsoftware/pearly-gates-go/redis"
)

func (h *Handler) CreateAPI (w http.ResponseWriter, r *http.Request) {
        r.ParseForm()
        APIName := r.Form["APIName"][0]
        Version := r.Form["Version"][0]
        Method := r.Form["Method"][0]
        HydraURI := r.Form["HydraURI"][0]
        Component := r.Form["Component"][0]
        Handler := r.Form["Handler"][0]
        ComponentURI := r.Form["ComponentURI"][0]

        APIValues := make(map[string]interface{})
        APIValues["APIName"] = APIName
        APIValues["Version"] = Version
        APIValues["Method"] = Method
        APIValues["HydraURI"] = HydraURI
        APIValues["Component"] = Component
        APIValues["Handler"] = Handler
        APIValues["ComponentURI"] = ComponentURI

        redis.DBConnect(APIName, APIValues)
    fmt.Fprintf(w,"API Created Successfully")
}

