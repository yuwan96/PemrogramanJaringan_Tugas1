package main

import (
	"fmt"
    "html"
    "github.com/spf13/viper"
    "net/http"
)

func main() {
	// Set lokasi file config
	viper.SetConfigFile("./config/env.json")
	
	// Tampilkan error jika file config tidak ditemukan
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file, %s", err)
	}
	
	// Set route dan konten ketika web diakses
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %q", html.EscapeString(r.URL.Path))
	})

	// Jalankan server di port yang sudah di-set config
	http.ListenAndServe(":" + viper.GetString("server.port"), nil)