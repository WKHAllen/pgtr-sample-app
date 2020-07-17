package services

import (
	"main/src/db"
)

var dbm *db.DBManager

// SetDBManager sets the package's global database manager
func SetDBManager(dbManager *db.DBManager) {
	dbm = dbManager
}
