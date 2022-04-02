package service

type AuthService interface {
	Verify(routeName string, routeVars map[string]string, token string) error
}
