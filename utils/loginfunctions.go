package utils

import (
	"fmt"
	"github.com/antigloss/go/logger"
	//"github.com/Unotechsoftware/pearly-gates-go/redis"
	_ "github.com/go-sql-driver/mysql"
	//"html/template"
	"crypto/md5"
	"encoding/hex"
	"io"
)

/*
func createAPI(w http.ResponseWriter, r *http.Request){
    r.ParseForm()
    fmt.Println(r.Form)
    APIName := r.Form["APIName"][0]
    Version :=  r.Form["Version"][0]
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

}

func editAPI(w http.ResponseWriter, r *http.Request){
    //r.ParseForm()


}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form["Username"][0]
	password := r.Form["Password"][0]
	db, err := sql.Open("mysql", "root:redhat@/access_manager")
	if err != nil {
		logger.Error(err.Error())
	}

	stmt, err := db.Prepare("SELECT count(*) as count_users FROM Users WHERE Username=? AND Password=?")
	if err != nil {
		logger.Error(err.Error())
	}

	h := md5.New()
	io.WriteString(h, password)
	encr_passwd := h.Sum(nil)
	enc_pwd_strg := hex.EncodeToString(encr_passwd)

	rows, err := stmt.Query(username, enc_pwd_strg)
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
		http.Redirect(w, r, "http://192.168.2.166/HydraHtml/login.html", http.StatusSeeOther)
	} else {

		http.Redirect(w, r, "http://192.168.2.166/HydraHtml/createAPI.html", http.StatusSeeOther)
	}

    db.Close()
}

func main() {
	mux := http.NewServeMux()
	th_basic := http.HandlerFunc(login)
	hf_register := http.HandlerFunc(register)
    hf_createAPI := http.HandlerFunc(createAPI)
	mux.Handle("/login", th_basic)
    mux.Handle("/createAPI",hf_createAPI)
	mux.Handle("/register", hf_register)
	log.Fatal(http.ListenAndServe(":8012", mux))

}

func register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	db, err := sql.Open("mysql", "root:redhat@/access_manager")
	if err != nil {
		logger.Error(err.Error())
	}

	username := r.Form["Username"][0]
	password := r.Form["Password"][0]
	email := r.Form["Email"][0]
	component_slice := r.Form["Component"]

	var Count_records int

	stmt, err := db.Prepare("SELECT count(*) AS Count_records FROM Users WHERE Username=? AND Password=?")
	if err != nil {
		logger.Error(err.Error())
	}

	h := md5.New()
	io.WriteString(h, password)
	encr_passwd := h.Sum(nil)
	enc_pwd_strg := hex.EncodeToString(encr_passwd)

	res, err := stmt.Query(username, enc_pwd_strg)
	if err != nil {
		logger.Error(err.Error())
	}
	Count_records = 0
	for res.Next() {
		err = res.Scan(&Count_records)
		if err != nil {
			logger.Error(err.Error())
		}
	}

	if Count_records == 0 {
		for _, vals := range component_slice {
			_, err = db.Exec("INSERT INTO Users(username, password,email,component,create_time,update_time) VALUES(?,?,?,?,CURDATE(),CURDATE())", username, enc_pwd_strg, email, vals)
			if err != nil {
				logger.Error(err.Error())
				fmt.Println("Insertion failed")
			}

			fmt.Println("Record created!")
		}
	}
}*/

func CheckIfRegistered(username string, password string) bool {

	db, err := DBConnect()
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	encPassword := CreateMD5Password(password)
	sqlStrng := "SELECT count(*) AS Count_records FROM Users WHERE Username=" + username + "AND Password=" + encPassword
	rows, err := DBSelect(db, sqlStrng)
	db.Close()
	if err != nil {
		logger.Error(err.Error())
		return false
	}

	Count_records := 0
	for rows.Next() {
		err = rows.Scan(&Count_records)
		if err != nil {
			logger.Error(err.Error())
			return false
		}
	}
	if Count_records == 0 {
		return false
	} else {
		return true
	}

}

func CreateMD5Password(password string) string {

	h := md5.New()
	io.WriteString(h, password)
	encr_passwd := h.Sum(nil)
	enc_pwd_strg := hex.EncodeToString(encr_passwd)
	return enc_pwd_strg

}

func CreateUser(username string, password string, component string, email string) string {

	if !CheckIfRegistered(username, password) {

		db, err := DBConnect()
        

        password = CreateMD5Password(password)
		sqlStrng := "INSERT INTO Users(username, password,email,component,create_time,update_time) VALUES(\"" + username + "\",\"" + password + "\",\"" + email + "\",\"" + component + "\",CURDATE(),CURDATE())"

		_, err = DBInsert(db, sqlStrng)
		if err != nil {
			return err.Error()
		}
        return "User Registered Successfully"
	} else {

		return "User Already Registered"
	}

}
