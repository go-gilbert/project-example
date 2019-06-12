package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/sirupsen/logrus"
	"github.com/go-gilbert/project-example/server/config"
	"github.com/go-gilbert/project-example/server/feed/sources"
	"github.com/go-gilbert/project-example/server/web"
)

const defaultConfigFile = "config.json"

var (
	version = "dev"
	commit  = "local build"
)

func main() {
	fmt.Printf("FeedViewer version %s (%s)\n", version, commit)
	if err := bootstrap(); err != nil {
		logrus.Fatalf("failed to start service, %s", err)
		return
	}

	// Handle ^C signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	for range c {
		logrus.Info("shutting down service...")
		exit(0)
	}
}

func exit(code int) {
	if err := web.Shutdown(); err != nil {
		logrus.Errorf("HTTP service stop with error: %s", err)
	}

	sources.Dispose()
	os.Exit(code)
}

func bootstrap() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s", r)
		}
	}()

	cfgFlag := flag.String("config", defaultConfigFile, "Configuration file path")
	flag.Parse()

	if err := config.Load(*cfgFlag); err != nil {
		return err
	}

	if config.Main.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: false,
		ForceColors:   true,
	})

	if err := sources.Load(config.Main.Sources); err != nil {
		return fmt.Errorf("failed to load source libraries: %s", err)
	}

	go func() {
		if err := web.Load(); err != nil {
			exit(1)
			panic(err)
		}
	}()

	return nil
}
