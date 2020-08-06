package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"nrei.name/blogapi/post"
)

const (
	defaultPort = "8080"
)

func main() {
	// var (
	// 	addr = utils.EnvGet("PORT", defaultPort)
	// 	ctx  = context.Background()
	// )

	var logger log.Logger

	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "", log.DefaultTimestampUTC)

	errs := make(chan error, 2)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	// @nrei debug post.
	var post = post.New(post.ID("123"), "my super long text of 80 chars will be here and the whole phrase should be sliced", "my awesome content")
	logger.Log(fmt.Printf("post data\n--\nid: %s\ncreateDate: %s\nmodifyDate: %s\ntitle: %s\ncontent: %s\n", post.ID, post.CreateDate, post.ModifyDate, post.Title, post.Content))
	// @nrei debug post end.

	logger.Log("terminated", <-errs)
}
