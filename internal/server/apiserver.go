package server

import (
	"CustomServerTemplate/internal/config"
	"CustomServerTemplate/internal/repository"
	"database/sql"
	"flag"
	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var dbCfgPath string

const (
	AdminRole          = "admin"
	ctxKeyToken ctxKey = iota
)

type ctxKey int8

type APIServer struct {
	Config *config.Config
	Router *mux.Router
	db     *repository.DB
}

func New(config *config.Config) *APIServer {
	return &APIServer{
		Config: config,
		Router: mux.NewRouter(),
	}
}

func (server *APIServer) Start() error {
	testRep()
	server.configureRouter()
	return http.ListenAndServe(server.Config.Port, server.Router)
}

func (server *APIServer) configureDB() error {
	server.db = repository.NewDB(server.Config)
	return nil
}

func testRep() {
	dbData := repository.DBData{}

	flag.StringVar(&dbCfgPath, "dbcfg", "config/dbConfig.toml", "ss")
	_, err := toml.DecodeFile(dbCfgPath, &dbData)
	sqldatabase, err := sql.Open("mysql", "root:1234@/"+dbData.TableName)
	database := repository.DB{}
	database.ConDB = sqldatabase

	rep := &repository.UserRepository{}
	rep.Configure(&database)

	if err != nil {
		log.Fatalln(err)
	}

	rep.Create(dbData)
}
