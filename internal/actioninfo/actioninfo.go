package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			log.Printf("parse error: %v\n", err)
			continue
		}
		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("action info error: %v\n", err)
			continue
		}
		fmt.Println(info)
	}
}
