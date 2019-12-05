package settings

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	ForwordDateStart int
	ForwordDateMax   int
)

func init() {
	checkEnv()
	LoadSetting()
}

func checkEnv() {
	_ = godotenv.Load()
	needChecks := []string{
		"FORWORD_DATE_START", "FORWORD_DATE_MEX",
	}

	for _, envKey := range needChecks {
		if os.Getenv(envKey) == "" {
			log.Fatalf("env %s missed", envKey)
		}
	}
}

func LoadSetting() {
	ForwordDateStart = loadIntFatal("FORWORD_DATE_START")
	ForwordDateMax = loadIntFatal("FORWORD_DATE_MEX")
}

func loadIntFatal(e string) int {
	intVar, err := strconv.Atoi(os.Getenv(e))
	if err != nil {
		log.Fatalf("env %s invalid\n", e)
	}

	return intVar
}
