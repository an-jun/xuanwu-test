package sign

import (
	"crypto/sha256"
	"encoding/hex"
	"sort"
	"strings"
)

type SignUtils struct {
}

func (SignUtils) Sign(appId, appSecret, timestamp, nonce, body string) string {
	arr := []string{body, timestamp, nonce, appSecret}
	sort.Strings(arr)
	str := strings.Join(arr, "")
	hash := sha256.New()
	hash.Write([]byte(str))
	md := hash.Sum(nil)
	sign := hex.EncodeToString(md)
	return sign
}
