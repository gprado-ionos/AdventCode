package main

import (
	"regexp"
	"strings"
)

type Record struct {
	springStatus        string
	groupDamagedSprings []string
}

func buildRecords(entry []string) []Record {
	var records []Record
	for _, line := range entry {
		entryArr := strings.Split(line, " ")
		groupDamagedSprings := strings.Split(entryArr[1], ",")
		record := Record{entryArr[0], groupDamagedSprings}
		records = append(records, record)
	}
	return records
}

func getPossibleArrangements(records []Record) int {
	for _, record := range records {
		regexp.MustCompile(`[?]+`).FindAllStringIndex(record.springStatus, -1)
		//for i := 0; i < len(record.groupDamagedSprings); i++ {
		//	continuousDamagedSprings, err := strconv.Atoi(record.groupDamagedSprings[i])
		//	if err != nil {
		//		return 0
		//	}
		//	regex := "^[#]{"+record.groupDamagedSprings[i]+"}|[#]{"+record.groupDamagedSprings[i]+"}$|.[#]{"+record.groupDamagedSprings[i]+"}."
		//	index := regexp.MustCompile(regex).FindAllStringIndex(record.springStatus, -1)
		//	if len(index) > 1 {
		//		return 0
		//	}
		//
		//
		//}
		//fmt.Println(record)
	}
	return 0
}
