package serve

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"reflect"
	"time"

	"github.com/gorilla/mux"
)

var server *http.Server
var itemToServe map[string]string

func Start(item map[string]string) {
	itemToServe = item

	r := mux.NewRouter()
	r.HandleFunc("/getvoucher", serveVoucher).Methods(http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodOptions)
	r.HandleFunc("/redeemvoucher/{code}", redeemVoucher).Methods(http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodOptions)
	r.Use(mux.CORSMethodMiddleware(r))

	server = &http.Server{
		Addr:              ":8081",
		Handler:           r,
		TLSConfig:         &tls.Config{},
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      map[string]func(*http.Server, *tls.Conn, http.Handler){},
		ConnState: func(net.Conn, http.ConnState) {
		},
		ErrorLog: &log.Logger{},
	}

	defer Stop()
	server.ListenAndServe()
}

func Stop() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	server.Shutdown(ctx)
}

func serveVoucher(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.WriteHeader(http.StatusOK)

	fmt.Fprintf(res, itemToServe[mapRandomKeyGet(itemToServe).(string)])
}

func redeemVoucher(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(req)
	res.WriteHeader(http.StatusOK)

	for _, value := range itemToServe {
		if value == vars["code"] {
			value = "Voucher redeemed"
			fmt.Fprintf(res, value)
			return
		}
	}
}

func mapRandomKeyGet(mapI interface{}) interface{} {
	keys := reflect.ValueOf(mapI).MapKeys()

	return keys[rand.Intn(len(keys))].Interface()
}
