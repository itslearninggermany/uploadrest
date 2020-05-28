package uploadrest

import "github.com/jinzhu/gorm"

type UniventionRestSetup struct {
	gorm.Model
	FirstRun       bool   `gorm:"-" json:"first_run"`
	Uuid           string `gorm:"unique" json:"uuid"`
	OrganisationID string `gorm:"unique" json:"-"`
	InstitutionID  string `json:"-"`
	AESkey         string `gorm:"VARCHAR 500" json:"aeskey"`
	AuthKey        string `gorm:"VARCHAR 500" json:"auth_key"`
}

func GetUniventionRestSetup(uuid string, db *gorm.DB) (setup UniventionRestSetup, err error) {
	err = db.Where("uuid = ?", uuid).Last(&setup).Error
	return setup, err
}

func (p *UniventionRestSetup) Initial(AESKey string, authKey string) {
	p.AESkey = AESKey
	p.AuthKey = authKey
}
