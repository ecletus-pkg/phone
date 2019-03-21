package phone

import (
	"github.com/aghape/admin"
)

func AddSubResource(res *admin.Resource, value interface{}, fieldName ...string) *admin.Resource {
	cfg := &admin.Config{Setup: func(r *admin.Resource) {
		r.SetI18nModel(&Phone{})
		PrepareResource(r)
		res.SetMeta(&admin.Meta{Name: fieldName[0], Resource: r})
	}}
	if len(fieldName) == 0 || fieldName[0] == "" {
		fieldName = []string{"Phones"}
		res.Meta(&admin.Meta{
			Name:  fieldName[0],
			Label: GetResource(res.GetAdmin()).PluralLabelKey(),
		})
	} else {
		cfg.LabelKey = res.ChildrenLabelKey(fieldName[0])
	}
	return res.AddResource(&admin.SubConfig{FieldName: fieldName[0]}, value, cfg)
}

func PrepareResource(res *admin.Resource) {
	res.SetMeta(&admin.Meta{Name: "CountryCode"}).TemplateData = map[string]interface{}{"attrs": "maxlength=\"2\""}
	res.EditAttrs(&admin.Section{Rows: [][]string{{"CountryCode", "Number", "Note"}}})
	res.ShowAttrs(res.EditAttrs())
	res.NewAttrs(res.EditAttrs())
	res.IndexAttrs("CountryCode", "Number", "Note")
}

func GetResource(Admin *admin.Admin) *admin.Resource {
	return Admin.GetResourceByID("Phone")
}
