package main

import (
	log "github.com/cihub/seelog"
)

func main() {
	logger, err := log.LoggerFromConfigAsFile("./user/log/log.xml")

	if err != nil {
		log.Critical("err parsing config log file", err)
		return
	}
	log.ReplaceLogger(logger)

	defer log.Flush()
	log.Debug("Hello from seelog!")
}
