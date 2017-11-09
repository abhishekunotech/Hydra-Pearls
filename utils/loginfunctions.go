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


func CheckIfRegistered(username string, password string) bool {

	db, err := DBConnect()
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	encPassword := CreateMD5Password(password)
	sqlStrng := "SELECT count(*) AS Count_records FROM Users WHERE Username=\"" + username + "\"AND Password=\"" + encPassword + "\""
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
