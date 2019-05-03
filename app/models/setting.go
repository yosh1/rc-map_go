package models

import (
	"app/common"

	"github.com/jinzhu/gorm"
)

type Setting struct {
	TenantID int    `json:"tenant_id" binding:"required"`
	Tos      string `json:"tos" binding:"required"`
}

type SettingRepository struct {
}

func NewSettingRepository() SettingRepository {
	return SettingRepository{}
}

func (m SettingRepository) Get(tenantId int) (*Setting, bool) {
	db := common.GetDB()
	var setting Setting
	if err := db.Where(Setting{TenantId: tenantId}).Find(&setting).Error; gorm.IsRecordNotFoundError(err) {
		return &setting, false
	}
	return &setting, true
}

func (m SettingRepository) Post(tenantId int, setting Setting) bool {
	setting.TenantId = tenantId
	db := common.GetDB()
	if err := db.Create(setting).Error; err != nil {
		return false
	}
	return true
}

func (m SettingRepository) Put(tenantId int, setting Setting) bool {
	db := common.GetDB()
	if err := db.Where(Setting{TenantId: tenantId}).Find(&setting).Error; gorm.IsRecordNotFoundError(err) {
		return false
	}
	db.Model(&setting).Where("tenant_id=?", tenantId).Update(setting)
	return true
}

func (m SettingRepository) Delete(tenantId int) bool {
	db := common.GetDB()
	var setting Setting
	if err := db.Where(Setting{TenantId: tenantId}).Find(&setting).Error; gorm.IsRecordNotFoundError(err) {
		return false
	}
	db.Where("tenant_id=?", tenantId).Delete(&setting)
	return true
}
