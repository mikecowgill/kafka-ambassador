package main

import (
	"flag"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/anchorfree/kafka-ambassador/pkg/config"
	"github.com/anchorfree/kafka-ambassador/pkg/logger"
	"github.com/anchorfree/kafka-ambassador/pkg/servers"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	defaults map[string]interface{} = map[string]interface{}{
		"global.log.level":            "info",
		"global.log.encoding":         "json",
		"global.log.outputPaths":      ("stdout"),
		"global.log.errorOutputPaths": ("stderr"),
		"global.log.encoderConfig":    logger.NewEncoderConfig(),
		"server.http.listen":          ":19092",
		// Please take a look at:
		// https://github.com/edenhill/librdkafka/blob/master/CONFIGURATION.md
		// for more configuration parameters
		"kafka.compression.codec":                     "gzip",
		"kafka.batch.num.messages":                    100000,
		"kafka.socket.timeout.ms":                     10000, // mark connection as stalled
		"kafka.message.timeout.ms":                    60000, // try to deliver message with retries
		"kafka.max.in.flight.requests.per.connection": 20,
		"server.grpc.max.request.size":                4 * 1024 * 1024,
		"server.grpc.monitoring.histogram.enable":     true,
		"server.grpc.monitoring.enable":               true,
	}
)

func main() {
	var err error
	var configPathName string
	flag.StringVar(&configPathName, "config", "", "Configuration file to load")
	flag.Parse()

	// We need to shut down gracefully when the user hits Ctrl-C.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGUSR1, syscall.SIGTERM)

	s := new(servers.T)
	s.Wg = new(sync.WaitGroup)
	c := &config.T{
		Filename:  configPathName,
		EnvPrefix: "ka",
	}
	s.Config, err = c.ReadConfig(defaults)
	if err != nil {
		return
	}
	// logs
	cfg := logger.NewLogConfig(s.Config.Sub("global.log"))
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	s.Logger = logger.Sugar()
	// metrics
	s.Prometheus = prometheus.NewRegistry()

	// servers
	kafkaParams, err := config.KafkaParams(s.Config)
	if err != nil {
		return
	}
	s.Producer.Logger = s.Logger
	s.Producer.Config = config.ProducerConfig(s.Config)
	s.Producer.Init(&kafkaParams, s.Prometheus)
	s.Start()
	signal := <-sig
	switch signal {
	case syscall.SIGTERM, syscall.SIGINT:
		s.Stop()
		for {
			if !s.Producer.QueueIsEmpty() {
				s.Logger.Info("We still have messages in queue, waiting")
				time.Sleep(5 * time.Second)
			} else {
				s.Logger.Info("Queue is empty, shut down properly")
				s.Producer.Producer.Close()
				break
			}
		}
	}
}