package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func json_demo() {
	type Info struct {
		EntityID uint64 `json:"entity_id"`
		ShopIdx  uint32 `json:"shop_idx"`
	}

	var entityID uint64 = 10114
	var moneyShopIdx uint32 = 1
	serverExtensionInfo := Info{
		EntityID: entityID,
		ShopIdx:  moneyShopIdx,
	}
	fmt.Printf("serverExtensionInfo:%+v\n", serverExtensionInfo)
	serverExtensionInfoJson, err := json.Marshal(serverExtensionInfo)
	fmt.Printf("err:%v\n", err)
	serverExtensionInfoIndentJson, err := json.MarshalIndent(serverExtensionInfo, "", " 	")
	fmt.Printf("Indent err:%v\n", err)
	if err != nil {
		log.Fatalf("Json marshaling failed:%s", err)
		return
	}
	fmt.Printf("err:%v\n", err)
	fmt.Printf("serverExtensionInfoJson:%s\n", serverExtensionInfoJson)
	fmt.Printf("serverExtensionInfoJson:%s\n", serverExtensionInfoIndentJson)
	fmt.Printf("serverExtensionInfoJson string:%s\n", string(serverExtensionInfoJson))
	info := &Info{}
	err = json.Unmarshal(serverExtensionInfoJson, info)
	if err != nil {
		log.Fatalf("Json marshaling failed:%s", err)
		return
	}
	fmt.Printf("info:%+v\n", info)
	return
}
