package jobs

import (
	"fmt"
)

// 拼接关键词主逻辑
func SpliceKeyword(siteCity string, domain string, category string) string {
	keyword := fmt.Sprintf("site:%s.%s inurl:%s", siteCity, domain, category)
	return keyword
}
