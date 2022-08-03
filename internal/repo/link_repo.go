package repo

import (
	"fmt"

	"github.com/adnanbrq/slugify/internal/db"
	"github.com/adnanbrq/slugify/internal/entity"
	"github.com/adnanbrq/slugify/internal/errs"
	"gorm.io/gorm"
)

type LinkRepo struct{}

func (l LinkRepo) LinkExists(slug string) bool {
	var count int64
	if !db.IsConnected() {
		return false
	}

	res := db.DB.Model(&entity.Link{}).Select("id").Where("slug = ?", slug).Count(&count)
	if res.Error != nil {
		return false
	}

	return count > 0
}

func (LinkRepo) Create(link *entity.Link) error {
	if !db.IsConnected() {
		fmt.Println("[LinkRepo:Create] Not connected")
		return errs.ErrDbNotConnected
	}

	if res := db.DB.Create(link); res.Error != nil {
		fmt.Printf("%v", res.Error)
		return errs.ErrDbError
	}

	return nil
}

func (repo LinkRepo) GetBySlug(slug string) (*entity.Link, error) {
	if !db.IsConnected() {
		fmt.Println("[LinkRepo:GetBySlug] Not connected")
		return nil, errs.ErrDbNotConnected
	}

	if exists := repo.LinkExists(slug); exists == false {
		fmt.Printf("[LinkRepo:GetBySlug] Link with slug (%s) not found.\n", slug)
		return nil, errs.ErrLinkNotFound
	}

	link := new(entity.Link)
	if res := db.DB.Where("slug = ?", slug).Find(link); res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, errs.ErrLinkNotFound
		}

		fmt.Printf("%v\n", res.Error)
		return nil, errs.ErrDbError
	}

	return link, nil
}
