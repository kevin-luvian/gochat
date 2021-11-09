package sql

import (
	"database/sql"
	"fmt"
	"gochat/lib/database/sql/query"
	"strconv"
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

func MakeSQLDB(
	driver string,
	host string,
	port string,
	user string,
	password string,
	dbname string,
) SQLDB {
	iport, err := strconv.Atoi(port)
	if err != nil {
		logrus.Panic("Port must be an integer")
	}
	return SQLDB{
		driver:   driver,
		host:     host,
		port:     iport,
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

func (s *SQLDB) Ping() bool {
	if err := s.database.Ping(); err != nil {
		return false
	}
	return true
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

func (s *SQLDB) Close() {
	if err := s.database.Close(); err != nil {
		logrus.Panic("Cant close connection ", err)
	}

	logrus.Info(s.driver, " db closed!")
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
		case strings.Contains(strings.ToLower(err.Error()), "already exists"):
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
