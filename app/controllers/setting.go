
package controllers

import (
  "app/models"
)

type Setting struct {
}

func NewSetting() Setting {
  return Setting{}
}

func (c Setting) GetByTenantId(tenantId int) interface{} {
  repo := models.NewSettingRepository()
  setting, result := repo.Get(tenantId)
  if result {
    return setting
  }
  return result
}

func (c Setting) PostSetting(tenantId int, setting models.Setting) {
  repo := models.NewSettingRepository()
  repo.Post(tenantId, setting)
}

func (c Setting) PutSetting(tenantId int, setting models.Setting) bool {
  repo := models.NewSettingRepository()
  result := repo.Put(tenantId, setting)
  if result {
    return true
  }
  return false
}

func (c Setting) DeleteSetting(tenantId int) bool {
  repo := models.NewSettingRepository()
  result := repo.Delete(tenantId)
  if result == false {
    return false
  }
  return true
}
