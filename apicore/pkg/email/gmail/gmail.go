// @autor @jeffotoni
// actived count
// https://myaccount.google.com/lesssecureapps
package gemail

import (
    "bytes"

    "fmt"
    "html/template"
    "net/smtp"
    "os"

    "github.com/jeffotoni/gusermeli/apicore/pkg/fmts"
)

var (
    from     = os.Getenv("GMAIL_FROM")
    password = os.Getenv("GMAIL_PASSWORD")

    smtpHost = "smtp.gmail.com"
    smtpPort = "587"
)

var tmplEmail = `<table width="650" border="0" cellpadding="0" cellspacing="0" align="center" style="background-color: #FFF; font-family: Helvetica, Arial, sans-serif; line-height: 1.5; font-size: 16px; color: #545454; border: 1px solid #EAEAEA;">
    <tr>
        <td style="padding: 15px; text-align: center; border-bottom: 1px solid #F2F2F2" colspan="3">
            <img style="height: 100px;" src="https://s3.amazonaws.com/domain.com/setpassword/img/logo.png" alt="" />
        </td>
    </tr>
    <tr>
        <td style="padding: 50px 110px; text-align: center; border-bottom: 1px solid #F2F2F2">

            <div style="color: #269cea; font-size: 24px; font-weight: bold; margin: 0 auto 30px">Welcome to Meli</div>

            <p>Your account was created successfully, just below click to confirm your email.</p>

            <p>
                <span style="font-weight:bold;">E-mail:</span> <a href="" target="_blank" style="color:#269cea;text-decoration:none;">{{.USER_EMAIL}}</a>
            </p>

            <a style="text-decoration: none; color: #f6f6f6;" href="{{.LINK_PASSWORD_RECOVERY}}">
                <div style="background-color: #2CBE45; font-size: 16px; margin: 20px auto; padding: 12px; text-align: center; width: 380px; border-radius: 4px;">Confirm Email</div>
            </a>
        </td>
    </tr>
    <tr>
        <td style="padding:30px; font-size: 14px;text-align: center;">
            <a href="{{.DOMAIN}}" style="color: #969696; margin-bottom: 10px; display: block; vertical-align: middle; text-decoration: none;">{{.DOMAIN}}</a>
            Brz, 2021 - All rights reserved.
        </td>
    </tr>
</table>`

// somente uma simulacao enviar true
func SendUser(to []string, subject, user_email, link_pass, domain string) (err error) {
    return
    auth := smtp.PlainAuth("", from, password, smtpHost)

    var body bytes.Buffer
    mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
    body.Write([]byte(fmt.Sprintf("Subject: %s \n%s\n\n", subject, mimeHeaders)))

    t := template.Must(template.New("sendUser").Parse(tmplEmail))
    t.Execute(&body, struct {
        USER_EMAIL             string
        LINK_PASSWORD_RECOVERY string
        DOMAIN                 string
    }{
        USER_EMAIL:             user_email,
        LINK_PASSWORD_RECOVERY: link_pass,
        DOMAIN:                 domain,
    })

    err = smtp.SendMail(fmts.ConcatStr(smtpHost, ":", smtpPort), auth, from, to, body.Bytes())
    if err != nil {
        return
    }
    return
}
