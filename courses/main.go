package main

import (
	"fmt"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"
)

// var PORT = ":1443"

type handler struct {
}

func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	res, _ := w.Write([]byte("Hello world!\n"))
	fmt.Println(res)
}

func Default(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "This is an example HTTPS server!\n")
}

func DefaultTest(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "This is SPARTA!\n")
}

func main()  {
	hostPort := "127.0.0.1:1443"
	router := mux.NewRouter()
	router.HandleFunc("/api/v0/", DefaultTest).Methods(http.MethodGet)
	router.HandleFunc("/", Default).Methods(http.MethodGet)
	// http.HandleFunc("/", Default)

	caCert, err := ioutil.ReadFile("courses/client.crt")
	if err != nil {
		fmt.Println(err)
		return
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	cfg := &tls.Config{
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs: caCertPool,
	}

	srv := &http.Server{
		Addr: hostPort,
		// Handler: &handler{},
		Handler: router,
		TLSConfig: cfg,
	}
	
	fmt.Println("Listening to port number", hostPort)
	fmt.Println(srv.ListenAndServeTLS(
		"courses/server.crt",
		"courses/server.key",
	))

	fmt.Println("Listening to port number", "127.0.0.1:1443")
	// err := http.ListenAndServeTLS(
	// 	"127.0.0.1:1443",
	// 	"courses/server.crt",
	// 	"courses/server.key",
	// 	router,
	// )
	// if err != nil {
	// 	fmt.Println("ListenAndServeTLS: ", err)
	// 	return
	// }
}
