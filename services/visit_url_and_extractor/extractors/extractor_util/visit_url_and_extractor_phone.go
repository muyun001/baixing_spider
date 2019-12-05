package extractor_util

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/panwenbin/ghttpclient"
)

// 访问原网页并解析联系方式
func VisitSourceUrlAndParse(url string) string {
	body, _ := ghttpclient.Get(url, nil).ReadBodyClose()
	dom, _ := goquery.NewDocumentFromReader(bytes.NewReader(body))
	return dom.Find(".viewad-contact a.show-contact[data-phone]").AttrOr("data-phone", "")
}
