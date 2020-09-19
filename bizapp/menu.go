package bizapp

import "fmt"

type Menu interface {
	GetId() string
	GetParentId() string
	GetName() string
	GetLevel() int
	GetChildren() []Menu
	AddChild(id string, name string, options ...MenuOption)
	AddChildren(children ...Menu)
	Each(func(Menu) bool) bool
	setParentId(parentId string)
	setLevel(level int)
	Copy(copyAll bool) Menu // 复制菜单返回一下新的菜单，如果 copyAll 为 true 则复制所有子菜单，否则仅返回菜单本身
}

type menu struct {
	Id          string   `json:"id"`
	ParentId    string   `json:"parentId"`
	Name        string   `json:"name"`
	Subtitle    string   `json:"subtitle"`
	Type        string   `json:"type"`
	Level       int      `json:"level"`
	Icon        string   `json:"icon"`
	AuthCodes   []string `json:"authCodes"`
	Component   string   `json:"component"`
	Path        string   `json:"path"`
	Description string   `json:"description"`
	Status      string   `json:"status"`
	Children    []Menu   `json:"children"`
}

func (m *menu) GetId() string {
	return m.Id
}

func (m *menu) GetParentId() string {
	return m.ParentId
}

func (m *menu) GetName() string {
	return m.Name
}

func (m *menu) GetLevel() int {
	return m.Level
}

func (m *menu) GetChildren() []Menu {
	return m.Children
}

func (m *menu) AddChild(id string, name string, options ...MenuOption) {
	c := NewMenu(id, name, options...)
	c.setParentId(m.GetId())
	c.setLevel(m.GetLevel() + 1)
	m.AddChildren(c)
}

func (m *menu) AddChildren(children ...Menu) {
	for _, c := range children {
		c.setParentId(m.GetId())
		c.setLevel(m.GetLevel() + 1)
	}
	m.Children = append(m.Children, children...)
}

func (m *menu) Each(f func(Menu) bool) bool {
	for _, c := range m.GetChildren() {
		if f(c) == false {
			return false
		}
		if c.Each(f) == false {
			return false
		}
	}
	return true
}

func (m *menu) Copy(copyAll bool) Menu {

	menu := &menu{
		Id:          m.Id,
		ParentId:    m.ParentId,
		Name:        m.Name,
		Subtitle:    m.Subtitle,
		Type:        m.Type,
		Level:       m.Level,
		Icon:        m.Icon,
		AuthCodes:   m.AuthCodes,
		Component:   m.Component,
		Path:        m.Path,
		Description: m.Description,
		Status:      m.Status,
	}

	if copyAll == false {
		return menu
	}

	m.Each(func(m Menu) bool {
		new := m.Copy(copyAll)
		menu.AddChildren(new)
		return true
	})

	return menu
}

func (m *menu) setParentId(parentId string) {
	m.ParentId = parentId
}

func (m *menu) setLevel(level int) {
	m.Level = level
}

func NewMenu(id string, name string, options ...MenuOption) Menu {

	if id == "" {
		panic("menu id cannot be empty")
	}

	m := &menu{
		Id:   id,
		Name: name,
	}

	for _, option := range options {
		option(m)
	}

	if m.Id == m.ParentId {
		panic(fmt.Sprintf("menu id '%s' cannot equal parent id", id))
	}

	return m
}

func RootMenu() Menu {
	m := &menu{
		Id:       "",
		Name:     "ROOT",
		ParentId: "",
	}
	return m
}
