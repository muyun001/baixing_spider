package actions_test

import (
	"baixing_spider/actions"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

const dataFileBaiduPc = "./test_html/baidu_pc.html"
const dataFileBaiduMobile = "./test_html/baidu_mobile.html"

func TestParseBaiduPc(t *testing.T) {
	contents, err := ioutil.ReadFile(dataFileBaiduPc)
	if err != nil {
		t.Fatal("读取文件错误")
	}

	html := strings.Replace(string(contents), "\n", "", 1)
	res, err := actions.Parse(html)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}

func TestParseBaiduMobile(t *testing.T) {
	contents, err := ioutil.ReadFile(dataFileBaiduMobile)
	if err != nil {
		t.Fatal("读取文件错误")
	}

	html := strings.Replace(string(contents), "\n", "", 1)
	res, err := actions.Parse(html)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}