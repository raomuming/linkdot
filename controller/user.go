package controller

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/globalsign/mgo/bson"
	_ "github.com/raomuming/linkdot/auth"
	"github.com/raomuming/linkdot/model"
	"github.com/raomuming/linkdot/utils"
)

const (
	db         = "Linkdot"
	collection = "User"
)

func Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var wechatModel model.WechatModel
	if err := json.NewDecoder(r.Body).Decode(&wechatModel); err != nil {
		utils.ResponseWithJson(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	log.Println("code:", wechatModel.Code)
	log.Println("iv:", wechatModel.Iv)
	log.Println("crypted data", wechatModel.CryptedData)

	resp, err := http.Get(utils.GenerateWechatSessionGetUrl(wechatModel.Code))
	if err != nil {
		utils.ResponseWithJson(w, http.StatusInternalServerError, "Send request to wechat fail")
		return
	}

	defer resp.Body.Close()
	// https://stackoverflow.com/questions/11066946/partly-json-unmarshal-into-a-map-in-go
	var result map[string]*json.RawMessage
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		utils.ResponseWithJson(w, http.StatusInternalServerError, "Server decode error")
		return
	}

	if _, ok := result["errcode"]; ok {
		log.Println("Has error code")
		utils.ResponseWithJson(w, http.StatusInternalServerError, "Wechat server returns error")
		return
	}

	//utils.ResponseWithJson(w, http.StatusOK, "ok")
	//return
	var str string
	if err := json.Unmarshal(*result["session_key"], &str); err != nil {
		utils.ResponseWithJson(w, http.StatusInternalServerError, "Server decode session_key error")
		return
	}

	var cryptedData []byte
	var sessionKey []byte
	var iv []byte

	if cryptedData, err = base64.StdEncoding.DecodeString(wechatModel.CryptedData); err != nil {
		utils.ResponseWithJson(w, http.StatusInternalServerError, "Decode cryptedData error")
		return
	}

	if sessionKey, err = base64.StdEncoding.DecodeString(str); err != nil {
		utils.ResponseWithJson(w, http.StatusInternalServerError, "Decode sessionKey error")
		return
	}

	if iv, err = base64.StdEncoding.DecodeString(wechatModel.Iv); err != nil {
		utils.ResponseWithJson(w, http.StatusInternalServerError, "Decode iv error")
		return
	}

	var decoded string
	if decoded, err = utils.DecryptWxUserData(cryptedData, sessionKey, iv); err != nil {
		utils.ResponseWithJson(w, http.StatusInternalServerError, "Decrypt Wx user data error")
		return
	}
	log.Println("decoded data:", decoded)

	// example of decoded data:
	//{"openId":"oy4oL0btY8jYT1H2GozDCPGCMeo8","nickName":"饶木明","gender":1,"language":"zh_CN","city":"Xiamen","province":"Fujian","country":"China","avatarUrl":"https://wx.qlogo.cn/mmopen/vi_32/ajNVdqHZLLAXV6Z240SrANvfCY27icW54epiaLUfUjicROb5XmUjhbHiaFGK1aszGkYC8icvDy9vqpByGUcAibib4hHjQ/132","watermark":{"timestamp":1539444581,"appid":"wx577c587e3ea2cbd5"}}

	var userInfo map[string]*json.RawMessage
	if err := json.Unmarshal([]byte(decoded), &userInfo); err != nil {
		utils.ResponseWithJson(w, http.StatusInternalServerError, "Decode userInfo error")
		return
	}

}
