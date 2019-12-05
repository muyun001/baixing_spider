package actions

import (
	"baixing_spider/global"
	"baixing_spider/structs/logics"
	"encoding/csv"
	"fmt"
	"github.com/kevin-zx/baidu-seo-tool/search"
	"github.com/panwenbin/ghttpclient"
	"os"
	"strings"
	"time"
)

const resultFile = "./data/shenghuo_result.csv"

// 抓取网页
func Crawl() {
	basicKeyword := "site:suzhou.baixing.com inurl:%s"
	allInUrls := []string{
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

	// 写数据到csv
	file, err := os.OpenFile(resultFile, os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open file is failed, err: ", err)
	}
	defer file.Close()
	w := csv.NewWriter(file)
	_ = w.Write([]string{"标题", " 发布时间", " 公司名", " 服务内容", " 服务范围", " 联系人", " 联系电话", " 微信号", " 发布者", " 选择日期", "快照地址"})

	// 开始循环抓取数据
	for _, url := range allInUrls {
		forwardDateNum := 0 // 日期向前推移的天数
		needSearchKeyword := fmt.Sprintf(basicKeyword, url)

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
			//if global.WriteCount > 100 {
			//	break
			//}

			//设置日期向前推移的最大天数
			if forwardDateNum > logics.DAYS_时间向前推移的最大天数 {
				break
			}

			pageNum := 1
			startDate := time.Now().AddDate(0, 0, -forwardDateNum).Format("2006-01-02")

			// 对页码循环查询
			for {
				// 不设置日期，全量抓取
				//webc, err := search.GetBaiduPCSearchHtmlWithRN(needSearchKeyword, pageNum, 50)
				//if err != nil {
				//	return
				//}

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

				for _, res := range *baiduResults {
					body, err := ghttpclient.Get(res.CacheUrl, nil).TryUTF8ReadBodyClose()
					if err != nil {
						fmt.Println(err)
					}

					if !strings.Contains(string(body), "服务内容") && !strings.Contains(string(body), "服务范围") && !strings.Contains(string(body), "公司名称") {
						fmt.Println("此网页是列表页, url:", res.CacheUrl)
						continue
					}

					if strings.Contains(string(body), "很遗憾，你来晚了") && strings.Contains(string(body), "这条信息已经搞定了") {
						fmt.Println("内容已过期, url:", res.CacheUrl)
						continue
					}

					if strings.Contains(string(body), "来晚了已过期") {
						fmt.Println("内容已过期, url:", res.CacheUrl)
						continue
					}

					result, err := Parse(string(body))
					if err != nil {
						fmt.Println(err)
					}

					if result.ReleaseTime == "" && result.CompanyName == "" && result.ServiceRange == "" && result.ServiceContent == "" && result.ContactPeople == "" && result.ContactPhone == "" {
						fmt.Println("非列表页没解析到内容, url:", res.CacheUrl)
						continue
					}

					result.StartDate = startDate
					_ = w.Write([]string{result.Title, result.ReleaseTime, result.CompanyName, result.ServiceContent, result.ServiceRange, result.ContactPeople, result.ContactPhone, result.ContactWeixin, result.Poster, result.StartDate, res.CacheUrl})
					w.Flush()
					fmt.Println("文件写入成功：", result)
					global.WriteCount ++
					//time.Sleep(time.Second)
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
