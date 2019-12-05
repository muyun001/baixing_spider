package spider_cache_infos

import (
	"errors"
	"fmt"
	"github.com/kevin-zx/baidu-seo-tool/search"
	"strings"
)

// 获取百度的查询结果
func BaiduSearchResult(keyword string, pageNum int, rn int, searchDate string) (*[]search.SearchResult, bool, error) {
	webc, err := search.GetBaiduPCSearchHtmlWithRNAndTimeDayInterval(keyword, pageNum, rn, searchDate)
	if err != nil {
		return nil, false, errors.New(fmt.Sprintf("GetBaiduPCSearchHtmlWithRNAndTimeDayInterval error: %s", err))
	}

	baiduResults, err := search.ParseBaiduPCSearchResultHtml(webc)
	if err != nil {
		return nil, false, errors.New(fmt.Sprintf("ParseBaiduPCSearchResultHtml error: %s", err))
	}

	var haveNextPage bool
	if strings.Contains(webc, "下一页") {
		haveNextPage = true
	}

	return baiduResults, haveNextPage, nil
}
