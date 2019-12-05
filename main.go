package main

import (
	"baixing_spider/jobs"
	"baixing_spider/structs/logics"
	"encoding/csv"
	"fmt"
	"os"
)

// 分开
var subDomain = "suzhou"
var domain = "baixing.com"
var resultCsvFile = "./data/result.csv"
var allCategories = []string{
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

func main() {
	file, err := os.OpenFile(resultCsvFile, os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open file is failed, err: ", err)
	}
	defer file.Close()
	w := csv.NewWriter(file)
	//_ = w.Write([]string{"标题", " 发布时间", " 公司名", " 服务内容", " 服务范围", " 联系人", " 联系电话", " 微信号", " 发布者", " 选择日期", "快照地址"})

	for _, category := range allCategories {
		keyword := jobs.SpliceKeyword(subDomain, domain, category)
		go jobs.SpiderCacheInfos(keyword, logics.DATE_FORWORD_START, logics.DATE_FORWORD_MEX, logics.DATA_EVERY_PAGE_NUM)
		for {
			result, err := jobs.VisitUrlAndExtractor()
			if err != nil {
				continue
			}

			//jobs.SaveCsv(result)
			if result.Url == "" || (result.ExtractorResult.ContactPhone == "" && result.ExtractorResult.ContactWeixin == "" && result.ExtractorResult.ServiceContent == "") {
				continue
			}

			_ = w.Write([]string{result.ExtractorResult.Title,
				result.ExtractorResult.ReleaseTime,
				result.ExtractorResult.CompanyName,
				result.ExtractorResult.ServiceContent,
				result.ExtractorResult.ServiceRange,
				result.ExtractorResult.ContactPeople,
				result.ExtractorResult.ContactPhone,
				result.ExtractorResult.ContactWeixin,
				result.ExtractorResult.Poster,
				result.ChooseDate, result.Url})
			w.Flush()
			fmt.Println("文件写入成功：", result)
		}
	}
}
