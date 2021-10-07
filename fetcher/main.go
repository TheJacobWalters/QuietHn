package fetcher

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Story struct {
	Title  string
	Url    string
	Author string
}

func GetContent(numStories int) []Story {
	results := make([]Story, 0)
	for i := 1; i < numStories; i++ {
		url := fmt.Sprintf("http://hn.algolia.com/api/v1/items/%d", i)
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		var story Story
		err = json.Unmarshal(body, &story)
		if err != nil {
			panic(err)
		}
		results = append(results, story)
	}
	return results
}
