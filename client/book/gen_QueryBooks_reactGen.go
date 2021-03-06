// Code generated by reactGen. DO NOT EDIT.

package book

import "myitcv.io/react"

type QueryBooksElem struct {
	react.Element
}

func buildQueryBooks(cd react.ComponentDef) react.Component {
	return QueryBooksDef{ComponentDef: cd}
}

func buildQueryBooksElem(props QueryBooksProps, children ...react.Element) *QueryBooksElem {
	return &QueryBooksElem{
		Element: react.CreateElement(buildQueryBooks, props, children...),
	}
}

func (q QueryBooksDef) RendersElement() react.Element {
	return q.Render()
}

// SetState is an auto-generated proxy proxy to update the state for the
// QueryBooks component.  SetState does not immediately mutate q.State()
// but creates a pending state transition.
func (q QueryBooksDef) SetState(state QueryBooksState) {
	q.ComponentDef.SetState(state)
}

// State is an auto-generated proxy to return the current state in use for the
// render of the QueryBooks component
func (q QueryBooksDef) State() QueryBooksState {
	return q.ComponentDef.State().(QueryBooksState)
}

// IsState is an auto-generated definition so that QueryBooksState implements
// the myitcv.io/react.State interface.
func (q QueryBooksState) IsState() {}

var _ react.State = QueryBooksState{}

// GetInitialStateIntf is an auto-generated proxy to GetInitialState
func (q QueryBooksDef) GetInitialStateIntf() react.State {
	return QueryBooksState{}
}

func (q QueryBooksState) EqualsIntf(val react.State) bool {
	return q == val.(QueryBooksState)
}

// IsProps is an auto-generated definition so that QueryBooksProps implements
// the myitcv.io/react.Props interface.
func (q QueryBooksProps) IsProps() {}

// Props is an auto-generated proxy to the current props of QueryBooks
func (q QueryBooksDef) Props() QueryBooksProps {
	uprops := q.ComponentDef.Props()
	return uprops.(QueryBooksProps)
}

func (q QueryBooksProps) EqualsIntf(val react.Props) bool {
	return q == val.(QueryBooksProps)
}

var _ react.Props = QueryBooksProps{}
