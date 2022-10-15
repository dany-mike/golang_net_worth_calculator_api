package models

import (
	"errors"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
)

type Item struct {
	ID          uint64  `gorm:"primary_key;auto_increment" json:"id"`
	Title       string  `gorm:"size:255;not null;" json:"title"`
	Description string  `json:"description"`
	Price       float64 `gorm:"not null;" json:"price"`
	User        User    `json:"user"`
	UserID      uint32  `gorm:"not null" json:"user_id"`
}

func (p *Item) Prepare() {
	p.ID = 0
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
	p.Description = html.EscapeString(strings.TrimSpace(p.Description))
	p.User = User{}
}

func (p *Item) Validate() error {

	if p.Title == "" {
		return errors.New("required title")
	}
	if p.Price == 0 {
		return errors.New("required price")
	}
	if p.UserID < 1 {
		return errors.New("required user")
	}
	return nil
}

func (p *Item) SaveItem(db *gorm.DB) (*Item, error) {
	var err error
	err = db.Debug().Model(&Item{}).Create(&p).Error
	if err != nil {
		return &Item{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
		if err != nil {
			return &Item{}, err
		}
	}
	return p, nil
}

func (p *Item) FindAllItems(db *gorm.DB) (*[]Item, error) {
	var err error
	items := []Item{}
	err = db.Debug().Model(&Item{}).Limit(100).Find(&items).Error
	if err != nil {
		return &[]Item{}, err
	}
	if len(items) > 0 {
		for i, _ := range items {
			err := db.Debug().Model(&User{}).Where("id = ?", items[i].UserID).Take(&items[i].User).Error
			if err != nil {
				return &[]Item{}, err
			}
		}
	}
	return &items, nil
}

func (p *Item) FindItemByID(db *gorm.DB, pid uint64) (*Item, error) {
	var err error
	err = db.Debug().Model(&Item{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Item{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
		if err != nil {
			return &Item{}, err
		}
	}
	return p, nil
}

func (p *Item) UpdateItem(db *gorm.DB) (*Item, error) {
	var err error

	err = db.Debug().Model(&Item{}).Where("id = ?", p.ID).Updates(Item{Title: p.Title, User: p.User, Price: p.Price, Description: p.Description}).Error
	if err != nil {
		return &Item{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
		if err != nil {
			return &Item{}, err
		}
	}
	return p, nil
}

func (p *Item) DeleteItem(db *gorm.DB, pid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&Item{}).Where("id = ? and user_id = ?", pid, uid).Take(&Item{}).Delete(&Item{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Item not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
