package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

const appSecret = "KPBQCINfXQKZs@&2"

func app_demo() {
	id := uuid.NewV4()
	ids := id.String()
	appId := ids
	// appId = "106579"
	appId = "9c712cd1-bb4b-492f-9c9f-4d56d4b3f09e"

	// ts := time.Now().Unix()
	// ts := 1638436226141
	ts := 1638351038
	// ts := req.TimeStamp
	// appId := req.AppId
	uid := 1234567
	fmt.Printf("app_id=%s&ts=%d, appSecret=%s\n", appId, ts, appSecret)
	strForSign := fmt.Sprintf("app_id=%s&ts=%d&uid=%d%s", appId, ts, uid, appSecret)
	fmt.Println(strForSign)
	signMd5 := md5.New()
	signMd5.Write([]byte(strForSign))
	sign := signMd5.Sum(nil)
	// fmt.Println(string(sign))
	fmt.Println(hex.EncodeToString(sign))

}
