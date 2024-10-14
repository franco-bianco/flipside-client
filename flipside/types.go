package flipside

import "time"

// SwapData represents the data for a swap transaction
type SwapData struct {
	SwapProgram       string    `json:"swap_program"`
	BlockID           int64     `json:"block_id"`
	BlockTimestamp    time.Time `json:"block_timestamp"`
	TxID              string    `json:"tx_id"`
	ProgramID         string    `json:"program_id"`
	Swapper           string    `json:"swapper"`
	SwapFromMint      string    `json:"swap_from_mint"`
	SwapFromAmount    float64   `json:"swap_from_amount"`
	SwapToMint        string    `json:"swap_to_mint"`
	SwapToAmount      float64   `json:"swap_to_amount"`
	LogID             string    `json:"_log_id"`
	EzSwapsID         string    `json:"ez_swaps_id"`
	InsertedTimestamp time.Time `json:"inserted_timestamp"`
	ModifiedTimestamp time.Time `json:"modified_timestamp"`
}

// TransferData represents the data for a token transfer transaction
type TransferData struct {
	BlockTimestamp    time.Time `json:"block_timestamp"`
	BlockID           int64     `json:"block_id"`
	TxID              string    `json:"tx_id"`
	TxFrom            string    `json:"tx_from"`
	TxTo              string    `json:"tx_to"`
	Amount            float64   `json:"amount"`
	Mint              string    `json:"mint"`
	FactTransfersID   string    `json:"fact_transfers_id"`
	InsertedTimestamp time.Time `json:"inserted_timestamp"`
	ModifiedTimestamp time.Time `json:"modified_timestamp"`
}

// AddressInteractions represents interactions for a single address
type AddressInteractions struct {
	Address           string
	SentTransfers     []TransferData
	ReceivedTransfers []TransferData
}

// AddressInteractionsMap is a map of addresses to their interactions
type AddressInteractionsMap map[string]*AddressInteractions

type CreateQueryRunResponse struct {
	Jsonrpc string `json:"jsonrpc,omitempty"`
	ID      int    `json:"id,omitempty"`
	Result  struct {
		QueryRequest struct {
			ID             string `json:"id,omitempty"`
			SQLStatementID string `json:"sqlStatementId,omitempty"`
			UserID         string `json:"userId,omitempty"`
			Tags           struct {
				Source string `json:"source,omitempty"`
			} `json:"tags,omitempty"`
			MaxAgeMinutes     int       `json:"maxAgeMinutes,omitempty"`
			ResultTTLHours    int       `json:"resultTTLHours,omitempty"`
			UserSkipCache     bool      `json:"userSkipCache,omitempty"`
			TriggeredQueryRun bool      `json:"triggeredQueryRun,omitempty"`
			QueryRunID        string    `json:"queryRunId,omitempty"`
			CreatedAt         time.Time `json:"createdAt,omitempty"`
			UpdatedAt         time.Time `json:"updatedAt,omitempty"`
		} `json:"queryRequest,omitempty"`
		QueryRun struct {
			ID                    string `json:"id,omitempty"`
			SQLStatementID        string `json:"sqlStatementId,omitempty"`
			State                 string `json:"state,omitempty"`
			Path                  string `json:"path,omitempty"`
			FileCount             any    `json:"fileCount,omitempty"`
			LastFileNumber        any    `json:"lastFileNumber,omitempty"`
			FileNames             any    `json:"fileNames,omitempty"`
			ErrorName             any    `json:"errorName,omitempty"`
			ErrorMessage          any    `json:"errorMessage,omitempty"`
			ErrorData             any    `json:"errorData,omitempty"`
			DataSourceQueryID     any    `json:"dataSourceQueryId,omitempty"`
			DataSourceSessionID   any    `json:"dataSourceSessionId,omitempty"`
			StartedAt             any    `json:"startedAt,omitempty"`
			QueryRunningEndedAt   any    `json:"queryRunningEndedAt,omitempty"`
			QueryStreamingEndedAt any    `json:"queryStreamingEndedAt,omitempty"`
			EndedAt               any    `json:"endedAt,omitempty"`
			RowCount              any    `json:"rowCount,omitempty"`
			TotalSize             any    `json:"totalSize,omitempty"`
			Tags                  struct {
				Source string `json:"source,omitempty"`
			} `json:"tags,omitempty"`
			DataSourceID            string    `json:"dataSourceId,omitempty"`
			UserID                  string    `json:"userId,omitempty"`
			CreatedAt               time.Time `json:"createdAt,omitempty"`
			UpdatedAt               time.Time `json:"updatedAt,omitempty"`
			ArchivedAt              any       `json:"archivedAt,omitempty"`
			RowsPerResultSet        int       `json:"rowsPerResultSet,omitempty"`
			StatementTimeoutSeconds int       `json:"statementTimeoutSeconds,omitempty"`
			AbortDetachedQuery      bool      `json:"abortDetachedQuery,omitempty"`
		} `json:"queryRun,omitempty"`
		SQLStatement struct {
			ID             string `json:"id,omitempty"`
			StatementHash  string `json:"statementHash,omitempty"`
			SQL            string `json:"sql,omitempty"`
			ColumnMetadata any    `json:"columnMetadata,omitempty"`
			UserID         string `json:"userId,omitempty"`
			Tags           struct {
				Source string `json:"source,omitempty"`
			} `json:"tags,omitempty"`
			CreatedAt time.Time `json:"createdAt,omitempty"`
			UpdatedAt time.Time `json:"updatedAt,omitempty"`
		} `json:"sqlStatement,omitempty"`
	} `json:"result,omitempty"`
}

