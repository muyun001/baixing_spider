package extractors

import (
	"baixing_spider/services/visit_url_and_extractor/extractors/extractor_util"
	"baixing_spider/structs/models"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func BaixingShenghuoExtractorPc(html string) (models.HtmlExtractorResult, error) {
	res := models.HtmlExtractorResult{}
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return res, errors.New(fmt.Sprintf("BaixingShenghuoExtractorPc goquery error: %s", err.Error()))
	}

	// 解析发布时间
	releaseTime := dom.Find("div.viewad-info div.viewad-actions span[title]").AttrOr("title", "")
	if releaseTime != "" {
		res.ReleaseTime = strings.Replace(releaseTime, "首次发布于:", "", -1)
	}
	if res.ReleaseTime == "" {
		releaseTime := dom.Find(".viewad-actions span[title]").AttrOr("title", "")
		if releaseTime != "" {
			res.ReleaseTime = strings.Replace(releaseTime, "首次发布于：", "", -1)
		} else {
			res.ReleaseTime = dom.Find(".viewad-actions span[title]").Text()
		}
	}

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
			} else if strings.Contains(metaLabel, "维修类型") {
				res.ServiceContent = metaItem.Find(".content").Text()
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
			res.Title = strings.Replace(selection.Find("h1").Text(), "\n", "", -1)
			res.ContactPhone = selection.Find("section.viewad-contact a.show-contact[data-phone]").AttrOr("data-phone", "")
			res.ContactWeixin = strings.Replace(strings.Replace(selection.Find(".weixin-contact-promo .detail").Text(), "微信号:", "", -1), " ", "", -1)

			if res.ContactPhone == "" {
				res.ContactPhone = dom.Find("#mobileNumber strong").Text()
				if res.ContactPhone == "" { // 如果没有联系方式，则访问原网页解析
					href := dom.Find("#bd_snap_note a[href]").AttrOr("href", "")
					res.ContactPhone = extractor_util.VisitSourceUrlAndParse(href)
				}
			}

			res.Poster = selection.Find("section.poster-info div.poster-detail h3").Text()
			if res.Poster == "" {
				res.Poster = dom.Find(".poster-name").Text()
			}
		})
	}

	return res, nil
}

func BaixingShenghuoExtractorMobile(html string) (models.HtmlExtractorResult, error) {
	res := models.HtmlExtractorResult{}
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return res, err
	}

	// 解析标题，发布时间，联系电话，发布者
	res.Title = strings.Replace(dom.Find("section.title h1").Text(), "\n", "", -1)
	res.ReleaseTime = strings.Replace(dom.Find("time.datetime").Text(), "首次发布于:", "", -1)
	res.ContactPhone = dom.Find("span.contact-main-txt").Text()
	if res.ContactPhone == "" { // 如果没有联系方式，则访问原网页解析
		href := dom.Find("#bd_snap_note a[href]").AttrOr("href", "")
		res.ContactPhone = extractor_util.VisitSourceUrlAndParse(href)
	}

	res.Poster = dom.Find(".user-info strong").Text()
	if res.Poster == "" {
		res.Poster = dom.Find(".user-info a.block").Text()
	}

	// 解析服务内容，服务范围
	dom.Find("ul.bx-meta-list li.feature-content").Each(func(_ int, metaItem *goquery.Selection) {
		metaLabel := metaItem.Find("label").Text()
		if strings.Contains(metaLabel, "服务内容") {
			res.ServiceContent = metaItem.Find("small.tag a").Text()
		} else if strings.Contains(metaLabel, "服务范围") {
			res.ServiceRange = metaItem.Find("small.tag a").Text()
		}
	})

	// 解析公司名，联系人
	dom.Find("ul.bx-meta-list li.meta-item").Each(func(_ int, metaItem *goquery.Selection) {
		metaLabel := metaItem.Find("label").Text()
		if strings.Contains(metaLabel, "公司名称") {
			res.CompanyName = metaItem.Find("div").Text()
		} else if strings.Contains(metaLabel, "联系人") {
			res.ContactPeople = metaItem.Find("div").Text()
		} else if strings.Contains(metaLabel, "微信号") {
			res.ContactWeixin = metaItem.Find("div").Text()
		} else if strings.Contains(metaLabel, "更新时间") {
			res.ReleaseTime = metaItem.Find("div").Text()
		} else if strings.Contains(metaLabel, "服务内容") {
			res.ServiceContent = metaItem.Find("div").Text()
		} else if strings.Contains(metaLabel, "服务范围") {
			res.ServiceRange = metaItem.Find("div").Text()
		}
	})

	return res, nil
}
