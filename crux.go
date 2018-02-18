package main

import (
	"log"
	"net/http"
	"github.com/blk-io/crux/enclave"
	"github.com/blk-io/crux/server"
	"github.com/blk-io/crux/storage"
)

func main() {
	enc := enclave.Enclave{Db : storage.Init("/Users/Conor/code/go/blk-io/tmp/crux.db")}

	// TODO: Read key from configuration & add command line tool to generate
	tm := server.TransactionManager{Key : enclave.NewKey(), Enclave : enc}

	http.HandleFunc("/upcheck", tm.Upcheck)

	// TODO: Restrict to IPC
	http.HandleFunc("/send", tm.Send)
	http.HandleFunc("/receive", tm.Receive)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

	// TODO: Add support for propogation methods
}
