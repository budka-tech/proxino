package proxy_market

import (
	"log"
	"net/http"
	"net/url"
)

type ProxyMarket struct {
	client *http.Client
}

func NewProxyMarket(proxyUrl string) *ProxyMarket {
	proxy, err := url.Parse(proxyUrl)
	if err != nil {
		log.Printf("Ошибка при парсинге прокси URL %s: %v", proxyUrl, err)
	}

	// Создание пользовательского транспорта
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxy),
	}

	client := &http.Client{
		Transport: transport,
	}
	return &ProxyMarket{
		client: client,
	}
}

func (receiver *ProxyMarket) Random() *http.Client {
	return receiver.client
}
