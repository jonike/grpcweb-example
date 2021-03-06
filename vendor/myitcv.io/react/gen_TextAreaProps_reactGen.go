// Code generated by reactGen. DO NOT EDIT.

package react

// TextAreaProps defines the properties for the <textarea> element
type TextAreaProps struct {
	ClassName               string
	Cols                    uint
	DangerouslySetInnerHTML *DangerousInnerHTML
	DataSet
	DefaultValue string
	ID           string
	Key          string

	OnChange
	OnClick

	Placeholder string
	Ref
	Role  string
	Rows  uint
	Style *CSS
	Value string
}

func (t *TextAreaProps) assign(v *_TextAreaProps) {

	v.ClassName = t.ClassName

	v.Cols = t.Cols

	v.DangerouslySetInnerHTML = t.DangerouslySetInnerHTML

	if t.DataSet != nil {
		for dk, dv := range t.DataSet {
			v.o.Set("data-"+dk, dv)
		}
	}

	if t.DefaultValue != "" {
		v.DefaultValue = t.DefaultValue
	}

	if t.ID != "" {
		v.ID = t.ID
	}

	if t.Key != "" {
		v.Key = t.Key
	}

	if t.OnChange != nil {
		v.o.Set("onChange", t.OnChange.OnChange)
	}

	if t.OnClick != nil {
		v.o.Set("onClick", t.OnClick.OnClick)
	}

	v.Placeholder = t.Placeholder

	if t.Ref != nil {
		v.o.Set("ref", t.Ref.Ref)
	}

	v.Role = t.Role

	v.Rows = t.Rows

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	v.Style = t.Style.hack()

	v.Value = t.Value

}
