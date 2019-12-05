package visit_url_and_extractor

import (
	"fmt"
	"github.com/panwenbin/ghttpclient"
	"strings"
)

// 访问快照地址并过滤掉无效网页
func VisitUrlAndFilter(url string) (string, error) {
	body, err := ghttpclient.Get(url, nil).TryUTF8ReadBodyClose()
	if err != nil {
		fmt.Println("VisitUrlAndFilter ghttpclient get error")
		return "", err
	}

	if !strings.Contains(string(body), "服务内容") && !strings.Contains(string(body), "服务范围") && !strings.Contains(string(body), "公司名称") {
		fmt.Println(fmt.Sprintf("此网页是列表页, url:", url))
		return "", nil
	}

	if strings.Contains(string(body), "很遗憾，你来晚了") && strings.Contains(string(body), "这条信息已经搞定了") {
		fmt.Println(fmt.Sprintf("内容已过期, url:", url))
		return "", nil
	}

	if strings.Contains(string(body), "来晚了已过期") {
		fmt.Println(fmt.Sprintf("内容已过期, url:", url))
		return "", nil
	}

	return string(body), nil
}
