package echoapp

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Title string `json:"title"`
}

type Tag struct {
	gorm.Model
	Title string `json:"title"`
}

type TagGroup struct {
	gorm.Model
	TagId   uint `json:"tag_id"`
	GroupId uint `json:"group_id"`
}

func (TagGroup) TableName() string {
	return "tag_group"
}

type Group struct {
	gorm.Model
	Title      string    `gorm:"column:title" json:"title"`
	UserId     uint      `gorm:"column:user_id"json:"user_id"`
	Covers     []string  `gorm:"-" json:"covers"`
	CoversStr  string    `gorm:"column:covers" json:"covers_str"`
	Desc       string    `gorm:"column:desc" json:"desc"`
	CategoryId uint      `gorm:"column:category_id" json:"category_id"`
	TagIds     []uint    `gorm:"-" json:"tags"`
	Display    string    `json:"display"`
	Chapters   []Chapter `json:"chapters"`
}

func (g *Group) BeforeSave() (err error) {
	var str []byte
	if str, err = json.Marshal(g.Covers); err != nil {
		return err
	}
	g.CoversStr = string(str)
	return
}

type Chapter struct {
	gorm.Model
	Title     string     `json:"title"`
	GroupId   uint       `json:"group_id"`
	ParentId  uint       `json:"parent_id"`
	Resources []Resource `json:"resources"`
}

type Resource struct {
	gorm.Model
	Title     string     `json:"title"`
	UserId    uint       `json:"user_id"`
	GroupId   uint       `json:"group_id"`
	ChapterId uint       `json:"chapter_id"`
	Covers    string     `json:"covers"`
	Type      string     `json:"type"`
	Article   *Article   `json:"article" gorm:"ForeignKey:rid"`
	Testpaper *Testpaper `json:"testpaper" gorm:"ForeignKey:rid"`
}

type Article struct {
	gorm.Model
	Rid     int    `json:"rid"`
	Content string `json:"content"`
}

type Testpaper struct {
	gorm.Model
	Rid     int    `json:"rid"`
	Content string `json:"content"`
}

type ResourceService interface {
	SaveResource(resource Resource) error
	GetResourceById(id uint) (*Resource, error)
	//通过tagId查找资源
	GetResourcesByTagID(tagID uint, from, limit int) ([]*Resource, error)
	//用户上传的资源
	GetSelfResources(userId uint, from, limit int) ([]*Resource, error)
	//用户购买的资源
	GetUserPaymentResources(userId uint, from, limit int) ([]*Resource, error)
}
