package phone

import (
	"github.com/aghape/admin"
	"github.com/aghape/admin/adminplugin"
	"github.com/aghape/db"
	"github.com/aghape/plug"
)

type Plugin struct {
	plug.EventDispatcher
	db.DBNames
	adminplugin.AdminNames
}

func (p *Plugin) OnRegister() {
	p.AdminNames.OnInitResources(p, func(e *adminplugin.AdminEvent) {
		e.Admin.AddResource(&QorPhone{}, &admin.Config{Setup: PrepareResource})
	})
	db.Events(p).DBOnMigrateGorm(func(e *db.GormDBEvent) error {
		return e.DB.AutoMigrate(&QorPhone{}).Error
	})
}
