package library

import (
	"errors"
	"fmt"

	"library-app/domain"
)

type Catalog struct {
	Items   []*domain.Publication
	Members []*domain.Client

	nextBookID   int
	nextReaderID int
}

func NewCatalog() *Catalog {
	return &Catalog{
		Items:   make([]*domain.Publication, 0),
		Members: make([]*domain.Client, 0),
	}
}

func (c *Catalog) RegisterBook(name, author string, year int) (*domain.Publication, error) {
	if name == "" || author == "" {
		return nil, errors.New("название и автор обязательны к заполнению")
	}
	
	// Поиск дубликатов классическим циклом по индексу
	for i := 0; i < len(c.Items); i++ {
		if c.Items[i].Name == name && c.Items[i].Author == author {
			return nil, fmt.Errorf("издание %s (%s) уже зарегистрировано", name, author)
		}
	}

	c.nextBookID++
	newItem := &domain.Publication{
		Uid:      c.nextBookID,
		Name:     name,
		Author:   author,
		Release:  year,
		Borrowed: false,
	}
	c.Items = append(c.Items, newItem)
	return newItem, nil
}

func (c *Catalog) RegisterClient(first, last string) (*domain.Client, error) {
	for i := 0; i < len(c.Members); i++ {
		if c.Members[i].Name == first && c.Members[i].Surname == last {
			return nil, fmt.Errorf("пользователь %s %s уже существует", first, last)
		}
	}

	c.nextReaderID++
	newClient := &domain.Client{
		Uid:     c.nextReaderID,
		Name:    first,
		Surname: last,
		Enabled: true,
	}
	c.Members = append(c.Members, newClient)
	return newClient, nil
}

func (c *Catalog) FetchBook(id int) (*domain.Publication, error) {
	for _, item := range c.Items {
		if item.Uid == id {
			return item, nil
		}
	}
	return nil, fmt.Errorf("издание с номером %d отсутствует", id)
}

func (c *Catalog) FetchClient(id int) (*domain.Client, error) {
	for _, member := range c.Members {
		if member.Uid == id {
			return member, nil
		}
	}
	return nil, fmt.Errorf("посетитель с номером %d не найден", id)
}

func (c *Catalog) CheckoutItem(bookID, readerID int) error {
	b, err := c.FetchBook(bookID)
	if err != nil {
		return err
	}
	r, err := c.FetchClient(readerID)
	if err != nil {
		return err
	}
	if !r.Enabled {
		return fmt.Errorf("аккаунт пользователя %s заблокирован", r.Surname)
	}
	return b.Take(readerID)
}

func (c *Catalog) ReturnItem(bookID int) error {
	b, err := c.FetchBook(bookID)
	if err != nil {
		return err
	}
	return b.HandIn()
}

func (c *Catalog) ListItems() []*domain.Publication {
	return c.Items
}

func (c *Catalog) ListMembers() []*domain.Client {
	return c.Members
}

func (c *Catalog) SyncBooks(data []*domain.Publication) {
	c.Items = data
	c.recalculateIDs()
}

func (c *Catalog) SyncClients(data []*domain.Client) {
	c.Members = data
	c.recalculateIDs()
}

func (c *Catalog) recalculateIDs() {
	highBook := 0
	for i := 0; i < len(c.Items); i++ {
		if c.Items[i].Uid > highBook {
			highBook = c.Items[i].Uid
		}
	}
	c.nextBookID = highBook

	highReader := 0
	for i := 0; i < len(c.Members); i++ {
		if c.Members[i].Uid > highReader {
			highReader = c.Members[i].Uid
		}
	}
	c.nextReaderID = highReader
}
