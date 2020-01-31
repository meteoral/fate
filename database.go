package fate

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/godcong/fate/config"
	"github.com/goextension/log"

	"github.com/xormsharp/xorm"
	"net/url"
)

const mysqlSource = "%s:%s@tcp(%s)/%s?loc=%s&charset=utf8mb4&parseTime=true"
const createDatabase = "CREATE DATABASE `%s` CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_bin'"

var _ = mysql.Config{}
var DefaultTableName = "fate"

type Database interface {
	Sync(v ...interface{}) error
	CountWuGeLucky() (n int64, e error)
	InsertOrUpdateWuGeLucky(lucky *WuGeLucky) (n int64, e error)
	GetCharacter(fn func(engine *xorm.Engine) *xorm.Session) (*Character, error)
	GetCharacters(fn func(engine *xorm.Engine) *xorm.Session) ([]*Character, error)
	GetPoems(fn func(engine *xorm.Engine) *xorm.Session) ([]*Poem, error)
	FilterWuGe(last []*Character, wg chan<- *WuGeLucky) error
	Database() interface{}
}

type xormDatabase struct {
	*xorm.Engine
}

func (db *xormDatabase) Database() interface{} {
	return db.Engine
}

func (db *xormDatabase) FilterWuGe(last []*Character, wg chan<- *WuGeLucky) error {
	return filterWuGe(db.Engine, last, wg)
}

func (db *xormDatabase) GetCharacters(fn func(engine *xorm.Engine) *xorm.Session) ([]*Character, error) {
	return getCharacters(db.Engine, fn)
}

func (db *xormDatabase) Sync(v ...interface{}) error {
	return db.Engine.Sync2(v...)
}

func (db *xormDatabase) GetCharacter(fn func(engine *xorm.Engine) *xorm.Session) (*Character, error) {
	return getCharacter(db.Engine, fn)
}

func (db *xormDatabase) GetPoems(fn func(engine *xorm.Engine) *xorm.Session) ([]*Poem, error) {
	return getPoems(db.Engine, fn)
}

func (db *xormDatabase) InsertOrUpdateWuGeLucky(lucky *WuGeLucky) (n int64, e error) {
	return insertOrUpdateWuGeLucky(db.Engine, lucky)
}

func (db *xormDatabase) CountWuGeLucky() (n int64, e error) {
	return countWuGeLucky(db.Engine)
}

func InitDatabaseFromConfig(config config.Config) Database {
	engine := initSQL(config.Database)
	return &xormDatabase{
		Engine: engine,
	}
}

func initSQL(database config.Database) *xorm.Engine {
	dbURL := fmt.Sprintf(mysqlSource, database.User, database.Pwd, database.Addr(), "", url.QueryEscape("Asia/Shanghai"))
	dbEngine, e := xorm.NewEngine(database.Driver, dbURL)
	if e != nil {
		log.Panicw("connect database failed", "error", e)
	}
	defer dbEngine.Close()

	sql := fmt.Sprintf(createDatabase, database.Name)

	_, e = dbEngine.DB().Exec(sql)
	if e == nil {
		log.Infow("create database failed", "database", DefaultTableName)
	}
	u := fmt.Sprintf(mysqlSource, database.User, database.Pwd, database.Addr(), database.Name, url.QueryEscape("Asia/Shanghai"))
	eng, e := xorm.NewEngine("mysql", u)
	if e != nil {
		log.Panicw("connect table failed", "error", e)
	}
	if database.ShowSQL {
		eng.ShowSQL(true)
	}

	if database.ShowExecTime {
		eng.ShowExecTime(true)
	}
	return eng
}
