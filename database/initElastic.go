package database

import (
	"github.com/elastic/go-elasticsearch/v8"

	"log"
	"fmt"
	"sync"
)

var (
	lock &sync.Mutex{}
	singletonElastic elasticsearch.Client
	username := ""
	password := ""
)

func InitElastic() *elasticsearch.Client {
	if singletonElastic == nil {
		lock.Lock()
		defer lock.Unlock()

		if singletonElastic == nil {
			cfg := elasticsearch.Config{
				Addresses: [""],
				Username: username,
				Password: password,
			}

			es, err := elasticsearch.NewClient(cfg)
			if err != nil {
				log.Fatalln("Failed to get connection to elsaticsearch")
			}

			fmt.Println("Connected to Elasticsearch")

			return es
		}

		fmt.Println("Connected to Elsaticsearch")
	}

	fmt.Println("Connected to Elsaticsearch")
	
	return singletonElastic
}