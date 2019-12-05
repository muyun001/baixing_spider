package jobs

import (
	"baixing_spider/services/spider_cache_infos"
	"fmt"
	"time"
)

// 抓取快照信息
func SpiderCacheInfos(keyword string, forwardDateStartNum int, maxDateNum int, rn int) {
	isIncluded, err := spider_cache_infos.IsIncluded(keyword)
	if err != nil {
		fmt.Println(err)
		return
	}

	if !isIncluded {
		fmt.Println("此关键词未收录:", keyword)
		return
	}

	// 对日期循环查询
	for {
		//设置日期向前推移的最大天数
		if forwardDateStartNum > maxDateNum {
			break
		}

		pageNum := 1                                                                      // 从第一页开始查询
		searchDate := time.Now().AddDate(0, 0, -forwardDateStartNum).Format("2006-01-02") // 设置查询日期

		// 对页码循环查询
		for {
			baiduResults, haveNextPage, err := spider_cache_infos.BaiduSearchResult(keyword, pageNum, rn, searchDate)
			if err != nil {
				fmt.Println(err)
				return
			}

			if len(*baiduResults) == 0 {
				break
			}

			spider_cache_infos.SendCacheInfoToChan(baiduResults, searchDate)
			if haveNextPage {
				pageNum ++
			}
		}
		forwardDateStartNum ++
	}
}
