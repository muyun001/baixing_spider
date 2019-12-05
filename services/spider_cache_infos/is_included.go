package spider_cache_infos

import (
	"errors"
	"github.com/kevin-zx/baidu-seo-tool/search"
	"strings"
)

// 判断网页是否有收录
func IsIncluded(keyword string) (bool, error) {
	webc, err := search.GetBaiduPCSearchHtmlWithRN(keyword, 1, 10)
	if err != nil {
		return false, errors.New("GetBaiduPCSearchHtmlWithRN error")
	}

	if strings.Contains(webc, "很抱歉，没有找到与") {
		return false, nil
	}

	return true, nil
}

