package lib

import (
	"bytes"
	"gopkg.in/gomail.v2"
	"ledger/conf"
	"ledger/sdk"
	"strings"
)

var MailChan = make(chan *Email, 100)

type Email struct {
	To       string       //收件人，必传
	Cc       string       //抄送人，非必传传
	Subject  string       // 主题 ，必传
	Msg      string       // 邮件主题 ，必传
	FileName string       // 附件名 ，非必传
	File     bytes.Buffer // 附件内容，非必传
}


// 初始化一个监听邮件的协程
func init() {
	go func() {
		for {
			if mail, ok := <-MailChan; ok {
				if err := mail.SendEmail(); err != nil {
					sdk.Log.Errorf("send mail error %x", err)
				}
				sdk.Log.Println("init monitor mail success")
			} else {
				panic("监听邮件失败")
			}

		}
	}()
}



//添加发邮件任务
func RegisterMail(m *Email) {
	go func() {
		MailChan <- m
	}()
}

func (e *Email) SendEmail() error {
	m := gomail.NewMessage()
	////发送人
	m.SetHeader("From", conf.Config.Email.User)
	//收件人
	to := strings.Split(e.To, ";")
	m.SetHeader("To", to...)
	////抄送
	if e.Cc != "" {
		cc := strings.Split(e.Cc, ";")
		m.SetHeader("Cc", cc...)
	}
	//主题
	m.SetHeader("Subject", e.Subject)
	//内容
	m.SetBody("text/html", e.Msg)

	//拿到token，并进行连接,第4个参数是填授权码
	t := gomail.NewDialer(conf.Config.Email.Host, conf.Config.Email.Port, conf.Config.Email.User, conf.Config.Email.Password)

	if err := t.DialAndSend(m); err != nil {
		sdk.Log.Errorf("send mail error :%x", err)
		return err
	}
	return nil
}
