package imapclient

import (
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-imap"
)

type ImapClient struct {
	Host string
	Port string
	User string
	Password string
}

func (ic *ImapClient) Connect() (c *client.Client, err error) {
	// Connect to server
	c, err = client.DialTLS(ic.Host + ":" + ic.Port, nil)
	if err != nil {
		return nil, err
	}

	// Login
	if err := c.Login(ic.User, ic.Password); err != nil {
		return nil, err
	}

	return c, err
}

func (ic *ImapClient) FetchAll(folder string) (result []*imap.Message, err error) {
	c, err := ic.Connect()
	defer c.Logout()

	mbox, err := c.Select(folder, false)
	if err != nil {
		return result, err
	}

	set := new(imap.SeqSet)
	set.AddRange(1, mbox.Messages)

	messages := make(chan *imap.Message, 100)
	done := make(chan error, 1)
	go func() {
		done <- c.Fetch(set, []string{imap.EnvelopeMsgAttr}, messages)
	}()

	for msg := range messages {
		result = append(result, msg)
	}

	return result, err
}

func (ic *ImapClient) MailboxFolders() (folders []string, err error) {
	c, err := ic.Connect()
	defer c.Logout()

	// List mailboxes
	mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1)
	go func () {
		done <- c.List("", "*", mailboxes)
	}()

	for m := range mailboxes {
		folders = append(folders, m.Name)
	}
	return folders, nil
}