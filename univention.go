package uploadrest

import "github.com/jinzhu/gorm"

type univentionRestSetup struct {
	gorm.Model
	Uuid           string `gorm:"unique"`
	OrganisationID string `gorm:"unique"`
	InstitutionID  string
	AESkey         string `gorm:"VARCHAR 500"`
	AuthKey        string `gorm:"VARCHAR 500"`
}

func GetUniventionRestSetup(uuid string, db *gorm.DB) (setup *univentionRestSetup, err error) {
	err = db.Where("uuid = ?", uuid).Last(setup).Error
	return setup, err
}

func (p *univentionRestSetup) Initial(AESKey string, authKey string) {
	p.AESkey = AESKey
	p.AuthKey = authKey
}
