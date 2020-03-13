package go_hms_mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/timeforaninja/go-hms/sqlUtil"
	"github.com/timeforaninja/go-hms/structs"
)

type MYSQLhandler struct {
	db *sql.DB
}

// InitHandler initialises a new mysql handler
func InitMYSQLHandler() (error, sqlUtil.DBhandler) {
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		return err, nil
	}
	handler := &MYSQLhandler{db}
	handler.validateTables()
	return nil, handler
}

func (mh MYSQLhandler) validateTables() {
	// TODO: write this
}

func (mh MYSQLhandler) getUser(userID uint64) *structs.User {
	return &structs.User{}
}

func (mh MYSQLhandler) setUser(user *structs.User) {

}
