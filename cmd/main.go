package main

import (
	"HumoAcademy"
	"HumoAcademy/pkg/handler"
	"HumoAcademy/pkg/repository"
	"HumoAcademy/pkg/service"
	"HumoAcademy/schema"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
)

func main() {
	//logrus.SetFormatter(new(logrus.JSONFormatter)) //ошибки будут иметь формат json
	initLogs()
	if err := initConfig(); err != nil {
		log.Fatalf("error while reading config file. Error is %s", err.Error())
	}

	database, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.password"), //TODO: Почему если подключить env то sslmode выдаёт ошибку ?
	})

	if err != nil {
		log.Fatalf("error while initializing schema. Error is %s", err.Error())
	}

	//schema.DBDrop(database)
	schema.DBInit(database)

	fmt.Println("server is listening port 8181")
	log.Println("server is listening port 8181")
	repos := repository.NewRepository(database)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(HumoAcademy.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error while running http server. Error is %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func initLogs() {
	log.SetOutput(&lumberjack.Logger{
		Filename:   "logs/foo.log",
		MaxSize:    viper.GetInt(""), // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
		Compress:   true, // disabled by default
	})
}