package actions

import (
	"baixing_spider/structs/models"
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/panwenbin/ghttpclient"
	"regexp"
	"strings"
)

// 解析界面内容
func Parse(html string) (models.HtmlExtractorResult, error) {
	//res, err := ParseBaixingPc(html)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//if res.Title == "" && res.CompanyName == "" && res.ServiceContent == "" {
	//	res, err = ParseBaixingMobile(html)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//}

	res, err := ShenghuoServiceParsePc(html)
	if err != nil {
		fmt.Println(err)
	}

	if res.Title == "" && res.CompanyName == "" && res.ServiceContent == "" {
		res, err = ShenghuoServiceParseMobile(html)
		if err != nil {
			fmt.Println(err)
		}
	}

	return res, nil
}

// pc端解析
func ParseBaixingPc(html string) (models.HtmlExtractorResult, error) {
	res := models.HtmlExtractorResult{}
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return res, err
	}

	// 解析发布时间
	regexpStrs := []string{`title="首次发布于:(.*?)"`, `title='首次发布于：(.*?)'`}
	for _, reg := range regexpStrs {
		re := regexp.MustCompile(reg)
		subMatch := re.FindStringSubmatch(html)
		if len(subMatch) == 2 {
			res.ReleaseTime = subMatch[1]
			break
		}
	}

	dom.Find(".viewad-info .viewad-actions span[title]")

	// 解析公司名，服务内容，服务范围，联系人
	parseStrs := []string{".viewad-meta-item", ".viewad-meta2-item"}
	for _, parseStr := range parseStrs {
		if len(dom.Find(parseStr).Nodes) == 0 {
			continue
		}

		dom.Find(parseStr).Each(func(_ int, metaItem *goquery.Selection) {
			metaLabel := metaItem.Find("label").Text()
			if strings.Contains(metaLabel, "公司名称") {
				res.CompanyName = metaItem.Find(".content").Text()
			} else if strings.Contains(metaLabel, "服务内容") {
				res.ServiceContent = metaItem.Find(".content").Text()
			} else if strings.Contains(metaLabel, "服务范围") {
				res.ServiceRange = metaItem.Find(".content").Text()
			} else if strings.Contains(metaLabel, "联系人") {
				res.ContactPeople = metaItem.Find(".content").Text()
			}
		})
	}

	// 解析标题，联系电话，发布人
	parseStrs = []string{"div.viewad-main-info", ".viewad-content"}
	for _, parseStr := range parseStrs {
		if len(dom.Find(parseStr).Nodes) == 0 {
			continue
		}

		dom.Find(parseStr).Each(func(i int, selection *goquery.Selection) {
			res.Title = selection.Find("h1").Text()
			res.ContactPhone = selection.Find("section.viewad-contact a.show-contact[data-phone]").AttrOr("data-phone", "")
			res.ContactWeixin = strings.Replace(strings.Replace(selection.Find(".weixin-contact-promo .detail").Text(), "微信号:", "", -1), " ", "", -1)

			if res.ContactPhone == "" {
				res.ContactPhone = dom.Find("#mobileNumber strong").Text()
			}

			res.Poster = selection.Find(" section.poster-info div.poster-detail h3").Text()
			if res.Poster == "" {
				res.Poster = dom.Find(".poster-name").Text()
			}
		})
	}

	// 如果没有联系方式，则访问原网页解析
	if res.ContactPhone == "" {
		href := dom.Find("#bd_snap_note a[href]").AttrOr("href", "")
		res.ContactPhone = VisitSourceUrlAndParse(href)
	}

	return res, nil
}

// mobile端解析
func ParseBaixingMobile(html string) (models.HtmlExtractorResult, error) {
	res := models.HtmlExtractorResult{}
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return res, err
	}

	// 解析标题，发布时间，联系电话，发布者
	res.Title = dom.Find("section.title h1").Text()
	res.ReleaseTime = strings.Replace(dom.Find("time.datetime").Text(), "首次发布于:", "", -1)
	res.ContactPhone = dom.Find(".contact-container a.main-btn").AttrOr("data-value", "")
	if res.ContactPhone == "" {
		res.ContactPhone = dom.Find(".contact-inner a.contact-sms").AttrOr("data-value", "")
		// 如果没有联系方式，则访问原网页解析
		if res.ContactPhone == "" {
			href := dom.Find("#bd_snap_note a[href]").AttrOr("href", "")
			res.ContactPhone = VisitSourceUrlAndParse(href)
		}
	}

	res.Poster = dom.Find(".user-info strong").Text()
	if res.Poster == "" {
		res.Poster = dom.Find(".user-info a.block").Text()
	}

	// 解析公司名，服务内容，服务范围，联系人及发布时间
	dom.Find("ul.bx-meta-list li.meta-item").Each(func(_ int, metaItem *goquery.Selection) {
		metaLabel := metaItem.Find("label").Text()
		if strings.Contains(metaLabel, "公司名称") {
			res.CompanyName = metaItem.Find(".meta-action-content").Text()
			if res.CompanyName == "" {
				res.CompanyName = metaItem.Find("div").Text()
			}
		} else if strings.Contains(metaLabel, "服务内容") {
			res.ServiceContent = metaItem.Find(".meta-action-content").Text()
			if res.ServiceContent == "" {
				res.ServiceContent = metaItem.Find(".tag").Text()
			}
		} else if strings.Contains(metaLabel, "服务范围") {
			res.ServiceRange = metaItem.Find(".meta-action-content").Text()
			if res.ServiceRange == "" {
				res.ServiceRange = metaItem.Find(".tag").Text()
			}
		} else if strings.Contains(metaLabel, "联系人") {
			res.ContactPeople = metaItem.Find(".meta-action-content").Text()
			if res.ContactPeople == "" {
				res.ContactPeople = metaItem.Find("div").Text()
			}
		} else if strings.Contains(metaLabel, "更新时间") {
			if res.ReleaseTime == "" {
				res.ReleaseTime = metaItem.Find(".meta-action-content").Text()
			}
		}
	})

	return res, nil
}

// 访问原网页并解析联系方式
func VisitSourceUrlAndParse(url string) string {
	body, _ := ghttpclient.Get(url, nil).ReadBodyClose()
	dom, _ := goquery.NewDocumentFromReader(bytes.NewReader(body))
	return dom.Find(".viewad-contact a.show-contact[data-phone]").AttrOr("data-phone", "")
}
