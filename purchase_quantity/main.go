package main

import (
	"bytes"
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/an-jun/xuanwu-test/conf"
	"github.com/an-jun/xuanwu-test/sign"
)

type PurchaseQuantityRequest struct {
	BeginRow         int    `json:"begin_row"`
	Size             int    `json:"size"`
	SaleOrganization string `json:"sale_organization"`
}

func main() {
	url := "https://midapi-test.cpchina.cn/xuanwu/purchase_quantity"
	appId := conf.AppId
	appSecret := conf.AppSecret
	req := PurchaseQuantityRequest{
		// BeginRow:         0,
		Size:             100,
		SaleOrganization: "CN64",
	}
	resp := getData(url, appId, appSecret, req)
	io.Copy(os.Stdout, resp.Body)

}
func getData(url, appId, appSecret string, req PurchaseQuantityRequest) *http.Response {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := r.Intn(1000)
	nonce := strconv.Itoa(i)
	t := time.Now().Unix() * 1000
	timestamp := strconv.FormatInt(t, 10)

	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(req)
	if err != nil {
		println(err.Error())
	}
	body := b.String()
	signUtils := sign.SignUtils{}
	sign := signUtils.Sign(appId, appSecret, timestamp, nonce, body)
	url = url + "?app_id=" + appId + "&nonce=" + nonce + "&timestamp=" + timestamp + "&sign=" + sign
	println(url)
	res, _ := http.Post(url, "application/json; charset=utf-8", b)

	return res
}
