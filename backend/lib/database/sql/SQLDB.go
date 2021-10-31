package sql

import (
	"database/sql"
	"fmt"
	"gochat/lib/database/sql/query"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	"github.com/sirupsen/logrus"
)

type SQLDB struct {
	driver   string
	host     string
	port     int
	user     string
	password string
	dbname   string
	database *sql.DB
}

func MakeTestSQLDB() SQLDB {
	return SQLDB{
		driver:   "mysql",
		host:     "localhost",
		port:     3301,
		user:     "gouser",
		password: "gopassword",
		dbname:   "gomysqltest"}
}

func MakeDefaultSQLDB() SQLDB {
	return SQLDB{
		driver:   "mysql",
		host:     "localhost",
		port:     5444,
		user:     "gouser",
		password: "gopassword",
		dbname:   "gomysql"}
}

func MakeSQLDB(
	driver string,
	host string,
	port int,
	user string,
	password string,
	dbname string,
) SQLDB {
	return SQLDB{
		driver:   driver,
		host:     host,
		port:     port,
		user:     user,
		password: password,
		dbname:   dbname}
}

func (s *SQLDB) getInfoString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", s.user, s.password, s.host, s.port, s.dbname)
}

func (s *SQLDB) GetDatabase() *sql.DB {
	return s.database
}

func (s *SQLDB) Connect() {
	db, err := sql.Open(s.driver, s.getInfoString())
	if err != nil {
		logrus.Panic("Cant open db connection \n", err)
	}
	s.database = db

	if err := db.Ping(); err != nil {
		logrus.Panic("Connection is dead \n", err)
	}

	logrus.Info(s.driver, " db connected!")
}

func (s *SQLDB) DropCreateTables(models ...interface{}) {
	s.DropTables(models...)
	s.CreateTables(models...)
}

func (s *SQLDB) DropTables(models ...interface{}) {
	for _, model := range models {
		s.dropTable(model)
	}
}

func (s *SQLDB) CreateTables(models ...interface{}) {
	for _, model := range models {
		s.createTable(model)
	}
}

func (s *SQLDB) createTable(o interface{}) {
	mmeta := query.MakeModelMetadata(o)
	q := query.MakeCreateTableQuery(o)
	if _, err := s.database.Exec(q); err != nil {
		switch {
		case strings.Contains(err.Error(), "already exists"):
			logrus.Warn("CREATE TABLE FAILED: table ",
				mmeta.Tablename, " already exists in database")
		default:
			logrus.Panic("unknown error ", err)
		}
	} else {
		logrus.Info("Table ", mmeta.Tablename, " created")
	}
}

func (s *SQLDB) dropTable(o interface{}) {
	mmeta := query.MakeModelMetadata(o)
	q := query.MakeDropTableQuery(o)
	if _, err := s.database.Exec(q); err != nil {
		switch {
		case strings.Contains(strings.ToLower(err.Error()), "unknown table"):
			logrus.Warn("DROP TABLE FAILED: table ",
				mmeta.Tablename, " doesn't exists in database")
		default:
			logrus.Panic("DROP TABLE FAILED: unknown error ", err)
		}
	} else {
		logrus.Warn("Table ", mmeta.Tablename, " dropped")
	}
}
