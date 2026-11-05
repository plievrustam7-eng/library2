package notifications

import "fmt"

type Notifier interface {
	Notify(readerName, message string) error
}

type EmailNotifier struct {
	Email string
}

func (n *EmailNotifier) Notify(readerName, message string) error {
	fmt.Printf("[EMAIL → %s] %s: %s\n", n.Email, readerName, message)
	return nil
}

type SMSNotifier struct {
	Phone string
}

func (n *SMSNotifier) Notify(readerName, message string) error {
	fmt.Printf("[SMS → %s] %s: %s\n", n.Phone, readerName, message)
	return nil
}