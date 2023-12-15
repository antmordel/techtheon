package main

import (
	"fmt"
	"os"

	"github.com/antmordel/techtheon/foundation/logger"
	"github.com/mmcdole/gofeed"
	"go.uber.org/zap"
)

type Feed struct {
	Name   string
	Author string
	Link   string
}

type Article struct {
	Title   string
	Link    string
	PubDate string
	Content string
	Author  string
	Blog    string
}

var feeds = []Feed{
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

	log.Infow("startup", "status", "started")

	// Parse RSS
	fp := gofeed.NewParser()
	for _, feed := range feeds {
		if err := readFeed(log, fp, feed); err != nil {
			return err
		}
	}

	return nil
}

func readFeed(log *zap.SugaredLogger, fp *gofeed.Parser, feed Feed) error {

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

func parseArticle(item *gofeed.Item, feed Feed) (*Article, error) {

	article := Article{
		Title:   item.Title,
		Link:    item.Link,
		PubDate: item.Published,
		Content: item.Content,
		Author:  feed.Author,
	}

	return &article, nil
}
