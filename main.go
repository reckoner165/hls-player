package main

import (
  "encoding/json"
  "log"
  "net/http"
)

func main() {
  fs := http.FileServer(http.Dir("./public"))
  http.Handle("/", fs)

  configHandler:= http.HandlerFunc(handleRequest)
  http.Handle("/config", configHandler)

  log.Println("Listening on :4000...")
  err := http.ListenAndServe(":4000", nil)
  if err != nil {
    log.Fatal(err)
  }
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["src"] = "https://test-streams.mux.dev/x36xhzz/x36xhzz.m3u8"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}