package main

import (
	"flag"
	"github.com/daaku/go.flagenv"
	log "github.com/marcw/gogol"
	"github.com/marcw/libpoller"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
)

var (
	port = flag.String("port", "", "Port on which the http server will listen.")
)

func init() {
	flagenv.UseUpperCaseFlagNames = true
	flagenv.Parse()
}

func main() {
	store := poller.NewInMemoryStore()
	scheduler := poller.NewSimpleScheduler()
	config := poller.NewConfig(store, scheduler)

	p := poller.NewDirectPoller()
	librato, err := poller.NewLibratoBackend(os.Getenv("LIBRATO_USER"), os.Getenv("LIBRATO_TOKEN"), "poller", "poller.")
	if err != nil {
		log.Fatalln(err)
	}

	go httpInput(config)
	go p.Run(scheduler, librato, poller.NewHttpProbe("Poller (https://github.com/marcw/poller)", 10*time.Second), nil)
	go scheduler.Start()

	select {}
}

func httpInput(config *poller.Config) {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {})
	http.Handle("/checks", poller.NewConfigHttpHandler(config))
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatalln(err)
	}
}
