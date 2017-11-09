package handlers


import (
    "fmt"
"github.com/antigloss/go/logger"
 _ "github.com/go-sql-driver/mysql"
    "net/http"
    "github.com/Unotechsoftware/pearly-gates-go/utils"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
        r.ParseForm()
        username := r.Form["Username"][0]
        password := r.Form["Password"][0]

        db, err := utils.DBConnect()
        if err != nil {
                fmt.Println(err.Error())
        } else {



// SELECT count(*) as count_users FROM Users WHERE Username= "123" AND Password="1242421412424de2e12"
                sqlStrng := "SELECT count(*) as count_users FROM Users WHERE Username= \"" + username + "\" AND Password=\"" + utils.CreateMD5Password(password) + "\""
                rows, err := utils.DBSelect(db, sqlStrng)
                if err != nil {
                        logger.Error(err.Error())
                }
                var Count_records int
                Count_records = 0
                for rows.Next() {
                        err = rows.Scan(&Count_records)
                        if err != nil {
                                logger.Error(err.Error())
                        }

                }
                if Count_records == 0 {
                        http.Redirect(w, r, "http://192.168.2.166/Hydra-Pearls/login.html", http.StatusSeeOther)
                } else {

                        http.Redirect(w, r, "http://192.168.2.166/Hydra-Pearls/createAPI.html", http.StatusSeeOther)
                }

        }

        db.Close()
}
