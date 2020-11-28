package utils

import "fmt"

type Client struct {
}

func (c *Client) Notify(i int) {
	fmt.Println("subject status change to ", i)
}
