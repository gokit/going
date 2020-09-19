package bizapp

type MenuOption func(menu *menu)

type menuOptions struct{}

func MenuOptions() menuOptions {
	return menuOptions{}
}

func (menuOptions) ParentId(parentId string) MenuOption {
	return func(menu *menu) {
		menu.ParentId = parentId
	}
}

func (menuOptions) Subtitle(subtitle string) MenuOption {
	return func(menu *menu) {
		menu.Subtitle = subtitle
	}
}

func (menuOptions) Type(typ string) MenuOption {
	return func(menu *menu) {
		menu.Type = typ
	}
}

func (menuOptions) Icon(icon string) MenuOption {
	return func(menu *menu) {
		menu.Icon = icon
	}
}

func (menuOptions) AuthCodes(authCodes ...string) MenuOption {
	return func(menu *menu) {
		menu.AuthCodes = authCodes
	}
}

func (menuOptions) Component(component string) MenuOption {
	return func(menu *menu) {
		menu.Component = component
	}
}

func (menuOptions) Path(path string) MenuOption {
	return func(menu *menu) {
		menu.Path = path
	}
}

func (menuOptions) Description(description string) MenuOption {
	return func(menu *menu) {
		menu.Description = description
	}
}

func (menuOptions) Status(status string) MenuOption {
	return func(menu *menu) {
		menu.Status = status
	}
}

func (menuOptions) Children(children []Menu) MenuOption {
	return func(menu *menu) {
		menu.AddChildren(children...)
	}
}
