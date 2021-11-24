package main

import (
	"fmt"
	"strings"
)

func GetIDfromUrl(urlID string)  string{
	splits := strings.Split(urlID, "/")
	if splits[len(splits)-1] == "other" || splits[len(splits)-1] == "unspecified" {
		return splits[len(splits)-2]
	}else {
		return splits[len(splits)-1]
	}
}

func PrintErr(err error)  {
	fmt.Println("################################")
	fmt.Println(err)
	fmt.Println("################################")
}