package phone

import (
	"github.com/ecletus/admin"
	"github.com/ecletus-pkg/admin"
	"github.com/ecletus/db"
	"github.com/ecletus/plug"
)

type Plugin struct {
	plug.EventDispatcher
	db.DBNames
	admin_plugin.AdminNames
}

func (p *Plugin) OnRegister() {
	admin_plugin.Events(p).InitResources(func(e *admin_plugin.AdminEvent) {
		e.Admin.AddResource(&Phone{}, &admin.Config{Setup: PrepareResource, Invisible:true})
	})
	db.Events(p).DBOnMigrate(func(e *db.DBEvent) error {
		return e.AutoMigrate(&Phone{}).Error
	})
}
