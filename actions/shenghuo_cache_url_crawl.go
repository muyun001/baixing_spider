package actions

import (
	"baixing_spider/global"
	"baixing_spider/structs/logics"
	"encoding/csv"
	"fmt"
	"github.com/kevin-zx/baidu-seo-tool/search"
	"os"
	"strings"
	"time"
)

var shenghuoServiceResultFile = "./data/shenghuo_service_result.csv"
var shenghuoServiceCacheUrlsFile = "./data/shenghuo_service_cache_urls.txt"
var basicKeyword = "site:suzhou.baixing.com inurl:%s"
var allInUrls = []string{
	"/baomu/",
	"/canyin/",
	"/wupinhuishou/",
	"/banjia/",
	"/baojieqingxi/",
	"/xiyihuli/",
	"/jiadianweixiu/m35854/",
	"/kaisuo/",
	"/weixiu/",
	"/fangwuweixiu/",
	"/jiadianweixiu/",
	"/jiajuweixiu/",
	"/diannaoweixiu/",
	"/shoujiweixiu/",
	"/shumaweixiu/",
	"/jiatingzhuangxiu/",
	"/ruanzhuang/",
	"/jiancaizhuangshi/",
	"/zhuangxiu/",
	"/chaijiu/",
	"/jiatingzhuangxiu/m179300/",
	"/gerenzuche/",
	"/peijiapeilian/",
	"/peijiafuwu/",
	"/jiaxiaofuwu/",
	"/qichebaoyang/",
	"/baoxianfuwu/m178815/",
	"/sheyingfuwu/",
	"/siyi/",
	"/siyi/m36001/",
	"/qianzhengfuwu/",
	"/jiudianfuwu/",
	"/lvxingshe/",
	"/jipiaofuwu/",
	"/yule/",
	"/yundongjianshen/",
	"/yangshengbaojian/",
	"/xianhualipin/",
	"/bendimingzhan/",
	"/quanxinshangjia/",
	"/binzang/",
	"/gongyijianding/",
	"/zixun/",
	"/qitafuwu/",
}

func ShenghuoServiceCrawlCacheUrls() {
	file, err := os.OpenFile(shenghuoServiceCacheUrlsFile, os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open file is failed, err: ", err)
	}
	defer file.Close()
	w := csv.NewWriter(file)

	for _, inUrl := range allInUrls {
		forwardDateNum := 0 // 日期向前推移的天数
		needSearchKeyword := fmt.Sprintf(basicKeyword, inUrl)

		// 判断是否搜索到内容
		webc, err := search.GetBaiduPCSearchHtmlWithRN(needSearchKeyword, 1, 10)
		if err != nil {
			fmt.Println(err)
		}
		if strings.Contains(webc, "很抱歉，没有找到与") {
			continue
		}

		// 对日期循环查询
		for {
			//设置日期向前推移的最大天数
			if forwardDateNum > logics.DAYS_时间向前推移的最大天数 {
				break
			}

			pageNum := 1
			startDate := time.Now().AddDate(0, 0, -forwardDateNum).Format("2006-01-02")

			// 对页码循环查询
			for {
				// 设置日期抓取
				webc, err := search.GetBaiduPCSearchHtmlWithRNAndTimeDayInterval(needSearchKeyword, pageNum, 50, startDate)
				if err != nil {
					return
				}

				baiduResults, err := search.ParseBaiduPCSearchResultHtml(webc)
				if err != nil {
					fmt.Println("GetBaiduPcResultsByKeywordAndSearchDay error", err)
				}

				fmt.Println("writeCount:", global.WriteCount, "needSearchKeyword:", needSearchKeyword, "forwardDateNum:", forwardDateNum, "pageNum:", pageNum, "baiduResults lenth:", len(*baiduResults))
				if len(*baiduResults) == 0 {
					break
				}

				for i, res := range *baiduResults {
					// 取每日数据的前10条
					if i == 10 {
						break
					}
					_ = w.Write([]string{res.CacheUrl})
					w.Flush()
					fmt.Println("CacheUrl写入成功：", res.CacheUrl)
				}

				if strings.Contains(webc, "下一页") {
					pageNum ++
				} else {
					break
				}
			}
			forwardDateNum ++
		}
	}
}
