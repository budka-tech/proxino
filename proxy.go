package proxino

import "fmt"

type Proxy struct {
	Ip       string
	Port     string
	Username string
	Password string
}

func (p *Proxy) Url() string {
	url := fmt.Sprintf("%s:%s", p.Ip, p.Port)

	if p.Username != "" && p.Password != "" {
		url = fmt.Sprintf("%v:%v@%v", p.Username, p.Password, url)
	}

	return url
}
func (p *Proxy) Http() string {
	return fmt.Sprintf("http://%s", p.Url())
}
