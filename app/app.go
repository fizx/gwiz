package app

import (
	"fmt"
	pb "github.com/fizx/gwiz/generated/go/proto"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

type App struct {
	config *pb.AppConfig
}

func (app *App) Start() {
	log.Println(app.config)
}

func NewApp(name string) *App {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	config := &pb.AppConfig{}
	viper.SetConfigName("config")
	viper.AddConfigPath(fmt.Sprintf("/etc/%s/", name))
	viper.AddConfigPath(fmt.Sprintf("$HOME/.%s", name))
	viper.AddConfigPath(".")
	viper.AddConfigPath(filepath.Dir(os.Args[0]))
	if err := viper.ReadInConfig(); err != nil {
		log.Panicln("can't read config file:", err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("cannot parse config: %v", err)
	}
	log.Println("parsed config")
	return &App{config}
}
