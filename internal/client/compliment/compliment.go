package compliment

import "net/http"

type Client struct {
	host string
	//basePath string
	http.Client
}

func New(host string) *Client {
	return &Client{
		host: host,
		//client: http.Client{},
	}
}

//func (c *Client) GetCompliment() {
//
//}
