package flipside

import (
	"fmt"
	"strings"
)

func (c *Client) GetFirstBuyers(tokenAddress string, limit int) ([]SwapData, error) {
	var sb strings.Builder
	sb.WriteString(`
		WITH swaps AS (
			SELECT
				swap_program,
				block_id,
				block_timestamp,
				tx_id,
				program_id,
				swapper,
				swap_from_mint,
				swap_from_amount,
				swap_to_mint,
				swap_to_amount,
				_log_id,
				ez_swaps_id,
				inserted_timestamp,
				modified_timestamp,
				ROW_NUMBER() OVER (PARTITION BY swapper ORDER BY block_timestamp ASC) as rn
			FROM
				solana.defi.ez_dex_swaps
			WHERE
				swap_to_mint = '`)
	sb.WriteString(strings.ReplaceAll(tokenAddress, "'", "''"))
	sb.WriteString(`'
		)
		SELECT
			swap_program,
			block_id,
			block_timestamp,
			tx_id,
			program_id,
			swapper,
			swap_from_mint,
			swap_from_amount,
			swap_to_mint,
			swap_to_amount,
			_log_id,
			ez_swaps_id,
			inserted_timestamp,
			modified_timestamp
		FROM
			swaps
		WHERE
			rn = 1
		ORDER BY
			block_timestamp ASC
		LIMIT `)
	sb.WriteString(fmt.Sprintf("%d", limit))

	sql := sb.String()

	queryRun, err := c.createQueryRun(sql)
	if err != nil {
		return nil, err
	}

	results, err := c.waitForQueryResults(queryRun.Result.QueryRun.ID)
	if err != nil {
		return nil, err
	}

	var swapDataList []SwapData
	for _, record := range results.Result.Rows {
		data, ok := record.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("unexpected data type in result row")
		}

		blockID, err := parseFloat64(data["block_id"])
		if err != nil {
			return nil, fmt.Errorf("error parsing block_id: %w", err)
		}

		swapFromAmount, err := parseFloat64(data["swap_from_amount"])
		if err != nil {
			return nil, fmt.Errorf("error parsing swap_from_amount: %w", err)
		}

		swapToAmount, err := parseFloat64(data["swap_to_amount"])
		if err != nil {
			return nil, fmt.Errorf("error parsing swap_to_amount: %w", err)
		}

		swapData := SwapData{
			SwapProgram:       data["swap_program"].(string),
			BlockID:           int64(blockID),
			BlockTimestamp:    parseTimestamp(data["block_timestamp"].(string)),
			TxID:              data["tx_id"].(string),
			ProgramID:         data["program_id"].(string),
			Swapper:           data["swapper"].(string),
			SwapFromMint:      data["swap_from_mint"].(string),
			SwapFromAmount:    swapFromAmount,
			SwapToMint:        data["swap_to_mint"].(string),
			SwapToAmount:      swapToAmount,
			LogID:             data["_log_id"].(string),
			EzSwapsID:         data["ez_swaps_id"].(string),
			InsertedTimestamp: parseTimestamp(data["inserted_timestamp"].(string)),
			ModifiedTimestamp: parseTimestamp(data["modified_timestamp"].(string)),
		}

		swapDataList = append(swapDataList, swapData)
	}

	return swapDataList, nil
}

