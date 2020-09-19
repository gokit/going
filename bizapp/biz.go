package bizapp

import "fmt"

// 业务接口
type Biz interface {
	GetId() string
	GetName() string
	GetMenus() []Menu
	GetAuths() []Auth
	AddMenu(id string, name string, options ...MenuOption)
	AddMenus(menus ...Menu)
	AddAuth(id string, name string, options ...AuthOption)
	AddAuths(auths ...Auth)
	HasMenu(id string) bool
	HasAuth(id string) bool
	GetMenu(id string) Menu
	GetAuth(id string) Auth
	EachMenu(f func(Menu) bool)
	EachAuth(f func(Auth) bool)
	Build()
}

type MenuAction func(parent Menu)
type AuthAction func(parent Auth)

type biz struct {
	Id          string
	Name        string
	Menus       []Menu
	Auths       []Auth
	menuActions map[string][]MenuAction
	authActions map[string][]AuthAction
	menuMap     map[string]Menu
	authMap     map[string]Auth
}

func (b *biz) EachMenu(f func(Menu) bool) {
	for _, c := range b.GetMenus() {
		if f(c) == false {
			break
		}

		if c.Each(f) == false {
			break
		}
	}
}

func (b *biz) EachAuth(f func(Auth) bool) {
	for _, c := range b.GetAuths() {
		if f(c) == false {
			break
		}

		if c.Each(f) == false {
			break
		}
	}
}

func NewBiz(id string, name string) Biz {
	b := &biz{
		Id:   id,
		Name: name,
	}

	b.menuActions = make(map[string][]MenuAction)
	b.authActions = make(map[string][]AuthAction)
	b.menuMap = make(map[string]Menu)
	b.authMap = make(map[string]Auth)

	return b
}

func (b *biz) GetId() string {
	return b.Id
}

func (b *biz) GetName() string {
	return b.Name
}

func (b *biz) GetMenus() []Menu {
	return b.Menus
}

func (b *biz) GetAuths() []Auth {
	return b.Auths
}

func (b *biz) AddMenu(id string, name string, options ...MenuOption) {
	b.AddMenus(NewMenu(id, name, options...))
}

func (b *biz) AddMenus(menus ...Menu) {
	for _, m := range menus {

		if m.GetId() == "" {
			panic("menu id cannot be empty")
		}

		if m.GetId() == m.GetParentId() {
			panic(fmt.Sprintf("menu id '%s' cannot equal parent id", m.GetId()))
		}

		if _, ok := b.authMap[m.GetId()]; ok {
			panic(fmt.Sprintf("duplicate menu id '%s'", m.GetId()))
		}

		b.menuActions[m.GetParentId()] = append(b.menuActions[m.GetParentId()], func(parent Menu) {
			parent.AddChildren(m)
			// 递归添加子菜单
			for _, action := range b.menuActions[m.GetId()] {
				action(m)
			}
		})
	}
}

func (b *biz) AddAuth(id string, name string, options ...AuthOption) {
	b.AddAuths(NewAuth(id, name, options...))
}

func (b *biz) AddAuths(auths ...Auth) {
	for _, a := range auths {

		if a.GetId() == a.GetParentId() {
			panic(fmt.Sprintf("auth id '%s' cannot equal parent id", a.GetId()))
		}

		if _, ok := b.authMap[a.GetId()]; ok {
			panic(fmt.Sprintf("duplicate auth id '%s'", a.GetId()))
		}

		b.authActions[a.GetParentId()] = append(b.authActions[a.GetParentId()], func(parent Auth) {
			parent.AddChildren(a)
			// 递归添加子权限
			for _, action := range b.authActions[a.GetId()] {
				action(a)
			}
		})

	}
}

func (b *biz) reset() {
	b.menuMap = make(map[string]Menu)
	b.authMap = make(map[string]Auth)
	b.Menus = nil
	b.Auths = nil
}

func (b *biz) Build() {

	// 重置
	b.reset()

	menuRoot := RootMenu()

	for _, action := range b.menuActions[menuRoot.GetId()] {
		action(menuRoot)
	}

	b.Menus = menuRoot.GetChildren()

	b.EachMenu(func(m Menu) bool {
		b.menuMap[m.GetId()] = m
		return true
	})

	authRoot := RootAuth()

	for _, action := range b.authActions[authRoot.GetId()] {
		action(authRoot)
	}

	b.Auths = authRoot.GetChildren()

	b.EachAuth(func(a Auth) bool {
		b.authMap[a.GetId()] = a
		return true
	})
}

func (b *biz) HasMenu(id string) bool {
	if _, ok := b.menuMap[id]; ok {
		return true
	}
	return false
}

func (b *biz) HasAuth(id string) bool {
	if _, ok := b.authMap[id]; ok {
		return true
	}
	return false
}

func (b *biz) GetMenu(id string) Menu {
	if m, ok := b.menuMap[id]; ok {
		return m
	}
	return nil
}

func (b *biz) GetAuth(id string) Auth {
	if a, ok := b.authMap[id]; ok {
		return a
	}
	return nil
}
