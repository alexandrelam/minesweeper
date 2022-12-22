// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	flag.Parse()
	hub := newHub()
	go hub.run()

	r := mux.NewRouter()
	r.HandleFunc("/ws/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := mux.Vars(r)["name"]
		fmt.Println(name)
		serveWs(hub, w, r, name)
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
