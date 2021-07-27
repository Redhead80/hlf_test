package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type LedgerRestServer struct {
	client *LedgerClient
}

func StartRestServer(client *LedgerClient, port int) {
	restServer := LedgerRestServer{
		client: client,
	}

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/queryInfo", restServer.queryInfoHandler)
	router.HandleFunc("/queryBlockByNumber/{block}", restServer.queryBlockByNumber)
	router.HandleFunc("/queryBlockByHash/{hash}", restServer.queryBlockByHash)
	router.HandleFunc("/queryConfigBlock", restServer.queryConfigBlock)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

func sendResponse(w http.ResponseWriter, response interface{}) {
	bytes, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(bytes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (server *LedgerRestServer) queryInfoHandler(w http.ResponseWriter, _ *http.Request) {
	info, err := server.client.QueryInfo()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendResponse(w, info.BCI)
}

func (server *LedgerRestServer) queryBlockByNumber(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blockNumber, err := strconv.ParseInt(vars["block"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	block, err := server.client.QueryBlockByNumber(uint64(blockNumber))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendResponse(w, block)
}

func (server *LedgerRestServer) queryBlockByHash(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blockHash, err := base64.StdEncoding.DecodeString(vars["hash"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	block, err := server.client.QueryBlockByHash(blockHash)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendResponse(w, block)
}

func (server *LedgerRestServer) queryConfigBlock(w http.ResponseWriter, _ *http.Request) {
	block, err := server.client.QueryConfigBlock()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendResponse(w, block)
}
