package main

import (

	"fmt"
	"time"
	"crypto/sha1"
	"encoding/hex"
	b64"encoding/base64"
	"net/http"
	"io/ioutil"

	"math/rand"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("from_michi_with_love")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func main() {
	url := "https://suite7.emarsys.net/api/v2/event"
	fmt.Println("URL:>", url)
	var timestamp = time.Now().Format(time.RFC3339)
	user := "XXXXXXXXXXXXXXXXXXXXX"
	secret := "XXXXXXXXXXXXXXXXXXXXXXXX"
	nonce  := RandStringRunes(rand.Intn(36))
	text := (nonce + timestamp + secret)
	h :=sha1.New()
	h.Write([]byte(text))
	sha1 :=hex.EncodeToString(h.Sum(nil))
	passwordDigest := b64.StdEncoding.EncodeToString([]byte(sha1))

	req, err := http.NewRequest("GET",url,nil)
	header := string(" UsernameToken Username=\"" + user + "\",PasswordDigest=\"" + passwordDigest + "\",Nonce=\""+ nonce + "\",Created=\""+ timestamp + "\"")

    fmt.Printf(nonce)
	fmt.Printf("\n")
	req.Header.Set("X-WSSE", header)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}