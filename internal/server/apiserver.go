package server

import (
	"CustomServerTemplate/internal/config"
	"CustomServerTemplate/internal/repository"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
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
	//testRep()
	server.configureRouter()
	err := server.configureDB()
	if err != nil {
		return err
	}
	return http.ListenAndServe(server.Config.Port, server.Router)
}

func (server *APIServer) configureDB() error {
	db := repository.NewDB(server.Config)
	err := db.Open()
	if err != nil {
		return err
	}
	server.db = db
	server.db.CreateRepositories()
	return nil
}

/*func testRep() {
	dbData := repository.DBData{}

	flag.StringVar(&dbCfgPath, "dbcfg", "config/dbConfig.toml", "ss")
	_, err := toml.DecodeFile(dbCfgPath, &dbData)
	sqldatabase, err := sql.Open("mysql", "root:1234@/"+dbData.TableName)
	database := repository.DB{}
	database.conDB = sqldatabase

	rep := &repository.UserRepository{}
	rep.Configure(&database)

	if err != nil {
		log.Fatalln(err)
	}

	err = rep.Initialize(dbData)
	if err != nil {
		logrus.Error(err)
	}
} */
