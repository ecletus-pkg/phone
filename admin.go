package phone

import (
	"github.com/ecletus/admin"
)

func AddSubResource(setup func(res *admin.Resource), res *admin.Resource, value interface{}, fieldName ...string) error {
	return res.GetAdmin().OnResourcesAdded(func(e *admin.ResourceEvent) error {
		cfg := &admin.Config{Setup: func(r *admin.Resource) {
			r.SetI18nModel(&Phone{})
			PrepareResource(r)
			res.SetMeta(&admin.Meta{Name: fieldName[0], Resource: r})
			if setup != nil {
				setup(r)
			}
		}}
		if len(fieldName) == 0 || fieldName[0] == "" {
			fieldName = []string{"Phones"}
			res.Meta(&admin.Meta{
				Name:  fieldName[0],
				Label: e.Resource.PluralLabelKey(),
			})
		} else {
			cfg.LabelKey = res.ChildrenLabelKey(fieldName[0])
		}
		res.AddResource(&admin.SubConfig{FieldName: fieldName[0]}, value, cfg)
		return nil
	}, ResourceID)
}

func PrepareResource(res *admin.Resource) {
	res.SetMeta(&admin.Meta{Name: "CountryCode"}).TemplateData = map[string]interface{}{"attrs": "maxlength=\"2\""}
	res.EditAttrs(&admin.Section{Rows: [][]string{{"CountryCode", "Number", "Note"}}})
	res.ShowAttrs(res.EditAttrs())
	res.NewAttrs(res.EditAttrs())
	res.IndexAttrs("CountryCode", "Number", "Note")
}

const ResourceID = "Phone"
