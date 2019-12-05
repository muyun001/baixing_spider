package extractors

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

const shangwuDataFilePc = "./test_html/pc.html"
const shangwuDataFileMobile = "./test_html/mobile.html"

func TestBaixingShangwuExtractorPc(t *testing.T) {
	contents, err := ioutil.ReadFile(shangwuDataFilePc)
	if err != nil {
		t.Fatal("读取文件错误")
	}

	html := strings.Replace(string(contents), "\n", "", 1)
	res, err := BaixingShangwuExtractorPc(html)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}

func TestBaixingShangwuExtractorMobile(t *testing.T) {
	contents, err := ioutil.ReadFile(shangwuDataFileMobile)
	if err != nil {
		t.Fatal("读取文件错误")
	}

	html := strings.Replace(string(contents), "\n", "", 1)
	res, err := BaixingShangwuExtractorMobile(html)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}