func (c *Client) GetTransfersBetweenAddresses(addresses []string, limit int) (AddressInteractionsMap, error) {
	var sb strings.Builder
	sb.WriteString(`
		SELECT
			block_timestamp,
			block_id,
			tx_id,
			tx_from,
			tx_to,
			amount,
			mint,
			fact_transfers_id,
			inserted_timestamp,
			modified_timestamp
		FROM
			solana.core.fact_transfers
		WHERE
			(tx_from IN (`)

	// Add placeholders for addresses
	for i, addr := range addresses {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString("'")
		sb.WriteString(strings.ReplaceAll(addr, "'", "''"))
		sb.WriteString("'")
	}

	sb.WriteString(`) AND tx_to IN (`)

	// Add placeholders for addresses again
	for i, addr := range addresses {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString("'")
		sb.WriteString(strings.ReplaceAll(addr, "'", "''"))
		sb.WriteString("'")
	}

	sb.WriteString(`))
		AND tx_from != tx_to
		ORDER BY
			block_timestamp DESC
		LIMIT `)
	sb.WriteString(fmt.Sprintf("%d", limit))

	sql := sb.String()

	queryRun, err := c.createQueryRun(sql)
	if err != nil {
		return nil, err
	}

	results, err := c.waitForQueryResults(queryRun.Result.QueryRun.ID)
	if err != nil {
		return nil, err
	}

	interactionsMap := make(AddressInteractionsMap)

	// Initialize the map with empty slices for each address
	for _, addr := range addresses {
		interactionsMap[addr] = &AddressInteractions{
			Address:           addr,
			SentTransfers:     []TransferData{},
			ReceivedTransfers: []TransferData{},
		}
	}

	for _, record := range results.Result.Rows {
		data, ok := record.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("unexpected data type in result row")
		}

		blockID, err := parseFloat64(data["block_id"])
		if err != nil {
			return nil, fmt.Errorf("error parsing block_id: %w", err)
		}

		amount, err := parseFloat64(data["amount"])
		if err != nil {
			return nil, fmt.Errorf("error parsing amount: %w", err)
		}

		transferData := TransferData{
			BlockTimestamp:    parseTimestamp(data["block_timestamp"].(string)),
			BlockID:           int64(blockID),
			TxID:              data["tx_id"].(string),
			TxFrom:            data["tx_from"].(string),
			TxTo:              data["tx_to"].(string),
			Amount:            amount,
			Mint:              data["mint"].(string),
			FactTransfersID:   data["fact_transfers_id"].(string),
			InsertedTimestamp: parseTimestamp(data["inserted_timestamp"].(string)),
			ModifiedTimestamp: parseTimestamp(data["modified_timestamp"].(string)),
		}

		// Add to sender's sent transfers
		if interactions, exists := interactionsMap[transferData.TxFrom]; exists {
			interactions.SentTransfers = append(interactions.SentTransfers, transferData)
		}

		// Add to receiver's received transfers
		if interactions, exists := interactionsMap[transferData.TxTo]; exists {
			interactions.ReceivedTransfers = append(interactions.ReceivedTransfers, transferData)
		}
	}

	return interactionsMap, nil
}

func (c *Client) GetFirstSwaps(addresses []string) ([]FirstSwapData, error) {
	var sb strings.Builder
	sb.WriteString(`
		WITH first_swaps AS (
			SELECT
				swapper,
				MIN(block_timestamp) AS first_swap_timestamp
			FROM
				solana.defi.ez_dex_swaps
			WHERE
				swapper IN (`)

	// Add placeholders for addresses
	for i, addr := range addresses {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString("'")
		sb.WriteString(strings.ReplaceAll(addr, "'", "''"))
		sb.WriteString("'")
	}

	sb.WriteString(`)
		GROUP BY
			swapper
		)
		SELECT
			swapper AS address,
			first_swap_timestamp
		FROM
			first_swaps
		ORDER BY
			first_swap_timestamp ASC`)

	sql := sb.String()

	queryRun, err := c.createQueryRun(sql)
	if err != nil {
		return nil, err
	}

	results, err := c.waitForQueryResults(queryRun.Result.QueryRun.ID)
	if err != nil {
		return nil, err
	}

	var firstSwaps []FirstSwapData

	for _, record := range results.Result.Rows {
		data, ok := record.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("unexpected data type in result row")
		}

		address := data["address"].(string)
		timestamp := parseTimestamp(data["first_swap_timestamp"].(string))

		firstSwaps = append(firstSwaps, FirstSwapData{
			Address:   address,
			Timestamp: timestamp,
		})
	}

	return firstSwaps, nil
}
