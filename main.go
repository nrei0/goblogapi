package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"nrei.name/blogapi/post"
	"nrei.name/blogapi/utils"
)

const (
	defaultPort = "8080"
)

func main() {
	var (
		addr     = utils.EnvGet("PORT", defaultPort)
		httpAddr = flag.String("http.addr", ":"+addr, "HTTP listen address")
	)
	flag.Parse()

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	httpLogger := log.With(logger, "component", "http")

	psr := post.NewRepository()

	// For testing purposes.
	mockTestPostsData(psr)

	var ps post.Service
	ps = post.NewService(psr)

	mux := http.NewServeMux()

	mux.Handle("/posts/v1/", post.MakeHandler(ps, httpLogger))

	http.Handle("/", accessControl(mux))

	errs := make(chan error, 2)
	go func() {
		httpLogger.Log("transport", "http", "address", *httpAddr)
		errs <- http.ListenAndServe(*httpAddr, nil)
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("terminated", <-errs)
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

func mockTestPostsData(p post.Repository) {
	if err := p.Save(post.New(post.ID("123"), "my super long text of 80 chars will be here and the whole phrase should be sliced", "my awesome content")); err != nil {
		panic(err)
	}
}
