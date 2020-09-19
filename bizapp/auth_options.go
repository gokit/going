package bizapp

type AuthOption func(auth *auth)

type authOptions struct{}

func AuthOptions() authOptions {
	return authOptions{}
}

func (authOptions) ParentId(parentId string) AuthOption {
	return func(auth *auth) {
		auth.ParentId = parentId
	}
}

func (authOptions) Description(description string) AuthOption {
	return func(auth *auth) {
		auth.Description = description
	}
}

func (authOptions) Status(status string) AuthOption {
	return func(auth *auth) {
		auth.Status = status
	}
}
func (authOptions) Children(children []Auth) AuthOption {
	return func(auth *auth) {
		auth.AddChildren(children...)
	}
}
