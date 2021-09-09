package rpc

func (c *Client) GetChunk(chunkHash string) (*Chunk, error) {
	var res struct {
		GeneralResponse
		Result Chunk `json:"result"`
	}
	err := c.request("chunk", map[string]interface{}{"chunk_id": chunkHash}, &res)
	if err != nil {
		return nil, err
	}
	if res.Error != nil {
		return nil, res.Error
	}
	return &res.Result, nil
}

type Chunk struct {
	Author string `json:"author"`
	Header struct {
		BalanceBurnt         string        `json:"balance_burnt"`
		ChunkHash            string        `json:"chunk_hash"`
		EncodedLength        int           `json:"encoded_length"`
		EncodedMerkleRoot    string        `json:"encoded_merkle_root"`
		GasLimit             int64         `json:"gas_limit"`
		GasUsed              int64         `json:"gas_used"`
		HeightCreated        int64         `json:"height_created"`
		HeightIncluded       int64         `json:"height_included"`
		OutcomeRoot          string        `json:"outcome_root"`
		OutgoingReceiptsRoot string        `json:"outgoing_receipts_root"`
		PrevBlockHash        string        `json:"prev_block_hash"`
		PrevStateRoot        string        `json:"prev_state_root"`
		RentPaid             string        `json:"rent_paid"`
		ShardID              int           `json:"shard_id"`
		Signature            string        `json:"signature"`
		TxRoot               string        `json:"tx_root"`
		ValidatorProposals   []interface{} `json:"validator_proposals"`
		ValidatorReward      string        `json:"validator_reward"`
	} `json:"header"`
	Receipts     []interface{} `json:"receipts"`
	Transactions []struct {
		Actions []struct {
			Transfer *struct {
				Deposit string `json:"deposit"`
			} `json:"Transfer,omitempty"`
		} `json:"actions"`
		Hash       string `json:"hash"`
		Nonce      int    `json:"nonce"`
		PublicKey  string `json:"public_key"`
		ReceiverID string `json:"receiver_id"`
		Signature  string `json:"signature"`
		SignerID   string `json:"signer_id"`
	} `json:"transactions"`
}
