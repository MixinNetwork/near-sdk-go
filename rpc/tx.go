package rpc

func (c *Client) GetTx(transactionHash, senderId string) (*TransactionResult, error) {
	var res struct {
		GeneralResponse
		Result TransactionResult `json:"result"`
	}
	err := c.request("tx", []string{transactionHash, senderId}, &res)
	if err != nil {
		return nil, err
	}
	if res.Error != nil {
		return nil, res.Error
	}
	return &res.Result, nil
}
func (c *Client) BroadcastTxCommit(raw string) (string, error) {
	var res struct {
		GeneralResponse
		Result string `json:"result"`
	}
	err := c.request("broadcast_tx_commit", []string{raw}, &res)
	if err != nil {
		return "", err
	}
	if res.Error != nil {
		return "", res.Error
	}
	return res.Result, nil
}

func (c *Client) BroadcastTxAsync(raw string) (string, error) {
	var res struct {
		GeneralResponse
		Result string `json:"result"`
	}
	err := c.request("broadcast_tx_async", []string{raw}, &res)
	if err != nil {
		return "", err
	}
	if res.Error != nil {
		return "", res.Error
	}
	return res.Result, nil
}

type TransactionResult struct {
	ReceiptsOutcome []struct {
		BlockHash string `json:"block_hash"`
		ID        string `json:"id"`
		Outcome   struct {
			ExecutorID string        `json:"executor_id"`
			GasBurnt   int64         `json:"gas_burnt"`
			Logs       []interface{} `json:"logs"`
			ReceiptIds []string      `json:"receipt_ids"`
			Status     struct {
				SuccessValue string `json:"SuccessValue"`
			} `json:"status"`
			TokensBurnt string `json:"tokens_burnt"`
		} `json:"outcome"`
		Proof []struct {
			Direction string `json:"direction"`
			Hash      string `json:"hash"`
		} `json:"proof"`
	} `json:"receipts_outcome"`
	Status struct {
		SuccessValue     string `json:"SuccessValue,omitempty"`
		SuccessReceiptId string `json:"SuccessReceiptId,omitempty"`
		Failure          string `json:"Failure,omitempty"`
		Unknown          string `json:"Unknown,omitempty"`
	} `json:"status"`
	Transaction struct {
		Actions []struct {
			Transfer struct {
				Deposit string `json:"deposit"`
			} `json:"Transfer"`
		} `json:"actions"`
		Hash       string `json:"hash"`
		Nonce      int    `json:"nonce"`
		PublicKey  string `json:"public_key"`
		ReceiverID string `json:"receiver_id"`
		Signature  string `json:"signature"`
		SignerID   string `json:"signer_id"`
	} `json:"transaction"`
	TransactionOutcome struct {
		BlockHash string `json:"block_hash"`
		ID        string `json:"id"`
		Outcome   struct {
			ExecutorID string        `json:"executor_id"`
			GasBurnt   int64         `json:"gas_burnt"`
			Logs       []interface{} `json:"logs"`
			ReceiptIds []string      `json:"receipt_ids"`
			Status     struct {
				SuccessReceiptID string `json:"SuccessReceiptId"`
			} `json:"status"`
			TokensBurnt string `json:"tokens_burnt"`
		} `json:"outcome"`
		Proof []struct {
			Direction string `json:"direction"`
			Hash      string `json:"hash"`
		} `json:"proof"`
	} `json:"transaction_outcome"`
}
