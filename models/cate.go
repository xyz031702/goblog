package models

import (
	"time"
	"github.com/ilibs/gosql"
)

type Cates struct {
	Id        int       `form:"id" json:"id" db:"id"`
	Name      string    `form:"name" json:"name" db:"name"`
	Desc      string    `form:"desc" json:"desc" db:"desc"`
	Domain    string    `form:"domain" json:"domain" db:"domain"`
	CreatedAt time.Time `form:"-" json:"created_at" db:"created_at"`
	UpdatedAt time.Time `form:"-" json:"updated_at" db:"updated_at"`
}

func (c *Cates) DbName() string {
	return "default"
}

func (c *Cates) TableName() string {
	return "cates"
}

func (c *Cates) PK() string {
	return "id"
}

func CateGetList(start int, num int) ([]*Cates, error) {
	var m = make([]*Cates, 0)
	start = (start - 1) * num
	err := gosql.Model(&m).OrderBy("id desc").Limit(num).Offset(start).All()
	if err != nil {
		return nil, err
	}
	return m, nil
}
