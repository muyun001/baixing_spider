package extractors

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

const shenghuoDataFilePc = "./test_html/pc.html"
const shenghuoDataFileMobile = "./test_html/mobile.html"

func TestBaixingShenghuoExtractorPc(t *testing.T) {
	contents, err := ioutil.ReadFile(shenghuoDataFilePc)
	if err != nil {
		t.Fatal("读取文件错误")
	}

	html := strings.Replace(string(contents), "\n", "", 1)
	res, err := BaixingShenghuoExtractorPc(html)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}

func TestBaixingShenghuoExtractorMobile(t *testing.T) {
	contents, err := ioutil.ReadFile(shenghuoDataFileMobile)
	if err != nil {
		t.Fatal("读取文件错误")
	}

	html := strings.Replace(string(contents), "\n", "", 1)
	res, err := BaixingShenghuoExtractorMobile(html)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}
