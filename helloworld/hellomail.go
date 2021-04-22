package main

import (
	"flag"
	"log"
	"os"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

const (
	server   = "imap.exmail.qq.com:993" //"imap.yeah.net:993"
	username = "taojianyu@nstc.com.cn"  //"tjy@yeah.net"
	password = "Windows9813"            //"890378" imap-mail.outlook.com:993 jy.tao@live.com 98132219
)

func main() {
	s := *flag.String("s", server, "imap server address")
	u := *flag.String("u", username, "user login name")
	p := *flag.String("p", password, "password")

	if len(os.Args) > 1 {
		s = os.Args[1]
		u = os.Args[2]
		p = os.Args[3]
	}
	log.Printf("s=%s u=%s p=%s", s, u, p)
	var c *client.Client
	var err error

	log.Printf("Connecting to server '%s'...\n", s)
	c, err = client.DialTLS(s, nil)

	//连接失败报错
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")
	//登陆
	if err := c.Login(u, p); err != nil {
		log.Fatal(err)
	}
	log.Println("Logged in")
	mailboxes := make(chan *imap.MailboxInfo, 20)
	go func() {
		c.List("", "*", mailboxes)
	}()
	//列取邮件夹
	for m := range mailboxes {

		mbox, err := c.Select(m.Name, false)
		if err != nil {
			log.Fatal(err)
		}
		to := mbox.Messages
		log.Printf("%s : %d", m.Name, to)
	}
}
