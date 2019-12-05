package channels

import "baixing_spider/structs/models"

var CacheInfoChan chan models.CacheInfoResult

func init() {
	CacheInfoChan = make(chan models.CacheInfoResult, 100000)
}
