// Code generated by reactGen. DO NOT EDIT.

package react

// H3Props defines the properties for the <h3> element
type H3Props struct {
	ClassName               string
	DangerouslySetInnerHTML *DangerousInnerHTMLDef
	ID                      string
	Key                     string

	OnChange
	OnClick

	Role  string
	Style *CSS
}

func (h *H3Props) assign(v *_H3Props) {

	v.ClassName = h.ClassName

	v.DangerouslySetInnerHTML = h.DangerouslySetInnerHTML

	if h.ID != "" {
		v.ID = h.ID
	}

	if h.Key != "" {
		v.Key = h.Key
	}

	if h.OnChange != nil {
		v.o.Set("onChange", h.OnChange.OnChange)
	}

	if h.OnClick != nil {
		v.o.Set("onClick", h.OnClick.OnClick)
	}

	v.Role = h.Role

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	v.Style = h.Style.hack()

}