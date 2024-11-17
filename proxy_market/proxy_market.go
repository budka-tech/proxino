package proxy_market

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/budka-tech/proxino"
	"io"
	"math/rand"
	"net/http"
	"time"
)

type ProxyMarket struct {
	balance float32
	token   string
	pool    map[string]*proxino.Proxy
	List    []*http.Client
}

func NewProxyMarket(token string) *ProxyMarket {
	return &ProxyMarket{
		balance: 0,
		token:   token,
		pool:    make(map[string]*proxino.Proxy),
		List:    make([]*http.Client, 0, 10),
	}
}

func (receiver *ProxyMarket) GetList() {
	params := map[string]string{
		"type":       "all",
		"proxy_type": "",
		"page":       "1",
		"page_size":  "100000",
	}

	jsonParams, err := json.Marshal(params)
	if err != nil {
		fmt.Println("Ошибка при кодировании JSON:", err)
		return
	}

	resp, err := http.Post(receiver.Path("list"), "application/json", bytes.NewBuffer(jsonParams))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}

	var list ListResponse
	if err := json.Unmarshal(responseData, &list); err != nil {
		fmt.Println("Ошибка при декодировании ответа:", err)
		return
	}

	for _, item := range list.List.Data {
		proxy := &proxino.Proxy{
			Ip:       item.Ip,
			Port:     item.HttpPort,
			Username: item.Login,
			Password: item.Password,
		}
		url := proxy.Url()

		if _, ok := receiver.pool[url]; !ok {
			cli := proxino.NewClient(proxy)
			receiver.pool[url] = proxy
			receiver.List = append(receiver.List, cli)
		}
	}
}

func (receiver *ProxyMarket) Path(path string) string {
	return fmt.Sprintf("https://api.dashboard.proxy.market/dev-api/%s/%s", path, receiver.token)
}

func (receiver *ProxyMarket) Random() *http.Client {
	if len(receiver.List) == 0 {
		return nil
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(receiver.List))
	return receiver.List[randomIndex]
}