type GetQueryRunResponse struct {
	Jsonrpc string `json:"jsonrpc,omitempty"`
	ID      int    `json:"id,omitempty"`
	Result  struct {
		QueryRun struct {
			ID                    string    `json:"id,omitempty"`
			SQLStatementID        string    `json:"sqlStatementId,omitempty"`
			State                 string    `json:"state,omitempty"`
			Path                  string    `json:"path,omitempty"`
			FileCount             int       `json:"fileCount,omitempty"`
			LastFileNumber        int       `json:"lastFileNumber,omitempty"`
			FileNames             string    `json:"fileNames,omitempty"`
			ErrorName             any       `json:"errorName,omitempty"`
			ErrorMessage          any       `json:"errorMessage,omitempty"`
			ErrorData             any       `json:"errorData,omitempty"`
			DataSourceQueryID     any       `json:"dataSourceQueryId,omitempty"`
			DataSourceSessionID   string    `json:"dataSourceSessionId,omitempty"`
			StartedAt             time.Time `json:"startedAt,omitempty"`
			QueryRunningEndedAt   time.Time `json:"queryRunningEndedAt,omitempty"`
			QueryStreamingEndedAt time.Time `json:"queryStreamingEndedAt,omitempty"`
			EndedAt               time.Time `json:"endedAt,omitempty"`
			RowCount              int       `json:"rowCount,omitempty"`
			TotalSize             string    `json:"totalSize,omitempty"`
			Tags                  struct {
				SdkPackage  string `json:"sdk_package,omitempty"`
				SdkVersion  string `json:"sdk_version,omitempty"`
				SdkLanguage string `json:"sdk_language,omitempty"`
			} `json:"tags,omitempty"`
			DataSourceID string    `json:"dataSourceId,omitempty"`
			UserID       string    `json:"userId,omitempty"`
			CreatedAt    time.Time `json:"createdAt,omitempty"`
			UpdatedAt    time.Time `json:"updatedAt,omitempty"`
			ArchivedAt   any       `json:"archivedAt,omitempty"`
		} `json:"queryRun,omitempty"`
		RedirectedToQueryRun any `json:"redirectedToQueryRun,omitempty"`
	} `json:"result,omitempty"`
}

type QueryRunStatusResponse struct {
	State       string    `json:"state"`
	StartedAt   time.Time `json:"startedAt"`
	EndedAt     time.Time `json:"endedAt"`
	ElapsedTime float64   `json:"elapsedSeconds"`
	RecordCount int       `json:"recordCount"`
}

type GetQueryRunResultsResponse struct {
	Jsonrpc string `json:"jsonrpc,omitempty"`
	ID      int    `json:"id,omitempty"`
	Result  struct {
		ColumnNames []string      `json:"columnNames,omitempty"`
		ColumnTypes []string      `json:"columnTypes,omitempty"`
		Rows        []interface{} `json:"rows,omitempty"`
		Page        struct {
			CurrentPageNumber int `json:"currentPageNumber,omitempty"`
			CurrentPageSize   int `json:"currentPageSize,omitempty"`
			TotalRows         int `json:"totalRows,omitempty"`
			TotalPages        int `json:"totalPages,omitempty"`
		} `json:"page,omitempty"`
		SQL              string `json:"sql,omitempty"`
		Format           string `json:"format,omitempty"`
		OriginalQueryRun struct {
			ID                    string    `json:"id,omitempty"`
			SQLStatementID        string    `json:"sqlStatementId,omitempty"`
			State                 string    `json:"state,omitempty"`
			Path                  string    `json:"path,omitempty"`
			FileCount             int       `json:"fileCount,omitempty"`
			LastFileNumber        int       `json:"lastFileNumber,omitempty"`
			FileNames             string    `json:"fileNames,omitempty"`
			ErrorName             any       `json:"errorName,omitempty"`
			ErrorMessage          any       `json:"errorMessage,omitempty"`
			ErrorData             any       `json:"errorData,omitempty"`
			DataSourceQueryID     any       `json:"dataSourceQueryId,omitempty"`
			DataSourceSessionID   string    `json:"dataSourceSessionId,omitempty"`
			StartedAt             time.Time `json:"startedAt,omitempty"`
			QueryRunningEndedAt   time.Time `json:"queryRunningEndedAt,omitempty"`
			QueryStreamingEndedAt time.Time `json:"queryStreamingEndedAt,omitempty"`
			EndedAt               time.Time `json:"endedAt,omitempty"`
			RowCount              int       `json:"rowCount,omitempty"`
			TotalSize             string    `json:"totalSize,omitempty"`
			Tags                  struct {
				SdkPackage  string `json:"sdk_package,omitempty"`
				SdkVersion  string `json:"sdk_version,omitempty"`
				SdkLanguage string `json:"sdk_language,omitempty"`
			} `json:"tags,omitempty"`
			DataSourceID string    `json:"dataSourceId,omitempty"`
			UserID       string    `json:"userId,omitempty"`
			CreatedAt    time.Time `json:"createdAt,omitempty"`
			UpdatedAt    time.Time `json:"updatedAt,omitempty"`
			ArchivedAt   any       `json:"archivedAt,omitempty"`
		} `json:"originalQueryRun,omitempty"`
		RedirectedToQueryRun any `json:"redirectedToQueryRun,omitempty"`
	} `json:"result,omitempty"`
}

// FirstSwapData represents the first swap data for an address
type FirstSwapData struct {
	Address   string    `json:"address"`
	Timestamp time.Time `json:"timestamp"`
}
