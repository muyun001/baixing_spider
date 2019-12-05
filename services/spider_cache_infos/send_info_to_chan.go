package spider_cache_infos

import (
	"baixing_spider/channels"
	"baixing_spider/structs/models"
	"github.com/kevin-zx/baidu-seo-tool/search"
)

// 将快照信息和日期放入channel
func SendCacheInfoToChan(cacheInfos *[]search.SearchResult, searchDate string) {
	for _, res := range *cacheInfos {
		cacheInfo := models.CacheInfoResult{
			CacheInfo: res.CacheUrl,
			Date:      searchDate,
		}

		channels.CacheInfoChan <- cacheInfo
	}
}
