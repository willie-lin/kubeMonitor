package nextserver

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/migrate"
	"go.uber.org/zap"
	//"log"
)

//type DatabaseCfg struct {
//	User     string `json:"user"`
//	Password string `json:"password"`
//	DbName   string `json:"db_name"`
//	Host     string `json:"host"`
//	Port     int    `json:"port"`
//	//DbPath   string      `json:"db_path"`
//	Type	 string `json:"type"`
//	Sslmode  string `json:"sslmode"`
//	//Typea     string     `json:"typea"`
//
//}



func (s *NextServer)  NewClient() (*ent.Client, error) {

	var client *ent.Client
	var err error
	dbCfg := s.config.Database
	switch dbCfg.Type {
	case "sqlite3":
		client, err := ent.Open(dbCfg.Type, fmt.Sprintf("file:%s?_busy_timeout=10000)&_fk=1", dbCfg.DbName))
		if err != nil {
			return client, fmt.Errorf("failed opening connection to sqlite: %v", err)
		}
	case "mysql":
		client, err := ent.Open(dbCfg.Type, fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			dbCfg.Host, dbCfg.Port, dbCfg.User, dbCfg.Password, dbCfg.DbName, dbCfg.SslMode))
		if err != nil {
			return client, fmt.Errorf("failed opening connection to mysql: %v", err)
		}
	case "postgres", "postgresql":
		client, err := ent.Open(dbCfg.Type, fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			dbCfg.Host, dbCfg.Port, dbCfg.User, dbCfg.Password, dbCfg.DbName, dbCfg.SslMode))
		if err != nil {
			return client, fmt.Errorf("failed opening connection to postgres: %v", err)
		}
	default:
		return client, fmt.Errorf("unknown database type")
	}
	return client, err
}



//
//func NewClient() (*ent.Client, error) {
//	var dfg = &DatabaseCfg{
//		User:     "root",
//		Password: "root1234",
//		DbName:   "terminal",
//		//Host:     "db",
//		Host: "127.0.0.1",
//		Port: 3306,
//		Type: "mysql",
//		//Type: "sqlite3",
//	}
//	var client *ent.Client
//	var err error
//	//drv, err := sql.Open("mysql", "root:root1234@tcp(127.0.0.1:3306)/terminal?charset=utf8&parseTime=true")
//	switch dfg.Type {
//	case "sqlite3":
//		client, err = ent.Open(dfg.Type, fmt.Sprintf("file:%s?_busy_timeout=100000&_fk=1", dfg.DbName))
//		if err != nil {
//			return client, fmt.Errorf("failed opening connection to sqlite: %v", err)
//		}
//	case "mysql":
//		client, err = ent.Open(dfg.Type, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",
//			dfg.User, dfg.Password, dfg.Host, dfg.Port, dfg.DbName))
//		if err != nil {
//			return client, fmt.Errorf("failed opening connection to mysql: %v", err)
//		}
//	case "postgres", "postgresql":
//		client, err = ent.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
//			dfg.Host, dfg.Port, dfg.User, dfg.DbName, dfg.Password))
//		if err != nil {
//			return client, fmt.Errorf("failed opening connection to postgres: %v", err)
//		}
//	default:
//		return client, fmt.Errorf("unknown database type")
//	}
//	return client, err
//}

// AutoMigration auto migrate
func AutoMigration(client *ent.Client, ctx context.Context) {
	log, _ := zap.NewDevelopment()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatal("failed creating schema resources: %v", zap.Error(err))
		//log.Fatalf("failed creating schema resources: %v", err)
	}
}

// DebugMode debug mode
func DebugMode(err error, client *ent.Client, ctx context.Context) {
	log, _ := zap.NewDevelopment()
	err = client.Debug().Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatal("failed creating schema resources: %v", zap.Error(err))
		//log.Fatalf("failed creating schema resources: %v", err)
	}
}


