package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/johannesboyne/gofakes3"
	"github.com/johannesboyne/gofakes3/backend/s3mem"
	"github.com/sculley/someadmin-go/config"
	log "github.com/sirupsen/logrus"
)

type FakeS3 struct {
	General struct {
		Port         int           `mapstructure:"port"`
		ReadTimeout  time.Duration `mapstructure:"read_timeout"`
		WriteTimeout time.Duration `mapstructure:"write_timeout"`
	} `mapstructure:"general"`
}

func main() {
	// Set log level
	setLogLevel()

	// Load config
	var cfg FakeS3

	configLoader := config.FileConfig{
		Path: getEnvVar("CONFIG_PATH", "./conf/"),
		Name: "config",
		Type: "yaml",
	}

	configLoader.Load(&cfg)

	// ===================== Fake S3 setup ====================== //
	log.Info("Starting the fake-s3")

	backend := s3mem.New()

	faker := gofakes3.New(backend)

	http.HandleFunc("/", faker.Server().ServeHTTP)

	// Create an instance of http.Server
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.General.Port),
		ReadTimeout:  cfg.General.ReadTimeout * time.Second,
		WriteTimeout: cfg.General.WriteTimeout * time.Second,
	}

	// Start the http server
	log.Infof("Starting the server on port: %v", cfg.General.Port)

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", server.Addr, err)
	}

	log.Infof("Listening on port: %v", cfg.General.Port)
}

// setLogLevel sets the log level based on the LOG_LEVEL env var
func setLogLevel() {
	switch getEnvVar("LOG_LEVEL", "info") {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
}

func getEnvVar(k, d string) string {
	v := os.Getenv(k)
	if v == "" {
		v = d
	}

	return v
}
