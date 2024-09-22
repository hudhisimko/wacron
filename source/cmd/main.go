package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/posflag"
	"github.com/knadh/koanf/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	flag "github.com/spf13/pflag"

	"simko/wacron/db"
	"simko/wacron/messenger"
)

var (
	k                  = koanf.New(".")
	defaultConfigPaths = []string{"./config.yml", "/config/config.yml"}
)

func init() {
	log.Logger = log.Output(zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
		w.Out = os.Stderr
		w.TimeFormat = time.RFC3339
	}))
	log.Logger = log.Level(zerolog.InfoLevel)
}

func setupConfig() {
	f := flag.NewFlagSet("config", flag.ContinueOnError)
	f.Usage = func() {
		fmt.Println(f.FlagUsages())
		os.Exit(0)
	}
	// Path to one or more config files to load into koanf along with some config params.
	f.StringSliceP("config", "c", defaultConfigPaths, "path to one or more .yaml config files")
	f.Parse(os.Args[1:])

	// Load the config files provided in the command line.
	cFiles, _ := f.GetStringSlice("config")
	for _, c := range cFiles {
		provider := file.Provider(c)
		if err := k.Load(provider, yaml.Parser()); err != nil {
			log.Warn().Err(err).Msg("error loading file")
		} else {
			log.Info().Str("file", c).Msg("loaded config")
		}
	}

	// flags override
	if err := k.Load(posflag.Provider(f, ".", k), nil); err != nil {
		log.Warn().Err(err).Msg("error loading config from flag")
	}
}

func initDevelopment() {
	log.Logger = log.Level(zerolog.TraceLevel)
	if k.Bool("print_query") {
		db.EnableQueryLog()
	}
}

func main() {
	setupConfig()
	if k.Bool("development") {
		initDevelopment()
	}

	// configure needed settings
	if d := k.Duration("interval"); d > 0 {
		messenger.PendingMessagesIterationInterval = d
	}

	if n := k.Int("iteration_limit"); n > 0 {
		messenger.PendingMessagesIterationLimit = n
	}

	// Connect db for data
	dataDB, err := db.Open(k.String("db_data_dsn"))
	if err != nil {
		log.Fatal().Err(err).Msg("unable connect db (data)")
	}

	// Connect db for gateway
	gatewayDB, err := db.Open(k.String("db_gateway_dsn"))
	if err != nil {
		log.Fatal().Err(err).Msg("unable connect db (gateway)")
	}

	// Parse Schemas, this ensures that latter
	// we will use database as declared in dsn, not the hardcoded one from jetgen.
	dataSchema, err := db.DSNSchema(k.String("db_data_dsn"))
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	gatewaysSchema, err := db.DSNSchema(k.String("db_gateway_dsn"))
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	m := messenger.NewMessenger(dataDB, gatewayDB, dataSchema, gatewaysSchema, k.Int("worker"))

	// set operation mode
	var om messenger.OperationMode
	if err := k.Unmarshal("operation_mode", &om); err != nil {
		log.Fatal().Err(err).Msg("unable parse operation_mode")
	}
	m.OperationMode = om
	log.Info().Any("operation_mode", om).Send()

	ctx, cancel := context.WithCancel(context.Background())
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	go func() { <-interrupt; cancel() }()

	m.Start()

	go func() {
		log.Info().Msg("Metric server is listening on :3541")
		messenger.ListenMetric()
	}()

	<-ctx.Done()

	log.Info().Msg("Stopping...")
	m.Stop()

	log.Info().Msg("Exited")
}
