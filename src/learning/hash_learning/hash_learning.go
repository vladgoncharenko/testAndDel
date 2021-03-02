package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	//data := "v4ndq314c2s5c28r:93b78bc88be90d93ac282e50ae569fdd"
	//encode := b64
	str := base64.StdEncoding.EncodeToString([]byte("v4ndq314c2s5c28r:93b78bc88be90d93ac282e50ae569fdd"))
	fmt.Println(str)
}

func md5Work(str string, f float64) string {
	md := md5.New()
	md.Write([]byte(str + strconv.FormatFloat(f, 'f', 6, 64)))
	return hex.EncodeToString(md.Sum(nil))
}
