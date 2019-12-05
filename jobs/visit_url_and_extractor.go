package jobs

import (
	"baixing_spider/channels"
	"baixing_spider/services/visit_url_and_extractor"
	"baixing_spider/structs/models"
	"errors"
	"time"
)

// 浏览页面并解析
func VisitUrlAndExtractor() (models.SaveCsvResult, error) {
	if len(channels.CacheInfoChan) == 0 {
		time.Sleep(time.Second)
		return models.SaveCsvResult{}, errors.New("暂时没有页面需要抓取, 休息1秒")
	}

	cacheInfo := <-channels.CacheInfoChan
	html, err := visit_url_and_extractor.VisitUrlAndFilter(cacheInfo.CacheInfo)
	if err != nil {
		return models.SaveCsvResult{}, err
	}

	if html == "" {
		return models.SaveCsvResult{}, nil
	}

	// TODO: "baixingShenghuo"
	extractor := ChooseExtractor("baixingShenghuo")
	result, err := extractor.Extractor(html)
	if err != nil {
		return models.SaveCsvResult{}, err
	}

	finalResult := models.SaveCsvResult{
		ExtractorResult: result,
		Url:             cacheInfo.CacheInfo,
		ChooseDate:      cacheInfo.Date,
	}

	return finalResult, nil
}

func ChooseExtractor(platform string) (visit_url_and_extractor.ExtractorInterface) {
	var extractor visit_url_and_extractor.ExtractorInterface
	switch platform {
	case "baixingShenghuo":
		extractor = new(visit_url_and_extractor.BaixingShenghuoExtractor)
	case "baixingShangwu":
		extractor = new(visit_url_and_extractor.BaixingShangwuExtractor)
	}

	return extractor
}
