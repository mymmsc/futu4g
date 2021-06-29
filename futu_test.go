package futu4g

import (
	"encoding/json"
	"fmt"
	"github.com/mymmsc/futu4g/api"
	"github.com/mymmsc/futu4g/protocol/Qot_Common"
	"github.com/mymmsc/futu4g/protocol/Qot_GetStaticInfo"
	"testing"
)

func Test(t *testing.T) {
	// 1. 建立网络连接
	config := api.Config{
		Address: "127.0.0.1:11111",
	}
	cli, err := api.NewFutuAPI(&config)
	if err != nil {
		fmt.Printf("NewFutuAPI err: %v\n", err)
	}
	defer cli.Close()
	// 2. 查询
	market := Qot_Common.QotMarket_QotMarket_CNSH_Security
	marketInt32 := int32(market)
	secType := Qot_Common.SecurityType_SecurityType_Eqty
	secTypeInt32 := int32(secType)
	req := Qot_GetStaticInfo.Request{
		C2S: &Qot_GetStaticInfo.C2S{
			Market:  &marketInt32,
			SecType: &secTypeInt32,
		},
	}
	//response := make(<-chan Qot_GetStaticInfo.Response)
	ch, err := cli.QotGetStaticInfo(&req)
	if err != nil {
		fmt.Printf("NewFutuAPI err: %v\n", err)
	}
	response := <-ch
	fmt.Printf("data = %v\n", response)
	var basicList []Qot_Common.SecurityStaticBasic
	infoList := response.S2C.StaticInfoList
	for _, item := range infoList {
		if item.Basic == nil {
			continue
		}
		basicList = append(basicList, *item.Basic)
	}
	fmt.Print("---------------------------------\n")
	data, err := json.Marshal(basicList)
	if err == nil {
		fmt.Printf("%s", data);
	}

}
