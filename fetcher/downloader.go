package fetcher

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"simple-golang-crawler/model"
	"simple-golang-crawler/tool"
)

var _startUrlTem = "https://api.bilibili.com/x/web-interface/view?aid=%d"

func GenVideoFetcher(video *model.Video) FetchFun {
	referer := fmt.Sprintf(_startUrlTem, video.ParCid.ParAid.Aid)
	for i := int64(1); i <= video.ParCid.Page; i++ {
		referer += fmt.Sprintf("/?p=%d", i)
	}

	return func(url string) (bytes []byte, err error) {
		<-_rateLimiter.C
		client := http.Client{CheckRedirect: genCheckRedirectfun(referer)}

		request, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatalln(url, err)
			return nil, err
		}
		request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.13; rv:56.0) Gecko/20100101 Firefox/56.0")
		request.Header.Set("Accept", "*/*")
		request.Header.Set("Accept-Language", "en-US,en;q=0.5")
		request.Header.Set("Accept-Encoding", "gzip, deflate, br")
		request.Header.Set("Range", "bytes=0-")
		request.Header.Set("Referer", referer)
		request.Header.Set("Origin", "https://www.bilibili.com")
		request.Header.Set("Connection", "keep-alive")

		resp, err := client.Do(request)
		if err != nil {
			log.Fatalf("Fail to download the video %s,err is %s", video.ParCid.Part, err)
			return nil, err
		}

		if resp.StatusCode != http.StatusPartialContent {
			log.Fatalf("Fail to download the video %s,status code is %d", video.ParCid.Part, resp.StatusCode)
			return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
		}
		defer resp.Body.Close()

		aidPath := tool.GetAidFileDownloadDir(video.ParCid.ParAid.Aid, video.ParCid.ParAid.Title)
		filename := fmt.Sprintf("P%d_%s.flv", video.ParCid.Page, video.ParCid.Part)
		file, err := os.Create(filepath.Join(aidPath, filename))
		if err != nil {
			log.Fatalln(err)
			os.Exit(1)
		}
		defer file.Close()

		log.Println(video.ParCid.ParAid.Title + ":" + filename + " is downloading.")
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			log.Printf("Failed to download video %s", video.ParCid.Part)
			return nil, err
		}
		log.Println(video.ParCid.ParAid.Title + ":" + filename + " has finished.")

		return nil, nil
	}
}

func genCheckRedirectfun(referer string) func(req *http.Request, via []*http.Request) error {
	return func(req *http.Request, via []*http.Request) error {
		req.Header.Set("Referer", referer)
		return nil
	}
}
