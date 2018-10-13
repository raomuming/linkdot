package controller

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/http"

	_ "github.com/globalsign/mgo/bson"
	_ "github.com/raomuming/linkdot/auth"
	"github.com/raomuming/linkdot/model"
	"github.com/raomuming/linkdot/utils"
)

const (
	db         = "Linkdot"
	collection = "User"
	appid      = ""
	appsecret  = ""
)

func Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var wechatModel model.WechatModel
	if err := json.NewDecoder(r.Body).Decode(&wechatModel); err != nil {
		utils.ResponseWithJson(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	var buffer bytes.Buffer
	buffer.WriteString("https://api.weixin.qq.com/sns/jscode2session?appid=")
	buffer.WriteString(appid)
	buffer.WriteString("&secret=")
	buffer.WriteString(appsecret)
	buffer.WriteString("&js_code=")
	buffer.WriteString(wechatModel.Code)
	buffer.WriteString("&grant_type=authorization_code")

	resp, err := http.Get(buffer.String())
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

	var cryptedData []byte
	var sessionKey []byte
	var iv []byte

	if cryptedData, err = base64.StdEncoding.DecodeString(wechatModel.CryptedData); err != nil {
		utils.ResponseWithJson(w, http.StatusInternalServerError, "Decode cryptedData error")
		return
	}

	if sessionKey, err = base64.StdEncoding.DecodeString("sessionKey"); err != nil {
		utils.ResponseWithJson(w, http.StatusInternalServerError, "Decode sessionKey error")
		return
	}

	if iv, err = base64.StdEncoding.DecodeString(wechatModel.Iv); err != nil {
		utils.ResponseWithJson(w, http.StatusInternalServerError, "Decode iv error")
		return
	}

	var decoded string
	if decoded, err = utils.DecryptWxUserData(cryptedData, sessionKey, iv); err != nil {
		utils.ResponseWithJson(w, http.StatusInternalServerError, "Decrypt Wx dser data error")
		return
	}

	var userInfo map[string]*json.RawMessage
	if err := json.Unmarshal([]byte(decoded), &userInfo); err != nil {
		utils.ResponseWithJson(w, http.StatusInternalServerError, "")
	}
}
