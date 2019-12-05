package visit_url_and_extractor

import (
	"baixing_spider/services/visit_url_and_extractor/extractors"
	"baixing_spider/structs/models"
	"fmt"
)

type ExtractorInterface interface {
	Extractor(html string) (models.HtmlExtractorResult, error)
}

type BaixingShenghuoExtractor struct{}
type BaixingShangwuExtractor struct{}

func (baixingShenghuo BaixingShenghuoExtractor) Extractor(html string) (models.HtmlExtractorResult, error) {
	result, err := extractors.BaixingShenghuoExtractorPc(html)
	if err != nil {
		fmt.Println(err)
		return models.HtmlExtractorResult{}, err
	}

	if result.Title == "" && result.CompanyName == "" && result.ServiceContent == "" {
		result, err = extractors.BaixingShenghuoExtractorMobile(html)
		if err != nil {
			fmt.Println(err)
			return models.HtmlExtractorResult{}, err
		}
	}

	return result, nil
}

func (baixingShangwu BaixingShangwuExtractor) Extractor(html string) (models.HtmlExtractorResult, error) {
	result, err := extractors.BaixingShangwuExtractorPc(html)
	if err != nil {
		fmt.Println(err)
		return models.HtmlExtractorResult{}, err
	}

	if result.Title == "" && result.CompanyName == "" && result.ServiceContent == "" {
		result, err = extractors.BaixingShangwuExtractorMobile(html)
		if err != nil {
			fmt.Println(err)
			return models.HtmlExtractorResult{}, err
		}
	}

	return result, nil
}
