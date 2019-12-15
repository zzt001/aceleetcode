package db

import (
	"database/sql"
	"os"
	"path"

	_ "github.com/mattn/go-sqlite3"
	"github.com/zzt001/aceleetcode/data/entity"
	gorp "gopkg.in/gorp.v1"
)

const DB_FILE = `db.bin`

func NewDB(dataFolder string) (*gorp.DbMap, error) {
	if len(dataFolder) == 0 {
		dataFolder = "./data"
	}
	// create data folder if not exist
	if _, err := os.Stat(dataFolder); os.IsNotExist(err) {
		err := os.Mkdir(dataFolder, 0755)
		if err != nil {
			return nil, err
		}
	}

	db, err := sql.Open("sqlite3", path.Join(dataFolder, DB_FILE))
	if err != nil {
		return nil, err
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	// add tables here
	err = AddProblemTable(dbmap)
	if err != nil {
		return nil, err
	}
	return dbmap, nil
}

func AddProblemTable(dbMap *gorp.DbMap) error {
	if dbMap != nil {
		dbMap.AddTableWithName(entity.Problem{}, "lc_problem").SetKeys(false, "problem_id")
	}
	return dbMap.CreateTablesIfNotExists()
}

func DropProblemTable(dbMap *gorp.DbMap) {
	if dbMap != nil {
		dbMap.DropTableIfExists(entity.Problem{})
	}
}
