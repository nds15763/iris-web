package HongShaoRouFrameWork

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type DBconnect struct {
	servername string `xml:"server:attr"`
	host       string `xml:"host"`
	port       string `xml:"port"`
	database   string `xml:"database"`
	charset    string `xml:"charset"`
	user       string `xml:"user"`
	password   string `xml:"password"`
}

func DbInit() {
	file, err := os.Open("E:\\tBeego\\HongShaoRouFrameWork\\DbConfig.xml")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := DBconnect{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
}
