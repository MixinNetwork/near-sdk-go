package rpc

func (c *Client) NodeStatus() (*NodeStatus, error) {
	var res struct {
		GeneralResponse
		Result NodeStatus `json:"result"`
	}
	err := c.request("status", []string{}, &res)
	if err != nil {
		return nil, err
	}
	if res.Error != nil {
		return nil, res.Error
	}
	return &res.Result, nil
}

type NodeStatus struct {
	Version struct {
		Version string `json:"version"`
		Build   string `json:"build"`
	} `json:"version"`
	ChainId string `json:"chain_id"`
}
