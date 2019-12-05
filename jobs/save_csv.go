package jobs

import (
	"baixing_spider/global"
	"baixing_spider/structs/models"
	"fmt"
)

func SaveCsv(result models.SaveCsvResult) {
	_ = global.W.Write([]string{result.ExtractorResult.Title,
		result.ExtractorResult.ReleaseTime,
		result.ExtractorResult.CompanyName,
		result.ExtractorResult.ServiceContent,
		result.ExtractorResult.ServiceRange,
		result.ExtractorResult.ContactPeople,
		result.ExtractorResult.ContactPhone,
		result.ExtractorResult.ContactWeixin,
		result.ExtractorResult.Poster,
		result.ChooseDate, result.Url})

	global.W.Flush()
	fmt.Println("文件写入成功：", result)
}
