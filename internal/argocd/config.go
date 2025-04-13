package argocd

type Config struct {
	App      string
	Server   string
	Username string
	Password string
	Insecure bool
}
