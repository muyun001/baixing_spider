package global

import (
	"encoding/csv"
	"fmt"
	"os"
)

var W *csv.Writer
var resultCsvFile = "./data/result.csv"

func init() {
	file, err := os.OpenFile(resultCsvFile, os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open file is failed, err: ", err)
	}
	defer file.Close()
	W := csv.NewWriter(file)
	_ = W.Write([]string{"标题", " 发布时间", " 公司名", " 服务内容", " 服务范围", " 联系人", " 联系电话", " 微信号", " 发布者", " 选择日期", "快照地址"})
}
