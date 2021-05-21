package parser

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"simple-golang-crawler/engine"
	"simple-golang-crawler/fetcher"
	"simple-golang-crawler/model"
	"simple-golang-crawler/tool"
	"strconv"

	"github.com/tidwall/gjson"
)

var _entropy = "rbMCKn@KuamXWlPMoJGsKcbiJKUfkPF_8dABscJntvqhRSETg"
var _paramsTemp = "appkey=%s&cid=%s&otype=json&qn=%s&quality=%s&type="
var _playApiTemp = "https://interface.bilibili.com/v2/playurl?%s&sign=%s"
var _quality = "80"

func GenGetAidChildrenParseFun(videoAid *model.VideoAid) engine.ParseFunc {
	return func(contents []byte, url string) engine.ParseResult {
		var retParseResult engine.ParseResult
		data := gjson.GetBytes(contents, "data").Array()
		appKey, sec := tool.GetAppKey(_entropy)

		var videoTotalPage int64
		for _, i := range data {
			cid := i.Get("cid").Int()
			page := i.Get("page").Int()
			part := i.Get("part").String()
			videoCid := model.NewVideoCidInfo(cid, videoAid, page, part)
			videoTotalPage += 1
			cidStr := strconv.FormatInt(videoCid.Cid, 10)

			params := fmt.Sprintf(_paramsTemp, appKey, cidStr, _quality, _quality)
			chksum := fmt.Sprintf("%x", md5.Sum([]byte(params+sec)))
			urlApi := fmt.Sprintf(_playApiTemp, params, chksum)
			req := engine.NewRequest(urlApi, GenVideoDownloadParseFun(videoCid), fetcher.DefaultFetcher)
			retParseResult.Requests = append(retParseResult.Requests, req)
		}

		videoAid.SetPage(videoTotalPage)
		item := engine.NewItem(videoAid)
		retParseResult.Items = append(retParseResult.Items, item)

		return retParseResult
	}
}

func GetRequestByAid(aid int64) *engine.Request {
	reqUrl := fmt.Sprintf(_getCidUrlTemp, aid)
	title := getTitle(aid)
	videoAid := model.NewVideoAidInfo(aid, title)
	reqParseFunction := GenGetAidChildrenParseFun(videoAid)
	req := engine.NewRequest(reqUrl, reqParseFunction, fetcher.DefaultFetcher)
	return req
}

func getTitle(aid int64) string {
	resp, err := http.Get(fmt.Sprintf("https://api.bilibili.com/x/web-interface/view?aid=%d", aid))
	if err != nil {
		log.Fatalf("http.Get() function error：%v\n", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("io.RedAll() function error：%v\n", err)
	}
	data := gjson.GetBytes(body, "data")
	title := data.Get("title").String()
	return title
}

func GetRequestByBvid(bvid string) *engine.Request {
	resp, err := http.Get(fmt.Sprintf("http://api.bilibili.com/x/web-interface/archive/stat?bvid=%s", bvid))
	if err != nil {
		log.Fatalf("http.Get() function error：%v\n", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("io.RedAll() function error：%v\n", err)
	}
	data := gjson.GetBytes(body, "data")
	aid := data.Get("aid").Int()
	return GetRequestByAid(aid)
}
