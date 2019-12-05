package actions_test

import (
	"baixing_spider/actions"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

const dataFileShenghuoServicePc = "./test_html/shenghuo_service_pc.html"
const dataFileShenghuoServiceMobile = "./test_html/shenghuo_service_mobile.html"

func TestShenghuoServiceParsePc(t *testing.T) {
	contents, err := ioutil.ReadFile(dataFileShenghuoServicePc)
	if err != nil {
		t.Fatal("读取文件错误")
	}

	html := strings.Replace(string(contents), "\n", "", 1)
	res, err := actions.ShenghuoServiceParsePc(html)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}

func TestShenghuoServiceParseMobile(t *testing.T) {
	contents, err := ioutil.ReadFile(dataFileShenghuoServiceMobile)
	if err != nil {
		t.Fatal("读取文件错误")
	}

	html := strings.Replace(string(contents), "\n", "", 1)
	res, err := actions.ShenghuoServiceParseMobile(html)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}
