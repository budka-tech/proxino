package proxy_market

type ListResponse struct {
	Success bool   `json:"success"`
	Balance string `json:"balance"`
	List    struct {
		Error    bool `json:"error"`
		Total    int  `json:"total"`
		PageSize int  `json:"page_size"`
		Data     []struct {
			Id         int    `json:"id"`
			Ip         string `json:"ip"`
			IpOut      string `json:"ip_out"`
			HttpPort   string `json:"http_port"`
			SocksPort  string `json:"socks_port"`
			Login      string `json:"login"`
			Password   string `json:"password"`
			Comment    string `json:"comment"`
			ExpiresAt  string `json:"expires_at"`
			IsHidden   string `json:"is_hidden"`
			AttachedIp string `json:"attached_ip"`
			BoughtAt   string `json:"bought_at"`
			Type       string `json:"type"`
			ProxyType  string `json:"proxy_type"`
		} `json:"data"`
	} `json:"list"`
}
