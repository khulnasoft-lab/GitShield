package main

import (
	"os"
	"os/signal"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/khulnasoft-lab/gitshield/v8/cmd"
)

func main() {
	// send all logs to stdout
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// this block sets up a go routine to listen for an interrupt signal
	// which will immediately exit gitshield
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)
	go listenForInterrupt(stopChan)

	cmd.Execute()
}

func listenForInterrupt(stopScan chan os.Signal) {
	<-stopScan
	log.Fatal().Msg("Interrupt signal received. Exiting...")
}
