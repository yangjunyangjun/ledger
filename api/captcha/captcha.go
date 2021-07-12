package captcha

import (
	"fmt"
	"github.com/mojocn/base64Captcha"
	"image/color"
)

var store = base64Captcha.DefaultMemStore


//  获取验证码
func GetCaptcha() (string, string) {
	// 生成字符串验证码
	driver := base64Captcha.DriverString{
		Height:          60,
		Width:           200,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          4,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}

	// 生成base64图片
	c := base64Captcha.NewCaptcha(driver.ConvertFonts(), store)

	// 获取
	id, b64s, err := c.Generate()
	if err != nil {
		fmt.Println("Register GetCaptchaPhoto get base64Captcha has err:", err)
		return "", ""
	}
	return id, b64s
}

// 验证验证码
func Verify(id string, val string) bool {
	if id == "" || val == "" {
		return false
	}
	// 同时在内存清理掉这个图片
	return store.Verify(id, val, true)
}
