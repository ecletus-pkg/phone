package phone

import (
	"github.com/moisespsena-go/aorm"
)

type Phone struct {
	aorm.AuditedModel
	CountryCode string `gorm:"size:2"`
	Phone       string `gorm:"size:255"`
	Note        string `gorm:"size:255"`
}

func (p *Phone) String() (s string) {
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
