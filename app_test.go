package gofs_api

import (
	"fmt"
	"testing"
)

var token string = "08e1c5869a2a4d1fbd10e23719794d44"

func Test_App_GetList(t *testing.T) {
	client := Create(Option{
		Http:        "https",
		Ip:          "approve.cftzqinzhou.com",
		Port:        443,
		ContextPath: "gofs",
	})
	list := client.GetAppService().GetList(token)
	fmt.Println(list)
}

func Test_App_GetPageList(t *testing.T) {
	client := Create(Option{
		Http:        "https",
		Ip:          "approve.cftzqinzhou.com",
		Port:        443,
		ContextPath: "gofs",
	})
	pageData := client.GetAppService().GetPageList(token, 1, 10, "")
	fmt.Printf("%v\n", pageData)
}
