package utils


import(
	"github.com/antigloss/go/logger"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"crypto/md5"
	"io"
	"encoding/hex"
)

    func GetAccessCount(username string, password string) int{
        db, err := sql.Open("mysql", "root:redhat@/access_manager")
	if err != nil{
		logger.Error(err.Error())
	}	
	h := md5.New()
	io.WriteString(h, password)
	encr_passwd := 	h.Sum(nil)
	enc_pwd_strg := hex.EncodeToString(encr_passwd)
        // insert
        stmt, err := db.Prepare("SELECT count(*) AS Count_records FROM Users WHERE Username=? AND Password=?")
        if err != nil{
	logger.Error(err.Error())
	}
        res, err := stmt.Query(username,enc_pwd_strg)
	if err != nil{
	logger.Error(err.Error())
	}
	var count_rows int
	count_rows = 0	
	for res.Next(){
	err = res.Scan(&count_rows)
	if err != nil{
		logger.Error(err.Error())
	}
	}
	return count_rows
}

func DBConnect()(*sql.DB, error){
     db, err := sql.Open("mysql", "root:redhat@/access_manager")
    return db, err

}


func DBInsert(db *sql.DB, query string)(sql.Result, error){

    stmt, err := db.Exec(query)
    
    return stmt, err
}


func DBSelect(db *sql.DB,query string) (*sql.Rows, error){
    stmt, err := db.Prepare(query)
    
    if err != nil{
        return nil,err
    }
    
    res , err := stmt.Query()
    return res, err
}

