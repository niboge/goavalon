package handle

import (
	"crypto/rand"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"avalon/conf"
	"avalon/db"
	"avalon/plugin"
	"avalon/util"
	"fmt"
	"log"
	"time"
)

type WxCode struct {
	Code string `json:"code"`
}

type WxSessionKey struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
}

type WechatUserData struct {
	Iv            string `json:"iv"`            // 加密初始向量
	EncryptedData string `json:"encryptedData"` // 加密数据
	ThirdKey      string `json:"thirdKey"`      // 第三方key
}

type JsonUserInfo struct {
	OpenID    string `json:"openId"`
	NickName  string `json:"nickName"`
	Gender    int    `json:"gender"`
	City      int    `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	AvatarUrl string `json:"avatarurl"`
}

func LoginWechatUser(req *http.Request, config *conf.Config, uid int, pwd string) (int, string) {
	return 404, "kaixin"
	decoder := json.NewDecoder(req.Body)
	var wxcode WxCode
	// var user db.User
	err := decoder.Decode(&wxcode)
	if err != nil {
		log.Fatal(err)
	}

	if true {
		now := time.Now()
		// userUpdateDate := u.UpdateDate
		// subTime := now.Sub(userUpdateDate)
		// days := subTime.Hours() / 24

		// 将用户信息存放在session当中,并返回第三方sessionkey，防止官方session在网络中传输

		// 保存用户信息
		// session.Set(thirdKey, "userInfo", u)

		rspStr := fmt.Sprintf("{'thirdKey':'%s'}", now)
		return 200, rspStr
	} else {
		// thirdKey := createThirdPatyKey(wxsessionKey, u, session)
		rspStr := fmt.Sprintf(`{"errorInfo":"userinfo need to register"}`)
		return 404, rspStr
	}
}

func RegisterWechatUser(req *http.Request, config *conf.Config, session *plugin.SessionManager) (int, string) {
	decoder := json.NewDecoder(req.Body)
	var wechatUserData WechatUserData
	err := decoder.Decode(&wechatUserData)
	if err != nil {
		log.Fatal(err)
	}
	wxsession, has := session.Get("wxsessionKey")
	// 转换为wxsession
	var twxsession = wxsession.(*WxSessionKey)
	fmt.Println(twxsession)
	if has == true {
		// 解密加密信息
		wxbiz := util.WxBizDataCrypt{AppID: config.Wechat.APPID, SessionKey: twxsession.SessionKey}
		jsonUserInfo, err := wxbiz.Decrypt(wechatUserData.EncryptedData, wechatUserData.Iv, true)
		if err != nil {
			log.Fatalln(err)
		}
		tJSONUserInfo := jsonUserInfo.(string)
		var userinfo JsonUserInfo
		json.Unmarshal([]byte(tJSONUserInfo), &userinfo)
		// 查看是更新还是插入
		var user db.User
		u, has := user.GetUserByOpenId(userinfo.OpenID)
		if has == true {
			u.AvatarURL = userinfo.AvatarUrl
			u.NickName = userinfo.NickName
			var gender string
			if userinfo.Gender == 0 {
				gender = "神秘性别"
			} else if userinfo.Gender == 1 {
				gender = "男"
			} else {
				gender = "女"
			}
			u.Gender = gender
			u.Update()
			session.Set("userInfo", u)
			rsp := fmt.Sprintf("{'thirdKey':'%s'}", wechatUserData.ThirdKey)
			return 200, rsp

		} else {
			var gender string
			if userinfo.Gender == 0 {
				gender = "神秘性别"
			} else if userinfo.Gender == 1 {
				gender = "男"
			} else {
				gender = "女"
			}
			newUser := db.User{OpenID: userinfo.OpenID, NickName: userinfo.NickName, AvatarURL: userinfo.AvatarUrl, Gender: gender}
			session.Set("userInfo", newUser)
			newUser.Insert()
			rsp := fmt.Sprintf("{'thirdKey':'%s'}", wechatUserData.ThirdKey)
			return 200, rsp
		}
	} else {
		rsp := fmt.Sprintf("{'errorInfo':'You Must Be wx.login'}")
		return 404, rsp
	}

}

func getWxSessionCode(wxcode *WxCode, config *conf.Config) *WxSessionKey {
	// 获取微信code
	var wxKey WxSessionKey
	wxSessionAddr := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		config.Wechat.APPID, config.Wechat.AppSecret, wxcode.Code)
	rsp, err := http.Get(wxSessionAddr)
	defer rsp.Body.Close()
	b, _ := ioutil.ReadAll(rsp.Body)
	body := string(b)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal([]byte(body), &wxKey)
	return &wxKey
}

func createThirdPatyKey(wxsessionKey *WxSessionKey, u *db.User, session *plugin.SessionManager) string {
	// 创建第三方key
	b := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		log.Fatal(err)
		return ""
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	session.Set("wxsessionKey", wxsessionKey)
	return uuid
}
