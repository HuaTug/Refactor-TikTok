package elasticsearch

import (
	"io/ioutil"
	"log"

	"github.com/olivere/elastic/v7"
)

var (
	elasticClient *elastic.Client
)

func Load() {
	var err error
	elasticClient, err = elastic.NewClient(
		elastic.SetURL(ElasticAddr),
		elastic.SetSniff(false),
		elastic.SetInfoLog(log.New(ioutil.Discard, "", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)),  //debug as os.stdout
		elastic.SetErrorLog(log.New(ioutil.Discard, "", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)), //debug as os.stderr
		elastic.SetTraceLog(log.New(ioutil.Discard, "", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)),
	)

	newVideoIndex()

	if err != nil {
		panic(err)
	}

}
