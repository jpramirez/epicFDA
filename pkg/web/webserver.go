package web

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/handlers"
	models "github.com/jpramirez/epicFDA/pkg/models"
	webapp "github.com/jpramirez/epicFDA/pkg/web/app"
)

//WebAgent is the main struct for this agent.
type WebOne struct {
	webConfig models.Config
}

//StartServer Starts the server using the variable sip and port, creates anew instance.
func (W *WebOne) StartServer() {

	go http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://"+r.Host+r.URL.String(), http.StatusMovedPermanently)
	}))

	handler := W.New()
	f, err := os.OpenFile(W.webConfig.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	srv := &http.Server{
		Handler:           handlers.LoggingHandler(f, handler),
		Addr:              W.webConfig.WebAddress + ":" + W.webConfig.WebPort,
		ReadHeaderTimeout: 20 * time.Second,
		ReadTimeout:       5 * time.Minute,
		WriteTimeout:      2 * time.Minute,
		IdleTimeout:       120 * time.Second,
		TLSConfig: &tls.Config{
			// Causes servers to use Go's default ciphersuite preferences,
			// which are tuned to avoid attacks. Does nothing on clients.
			PreferServerCipherSuites: true,
			// Only use curves which have assembly implementations
			CurvePreferences: []tls.CurveID{
				tls.CurveP256,
				tls.X25519, // Go 1.8 only
			},
			MinVersion: tls.VersionTLS12,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,

				// Best disabled, as they don't provide Forward Secrecy,
				// but might be necessary for some clients
				// tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
				// tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
			},
		},
	}

	err = srv.ListenAndServeTLS(W.webConfig.CrtFile, W.webConfig.KeyFile)

	if err != nil {
		log.Println("Error Starting web server")
	}
}

//NewWebAgent creates new instance.
func NewWebAgent(config models.Config, BuildVersion string, BuidTime string) (WebOne, error) {
	var webone WebOne
	log.Println("Starting Go Epic FDA ")
	log.Println("Version : " + BuildVersion)
	log.Println("Build Time : " + BuidTime)

	webone.webConfig = config
	log.Println("Listening on ", webone.webConfig.WebAddress, webone.webConfig.WebPort)

	// Stop the grpc verbose logging
	//grpclog.SetLogger(noplog)
	return webone, nil
}

//New creates a new handler
func (W *WebOne) New() http.Handler {
	app, err := webapp.NewApp(W.webConfig)

	if err != nil {
		log.Fatalln("Error creating WebApp", err)
		return nil
	}
	// API Calls
	api := app.Mux.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/liveness", app.Liveness)
	api.HandleFunc("/downloadIndex", app.DownloadIndex)
	api.HandleFunc("/fetchDrugDataSet", app.DownloadDrugDataSet)
	api.HandleFunc("/fetchAnimalDataSet", app.DownloadAnimalDataSet)
	api.HandleFunc("/fetchFoodDataSet", app.DownloadFoodDataSet)
	api.HandleFunc("/loadFoodEnforcement/{dataset}", app.LoadFoodEnforcement)

	//For UI functions
	app.Mux.HandleFunc("/", app.HandleIndex).Methods("GET")
	//Serve Static Folder
	fs := http.FileServer(http.Dir("./assets/static"))
	app.Mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	ch := make(chan os.Signal, 2)

	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ch
		log.Println("Closing system")
		os.Exit(0)
	}()
	return &app
}
