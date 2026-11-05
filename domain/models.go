package domain

import "fmt"

// Publication описывает книгу в системе
type Publication struct {
	Uid      int    `json:"uid"`
	Name     string `json:"title_name"`
	Author   string `json:"author_name"`
	Release  int    `json:"release_year"`
	Borrowed bool   `json:"borrowed"`
	HolderID int    `json:"holder_id"`
}

// Client описывает посетителя библиотеки
type Client struct {
	Uid     int    `json:"uid"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Enabled bool   `json:"enabled"`
}

func (p *Publication) String() string {
	return fmt.Sprintf("Книга: «%s» — %s (%d г.)", p.Name, p.Author, p.Release)
}

func (p *Publication) Take(clientID int) error {
	if p.Borrowed {
		return fmt.Errorf("издание '%s' уже находится у читателя с ID %d", p.Name, p.HolderID)
	}
	p.Borrowed = true
	p.HolderID = clientID
	return nil
}

func (p *Publication) HandIn() error {
	if !p.Borrowed {
		return fmt.Errorf("издание '%s' уже возвращено на полку", p.Name)
	}
	p.Borrowed = false
	p.HolderID = 0
	return nil
}

func (c *Client) String() string {
	return fmt.Sprintf("%s %s [ID: %d]", c.Surname, c.Name, c.Uid)
}

func (c *Client) Block() {
	c.Enabled = false
}

func (c *Client) Allow() {
	c.Enabled = true
}
