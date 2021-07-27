package main

import (
	"flag"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"log"
)


func main() {

	configPath := flag.String("config", "config.yaml", "path to the config file")
	flag.Parse()

	configuration, err := LoadConfigurationFromFile(*configPath)
	if err != nil {
		log.Fatalf("failed loading client configuration from the file %s: %v", *configPath, err)
	}

	clientConfiguration := configuration.Client

	sdk, err := fabsdk.New(config.FromFile(clientConfiguration.ConnectionProfile))
	if err != nil {
		log.Fatalf("failed initializing Fabric SDK: %v", sdk)
	}

	defer sdk.Close()

	ledgerClient, err := NewLedgerClient(sdk, clientConfiguration)
	if err != nil {
		log.Fatalf("failed initializing the ledger client: %v", err)
	}

	StartRestServer(ledgerClient, configuration.Application.Port)
}
