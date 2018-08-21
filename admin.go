package phone

import (
	"github.com/aghape/admin"
	"github.com/aghape/admin/admincommon"
)

func AddSubResource(res *admin.Resource, value interface{}, fieldName ...string) *admin.Resource {
	if len(fieldName) == 0 || fieldName[0] == "" {
		fieldName = []string{"Phones"}
	}
	return res.AddResource(&admin.SubConfig{FieldName: fieldName[0]}, value, &admin.Config{Setup: func(r *admin.Resource) {
		r.SetI18nModel(&Phone{})
		PrepareResource(r)
		res.SetMeta(&admin.Meta{Name: fieldName[0], Resource: r})
	}})
}

func PrepareResource(res *admin.Resource) {
	admincommon.RecordInfoFields(res)
	res.SetMeta(&admin.Meta{Name: "CountryCode"}).TemplateData = map[string]interface{}{"attrs": "maxlength=\"2\""}
	res.EditAttrs(&admin.Section{Rows: [][]string{{"CountryCode", "Phone", "Contact"}}})
	res.ShowAttrs(res.EditAttrs())
	res.NewAttrs(res.EditAttrs())
}

func GetResource(Admin *admin.Admin) *admin.Resource {
	return Admin.GetResourceByID("Phone")
}
