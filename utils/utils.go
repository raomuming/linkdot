package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"log"
	"net/http"
)

const (
	appid     = "wx577c587e3ea2cbd5"
	appsecret = "30fc0179adb73fcc040b5e64da9fe0a9"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// https://www.liwenbin.com/2018/02/%E3%80%90%E5%8E%9F%E3%80%91%E5%BC%80%E5%8F%91%E5%BE%AE%E4%BF%A1%E5%B0%8F%E7%A8%8B%E5%BA%8F%E4%B8%AD%E8%8E%B7%E5%8F%96unionid%E5%A4%B1%E8%B4%A5-%E9%99%84golang%E4%B8%8Ephp%E7%A4%BA/
func DecryptWxUserData(encryptedData, sessionKey, iv []byte) (string, error) {
	var aesBlockDecrypter cipher.Block
	aesBlockDecrypter, err := aes.NewCipher(sessionKey)
	if err != nil {
		return "", err
	}
	decrypted := make([]byte, len(encryptedData))
	aesDecrypter := cipher.NewCBCDecrypter(aesBlockDecrypter, iv)
	aesDecrypter.CryptBlocks(decrypted, encryptedData)

	return string(decrypted), nil
}

func GenerateWechatSessionGetUrl(code string) string {
	var buffer bytes.Buffer
	buffer.WriteString("https://api.weixin.qq.com/sns/jscode2session?appid=")
	buffer.WriteString(appid)
	buffer.WriteString("&secret=")
	buffer.WriteString(appsecret)
	buffer.WriteString("&js_code=")
	buffer.WriteString(code)
	buffer.WriteString("&grant_type=authorization_code")
	log.Println("generated url:", buffer.String())
	return buffer.String()
}
