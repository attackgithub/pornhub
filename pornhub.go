package pornhub

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

const apiURL = "http://www.pornhub.com/webmasters/"
const APITimeout = 3

func SearchVideos(search string) (PornhubSearchResult) {
	timeout := time.Duration(APITimeout * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, _ := client.Get(fmt.Sprintf(apiURL+"search?search=%s&thumbnail=all", url.QueryEscape(search)))
	b, _ := ioutil.ReadAll(resp.Body)
	var result PornhubSearchResult
	err := json.Unmarshal(b, &result)
	if err != nil {
   		log.Fatal(err)
	}
	return result
}

func GetVideoByID(ID string) (PornhubSingleVideo) {
	resp, _ := http.Get(fmt.Sprintf(apiURL+"video_by_id?id=%s", ID))
	b, _ := ioutil.ReadAll(resp.Body)
	var result PornhubSingleVideo
	err := json.Unmarshal(b, &result)
	if err != nil {
		log.Fatal(err)
	}
	return result

}

func GetVideoEmbedCode(ID string) (PornhubEmbedCode) {
	resp, _ := http.Get(fmt.Sprintf(apiURL+"video_embed_code?id=%s", ID))
	b, _ := ioutil.ReadAll(resp.Body)
	var result PornhubEmbedCode
	err := json.Unmarshal(b, &result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}