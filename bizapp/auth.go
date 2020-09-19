package bizapp

import "fmt"

type Auth interface {
	GetId() string
	GetParentId() string
	GetName() string
	GetLevel() int
	GetChildren() []Auth
	AddChildren(children ...Auth)
	Each(func(Auth) bool) bool
	setParentId(parentId string)
	setLevel(level int)
}

type auth struct {
	Id          string `json:"id"`
	ParentId    string `json:"parentId"`
	Name        string `json:"name"`
	Level       int    `json:"level"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Children    []Auth `json:"children"`
}

func (a *auth) GetId() string {
	return a.Id
}

func (a *auth) GetParentId() string {
	return a.ParentId
}

func (a *auth) GetName() string {
	return a.Name
}

func (a *auth) GetLevel() int {
	return a.Level
}

func (a *auth) GetChildren() []Auth {
	return a.Children
}

func (a *auth) AddChild(id string, name string, options ...AuthOption) {
	c := NewAuth(id, name, options...)
	c.setParentId(a.GetId())
	c.setLevel(a.GetLevel() + 1)
	a.AddChildren(c)
}

func (a *auth) AddChildren(children ...Auth) {
	for _, c := range children {
		c.setParentId(a.GetId())
		c.setLevel(a.GetLevel() + 1)
	}
	a.Children = append(a.Children, children...)
}

func (a *auth) Each(f func(Auth) bool) bool {
	for _, c := range a.GetChildren() {
		if f(c) == false {
			return false
		}
		if c.Each(f) == false {
			return false
		}
	}
	return true
}

func (a *auth) setParentId(parentId string) {
	a.ParentId = parentId
}

func (a *auth) setLevel(level int) {
	a.Level = level
}

func NewAuth(id string, name string, options ...AuthOption) Auth {

	if id == "" {
		panic("auth id cannot be empty")
	}

	a := &auth{
		Id:   id,
		Name: name,
	}

	for _, option := range options {
		option(a)
	}

	if a.Id == a.ParentId {
		panic(fmt.Sprintf("auth id '%s' cannot equal parent id", id))
	}

	return a
}

func RootAuth() Auth {
	a := &auth{
		Id:       "",
		Name:     "ROOT",
		ParentId: "",
	}
	return a
}
