package proxino

import (
	"log"
	"net/http"
	"net/url"
)

func NewClient(proxy *Proxy) *http.Client {
	proxyUrl, err := url.Parse(proxy.Http())
	if err != nil {
		log.Printf("Ошибка при парсинге прокси URL %s: %v", proxy, err)
	}

	// Создание пользовательского транспорта
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}

	client := &http.Client{
		Transport: transport,
	}

	return client
}
