// Code generated by immutableGen. DO NOT EDIT.

package book

//immutableVet:skipFile

import (
	"myitcv.io/immutable"
)

// _Imm_Messages is generated to an immutable
// type *books which we use in the state
//
// Messages is an immutable type and has the following template:
//
// 	[]string
//
type Messages struct {
	theSlice []string
	mutable  bool
	__tmpl   _Imm_Messages
}

var _ immutable.Immutable = new(Messages)
var _ = new(Messages).__tmpl

func NewMessages(s ...string) *Messages {
	c := make([]string, len(s))
	copy(c, s)

	return &Messages{
		theSlice: c,
	}
}

func NewMessagesLen(l int) *Messages {
	c := make([]string, l)

	return &Messages{
		theSlice: c,
	}
}

func (m *Messages) Mutable() bool {
	return m.mutable
}

func (m *Messages) Len() int {
	if m == nil {
		return 0
	}

	return len(m.theSlice)
}

func (m *Messages) Get(i int) string {
	return m.theSlice[i]
}

func (m *Messages) AsMutable() *Messages {
	if m == nil {
		return nil
	}

	if m.Mutable() {
		return m
	}

	res := m.dup()
	res.mutable = true

	return res
}

func (m *Messages) dup() *Messages {
	resSlice := make([]string, len(m.theSlice))

	for i := range m.theSlice {
		resSlice[i] = m.theSlice[i]
	}

	res := &Messages{
		theSlice: resSlice,
	}

	return res
}

func (m *Messages) AsImmutable(v *Messages) *Messages {
	if m == nil {
		return nil
	}

	if v == m {
		return m
	}

	m.mutable = false
	return m
}

func (m *Messages) Range() []string {
	if m == nil {
		return nil
	}

	return m.theSlice
}

func (m *Messages) WithMutable(f func(mi *Messages)) *Messages {
	res := m.AsMutable()
	f(res)
	res = res.AsImmutable(m)

	return res
}

func (m *Messages) WithImmutable(f func(mi *Messages)) *Messages {
	prev := m.mutable
	m.mutable = false
	f(m)
	m.mutable = prev

	return m
}

func (m *Messages) Set(i int, v string) *Messages {
	if m.mutable {
		m.theSlice[i] = v
		return m
	}

	res := m.dup()
	res.theSlice[i] = v

	return res
}

func (m *Messages) Append(v ...string) *Messages {
	if m.mutable {
		m.theSlice = append(m.theSlice, v...)
		return m
	}

	res := m.dup()
	res.theSlice = append(res.theSlice, v...)

	return res
}
func (s *Messages) IsDeeplyNonMutable(seen map[interface{}]bool) bool {
	if s == nil {
		return true
	}

	if s.Mutable() {
		return false
	}
	return true
}
