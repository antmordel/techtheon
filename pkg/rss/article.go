package rss

import (
	"bytes"
	"database/sql"
	"encoding/gob"

	"github.com/antmordel/techtheon/pkg/data/db"
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

var dbConnection *sql.DB

func init() {
	var err error
	dbConnection, err = db.Connect()
	if err != nil {
		panic(err)
	}
}

func produce(log *zap.SugaredLogger, article Article) error {

	hash := hash(article)
	log.Infow("article", article, "hash", hash)
	exists := checkIfArticleExists(hash)
	if !exists {
		log.Infow("produce to kafka", "article", article, "hash", hash)
	}

	return nil
}

func checkIfArticleExists(hash []byte) bool {
	row := dbConnection.QueryRow("SELECT * FROM article_hash WHERE article_hash = $1", hash)
	var result []byte
	err := row.Scan(&result)
	if err != nil {
		return false
	}
	return true
}

func hash(article Article) []byte {
	var b bytes.Buffer
	gob.NewEncoder(&b).Encode(article)
	return b.Bytes()
}
