package phone

import (
	"github.com/aghape/db/common"
	"github.com/aghape/db/common/mixins"
)

type QorPhone struct {
	common.Model
	mixins.CreationUpdationMixin
	CountryCode string `gorm:"size:2"`
	Phone       string `gorm:"size:255"`
	Note        string `gorm:"size:255"`
}

func (p *QorPhone) String() (s string) {
	if p.CountryCode != "" {
		s += "+"
		s += p.CountryCode
	}
	s += p.Phone
	if p.Note != "" {
		s += " [" + p.Note + "]"
	}
	return
}
