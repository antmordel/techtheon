package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/antmordel/techtheon/foundation/logger"
	"github.com/antmordel/techtheon/pkg/rss"
	"github.com/mmcdole/gofeed"
	"go.uber.org/automaxprocs/maxprocs"
	"go.uber.org/zap"
)

var feeds = []rss.Feed{
	{
		Name:   "Irrational Exuberance",
		Author: "Will Larson",
		Link:   "https://lethain.com/feeds.xml",
	},
}

func main() {

	// Construct the application logger.
	log, err := logger.New("rss-feed")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer log.Sync()

	// Perform the startup and shutdown sequence.
	if err := run(log); err != nil {
		log.Errorw("startup", "ERROR", err)
		os.Exit(1)
	}
}

func run(log *zap.SugaredLogger) error {

	// =========================================================================
	// GOMAXPROCS

	// Set the correct number of threads for the service
	// based on what is available either by the machine or quotas.
	if _, err := maxprocs.Set(); err != nil {
		return fmt.Errorf("maxprocs: %w", err)
	}
	log.Infow("startup", "GOMAXPROCS", runtime.GOMAXPROCS(0))

	// =========================================================================
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	sig := <-shutdown
	log.Infow("shutdown", "status", "shutdown started", "signal", sig)
	defer log.Infow("shutdown", "status", "shutdown complete", "signal", sig)

	return nil
}

func readFeed(log *zap.SugaredLogger, fp *gofeed.Parser, feed rss.Feed) error {

	log.Infow("parsing feed", "feed", feed.Name)

	f, err := fp.ParseURL(feed.Link)
	if err != nil {
		return err
	}

	for _, item := range f.Items {
		_, err := parseArticle(item, feed)
		if err != nil {
			return err
		}
	}

	return nil
}

func parseArticle(item *gofeed.Item, feed rss.Feed) (*rss.Article, error) {

	article := rss.Article{
		Title:   item.Title,
		Link:    item.Link,
		PubDate: item.Published,
		Content: item.Content,
		Author:  feed.Author,
	}

	return &article, nil
}
