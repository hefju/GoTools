package Net
import (
    "fmt"
    "net/smtp"
    "strings"
//    "github.com/hefju/CBChareClient/myconfig"
)
func SendEmail(title ,content string){
    //  email := newEmail("89176125@qq.com;hvfju@gmail.com", "IBM峰会资料下载地址", "IBM峰会资料下载地址IBM峰会资料下载地址IBM峰会资料下载地址")
    email := newEmail("89176125@qq.com;hvfju@gmail.com", title, content)

    err := sendEmail(email)

    fmt.Println(err)
}


type Email struct {
    to      string "to"
    subject string "subject"
    msg     string "msg"
}

func newEmail(to, subject, msg string) *Email {
    return &Email{to: to, subject: subject, msg: msg}
}

func sendEmail(email *Email) error {
    auth := smtp.PlainAuth("", myconfig.EmailUSER,myconfig.EmailPASSWORD,myconfig.EmailHOST)
    sendTo := strings.Split(email.to, ";")
    done := make(chan error, 1024)

    go func() {
        defer close(done)
        for _, v := range sendTo {

            str := strings.Replace("From: "+myconfig.EmailUSER+"~To: "+v+"~Subject: "+email.subject+"~~", "~", "\r\n", -1) + email.msg

            err := smtp.SendMail(
            myconfig.EmailSERVER_ADDR,
            auth,
            myconfig.EmailUSER,
            []string{v},
            []byte(str),
            )
            done <- err
        }
    }()

    for i := 0; i < len(sendTo); i++ {
        <-done
    }

    return nil
}