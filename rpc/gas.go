package rpc

func (c *Client) GetGasPrice(blockId string) (string, error) {
	var res struct {
		GeneralResponse
		Result struct {
			GasPrice string `json:"gas_price"`
		} `json:"result"`
	}
	err := c.request("gas_price", []string{blockId}, &res)
	if err != nil {
		return "", err
	}
	if res.Error != nil {
		return "", res.Error
	}
	return res.Result.GasPrice, nil
}
