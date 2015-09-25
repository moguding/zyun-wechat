package controllers

import (
	"crypto/sha1"
	"encoding/hex"
	"zyun-wechat-api/models"

	"github.com/astaxie/beego"
)

type WechatController struct {
	beego.Controller
}

func (c *WechatController) URLMapping() {
	c.Mapping("Bind", c.Bind)
}

func (c *WechatController) Bind() {
	mailStr := c.Ctx.Input.Params[":mail"]
	passwdStr := c.Ctx.Input.Params[":passwd"]
	wechatStr := c.Ctx.Input.Params[":wechat"]
	user, err := models.GetZyUserAccountByMail(mailStr)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {

		tmpPasswd := user.HashedPasswd
		tmpPasswdSalt := user.PasswdSalt

		var data string = passwdStr + "{" + tmpPasswdSalt + "}"
		// SHA1 hash
		hash := sha1.New()
		hash.Write([]byte(data))
		hashBytes := hash.Sum(nil)
		hexString := hex.EncodeToString(hashBytes)
		if tmpPasswd == hexString { //密码正确
			v := &models.ZyWechat{
				UserId: user.Id,
				Wechat: wechatStr,
			}
			models.AddZyWechat(v)
		}
	}

}
