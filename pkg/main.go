// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/alexandrelam/minesweeper/pkg/websocket"
	"github.com/gorilla/mux"
	jsoniter "github.com/json-iterator/go"
)

var addr = flag.String("addr", ":3001", "http service address")
var json = jsoniter.ConfigCompatibleWithStandardLibrary

func main() {
	flag.Parse()
	hub := websocket.NewHub()
	go hub.Run()

	r := mux.NewRouter()
	r.HandleFunc("/ws/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := mux.Vars(r)["name"]
		websocket.ServeWs(hub, w, r, name)
	})

	srv := &http.Server{
		Handler: r,
		Addr:    *addr,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
