package actions

import (
	"baixing_spider/structs/models"
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/panwenbin/ghttpclient"
	"io"
	"os"
	"strings"
	"time"
)

// 通过cacheUrl抓取网页内容，然后发送到channel
func ShenghuoServiceCrawlAndSendToChan() {
	cacheUrlChan := make(chan string, 10000)
	i := 1
	fp, err := os.Open(shenghuoServiceCacheUrlsFile) // 获取文件指针
	if err != nil {
		fmt.Println(err)
	}
	defer fp.Close()
	bufReader := bufio.NewReader(fp)

	file, err := os.OpenFile(shenghuoServiceResultFile, os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open file is failed, err: ", err)
	}
	defer file.Close()
	w := csv.NewWriter(file)

	go func() {
		for {
			cacheUrl, _, err := bufReader.ReadLine() // 按行读
			if err != nil {
				if err == io.EOF {
					err = nil
					break
				}
			} else {
				cacheUrlChan <- string(cacheUrl)
			}
		}
	}()

	for {
		if len(cacheUrlChan) == 0 {
			fmt.Println("暂无任务，休息1秒")
			time.Sleep(time.Second)
			continue
		}

		url := <-cacheUrlChan
		body, err := ghttpclient.Get(url, nil).TryUTF8ReadBodyClose()
		if err != nil {
			fmt.Println(err)
		}

		if !strings.Contains(string(body), "服务内容") && !strings.Contains(string(body), "服务范围") && !strings.Contains(string(body), "公司名称") {
			fmt.Println("此网页是列表页, url:", url)
			return
		}

		if strings.Contains(string(body), "很遗憾，你来晚了") && strings.Contains(string(body), "这条信息已经搞定了") {
			fmt.Println("内容已过期, url:", url)
			return
		}

		urlAndHtml := models.UrlAndHtml{
			Url:  url,
			Html: string(body),
		}

		//channels.ShenghuoServiceHtmlNeedParseChan <- urlAndHtml
		//
		//urlAndHtml := <-channels.ShenghuoServiceHtmlNeedParseChan

		res, err := ShenghuoServiceParsePc(urlAndHtml.Html)
		if err != nil {
			fmt.Println(err)
		}

		if res.Title == "" && res.CompanyName == "" && res.ServiceContent == "" {
			res, err = ShenghuoServiceParseMobile(urlAndHtml.Html)
			if err != nil {
				fmt.Println(err)
			}
		}

		err = w.Write([]string{res.Title, res.ReleaseTime, res.CompanyName, res.ServiceContent, res.ServiceRange, res.ContactPeople, res.ContactPhone, res.ContactWeixin, res.Poster, urlAndHtml.Url})
		w.Flush()
		if err != nil {
			fmt.Println("写入数据到csv出错", err)
		}

		fmt.Println(i)
		fmt.Println("数据写入文件成功", fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s, %s", res.Title, res.ReleaseTime, res.CompanyName, res.ServiceContent,
			res.ServiceRange, res.ContactPeople, res.ContactPhone, res.ContactWeixin, res.Poster))
		i ++
	}
}

//func crawlAndSendChan(url string) {
//	body, err := ghttpclient.Get(url, nil).TryUTF8ReadBodyClose()
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	if !strings.Contains(string(body), "服务内容") && !strings.Contains(string(body), "服务范围") && !strings.Contains(string(body), "公司名称") {
//		fmt.Println("此网页是列表页, url:", url)
//		return
//	}
//
//	if strings.Contains(string(body), "很遗憾，你来晚了") && strings.Contains(string(body), "这条信息已经搞定了") {
//		fmt.Println("内容已过期, url:", url)
//		return
//	}
//
//	urlAndHtml := models.UrlAndHtml{
//		Url:  url,
//		Html: string(body),
//	}
//
//	//channels.ShenghuoServiceHtmlNeedParseChan <- urlAndHtml
//	//
//	//urlAndHtml := <-channels.ShenghuoServiceHtmlNeedParseChan
//
//	res, err := ShenghuoServiceParsePc(urlAndHtml.Html)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	if res.Title == "" && res.CompanyName == "" && res.ServiceContent == "" {
//		res, err = ShenghuoServiceParseMobile(urlAndHtml.Html)
//		if err != nil {
//			fmt.Println(err)
//		}
//	}
//
//	_ = w.Write([]string{res.Title, res.ReleaseTime, res.CompanyName, res.ServiceContent, res.ServiceRange, res.ContactPeople, res.ContactPhone, res.ContactWeixin, res.Poster, urlAndHtml.Url})
//	w.Flush()
//	fmt.Println(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s, %s", res.Title, res.ReleaseTime, res.CompanyName, res.ServiceContent,
//		res.ServiceRange, res.ContactPeople, res.ContactPhone, res.ContactWeixin, res.Poster))
//}
