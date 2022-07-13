package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"tdd/config"
	"tdd/dao"
	"tdd/web"
)

func getConfigPath() string {
	const defaultConfigPath = "/etc/tdd/config.yml"

	configPath := ""
	flag.StringVar(&configPath, "config", defaultConfigPath, "config file path")
	flag.StringVar(&configPath, "c", defaultConfigPath, "config file path (shorthand)")

	flag.Parse()

	return configPath

}

func waitExit() {
	exitChan := make(chan os.Signal)
	signal.Notify(exitChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	<-exitChan
}

func main() {
	err := config.Init(getConfigPath())
	if err != nil {
		log.Fatalf("init config error %v\n", err)
	}

	log.Println("init database...")
	err = dao.Init(config.Get().Database)
	if err != nil {
		log.Fatalf("init dao error %v\n", err)
	}
	defer dao.Close()

	log.Printf("start web server")
	webServer := web.NewServer(config.Get().Web)
	if webServer == nil {
		log.Fatalln("new web server error")
	}
	webServer.Start()
	defer webServer.Stop()

	waitExit()
}
