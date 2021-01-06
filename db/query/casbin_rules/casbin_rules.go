package casbin_rules

import (
	db "github.com/JIeeiroSst/store/component"
	"github.com/JIeeiroSst/store/models/casbin_rules"
)

func CasbinRulesAll(casbin_rules casbin_rules.Casbin_rules){
	db.GetConn().Find(&casbin_rules)
}

func CasbinRoleById(casbin_rules casbin_rules.Casbin_rules,id int){
	db.GetConn().Find(&casbin_rules,id)
}

func CreateCasbinRules(casbin_rules casbin_rules.Casbin_rules){
	db.GetConn().Create(&casbin_rules)
}

func DeleteCasbinRules(casbin_rules casbin_rules.Casbin_rules,id int){
	db.GetConn().First(&casbin_rules,id)
	db.GetConn().Delete(&casbin_rules)
}

func UpdateCasbinRules(casbin_rules casbin_rules.Casbin_rules,id int){
	db.GetConn().Where("id = ?" ,id).First(&casbin_rules)
	db.GetConn().Save(&casbin_rules)
}
