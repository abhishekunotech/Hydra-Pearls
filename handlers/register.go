package handlers

import(
    "fmt"
    "net/http"
    "github.com/Unotechsoftware/pearly-gates-go/utils"
)

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
        r.ParseForm()


        username := r.Form["Username"][0]
        password := r.Form["Password"][0]
        email := r.Form["Email"][0]
        component_slice := r.Form["Component"]

    successVal := true

forLoop:        for _, vals := range component_slice {
                returnStrg := utils.CreateUser(username, password, email, vals)

        if returnStrg != "User Registered Successfully"{
            successVal = false
            break forLoop
        }
        }

    if successVal {
              http.Redirect(w, r, "http://192.168.2.166/Hydra-Pearls/login.html", http.StatusSeeOther)
    } else {
        fmt.Fprintf(w,"Error Occured in Creation of User\n<a href=\"http://192.168.2.166/register.html\">Return to Registration Page</a>")
    }

}

