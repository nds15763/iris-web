package DB

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/kataras/iris"
	"github.com/olivere/elastic"
)

type T interface{}
type Tweet struct {
	User    string
	Message string
}

var client *elastic.Client
var err error

//https://www.oschina.net/p/elastic-go  文档

func EsInit() error {
	// Create a client
	client, err = elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	if err != nil {
		// Handle error
		return err
	}
	return nil
}

func EsCreateIndex() error {
	// Create an index
	_, err = client.CreateIndex("goods").Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	return nil
}

func EsCreateIndexDoc() error {
	// Add a document to the index
	tweet := Tweet{User: "olivere", Message: "Take Five"}
	_, err = client.Index().
		Index("goods").
		Type("goods").
		Id("1").
		BodyJson(tweet).
		Refresh("true").
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	return nil
}

func EsQuery(ctx iris.Context) {
	// Search with a term query
	termQuery := elastic.NewTermQuery("user", "olivere")
	searchResult, err := client.Search().
		Index("goods").          // search in index "twitter"
		Query(termQuery).        // specify the query
		Sort("user", true).      // sort by "user" field, ascending
		From(0).Size(10).        // take documents 0-9
		Pretty(true).            // pretty print request and response JSON
		Do(context.Background()) // execute
	if err != nil {
		// Handle error
		panic(err)
	}
	// searchResult is of type SearchResult and returns hits, suggestions,
	// and all kinds of other information from Elasticsearch.
	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

	// Each is a convenience function that iterates over hits in a search result.
	// It makes sure you don't need to check for nil values in the response.
	// However, it ignores errors in serialization. If you want full control
	// over iterating the hits, see below.

	// var t T
	var tweet Tweet
	for _, item := range searchResult.Each(reflect.TypeOf(tweet)) {
		if t, ok := item.(Tweet); ok {
			fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
		}
	}
	// TotalHits is another convenience function that works even when something goes wrong.
	fmt.Printf("Found a total of %d tweets\n", searchResult.TotalHits())

	// Here's how you iterate through results with full control over each step.
	if searchResult.Hits.TotalHits > 0 {
		fmt.Printf("Found a total of %d tweets\n", searchResult.Hits.TotalHits)

		// Iterate through results
		for _, hit := range searchResult.Hits.Hits {
			// hit.Index contains the name of the index

			// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).
			var t T
			err := json.Unmarshal(*hit.Source, &t)
			if err != nil {
				// Deserialization failed
			}

			// Work with tweet
			// fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
		}
	} else {
		// No hits
		fmt.Print("Found no tweets\n")
	}
}

func EsDeleteIndex() {

}
