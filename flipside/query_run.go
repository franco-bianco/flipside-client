package flipside

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

func (c *Client) createQueryRun(sql string) (*CreateQueryRunResponse, error) {
	c.log.Info("creating query run...")

	params := map[string]interface{}{
		"sql":            sql,
		"dataSource":     "snowflake-default",
		"dataProvider":   "flipside",
		"resultTTLHours": 1,
		"maxAgeMinutes":  60,
		"tags": map[string]string{
			"source": "go-sdk",
		},
	}

	resp, err := c.sendRequest("createQueryRun", params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, _ := io.ReadAll(resp.Body)

	var response CreateQueryRunResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) waitForQueryResults(queryRunID string) (*GetQueryRunResultsResponse, error) {
	c.log.Info("waiting for query results...")

	time.Sleep(5 * time.Second)

	timeout := time.After(10 * time.Minute)
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-timeout:
			return nil, fmt.Errorf("query execution timed out after 10 minutes")
		case <-ticker.C:
			params := map[string]string{
				"queryRunId": queryRunID,
			}

			resp, err := c.sendRequest("getQueryRun", params)
			if err != nil {
				return nil, err
			}
			defer resp.Body.Close()

			responseBody, _ := io.ReadAll(resp.Body)

			var response GetQueryRunResponse
			err = json.Unmarshal(responseBody, &response)
			if err != nil {
				return nil, err
			}

			switch response.Result.QueryRun.State {
			case "QUERY_STATE_SUCCESS":
				return c.getQueryResults(queryRunID)
			case "QUERY_STATE_RUNNING", "QUERY_STATE_READY":
				// Continue waiting
			default:
				return nil, fmt.Errorf("query failed with state: %s", response.Result.QueryRun.State)
			}
		}
	}
}

func (c *Client) getQueryResults(queryRunID string) (*GetQueryRunResultsResponse, error) {
	c.log.Info("getting query results...")

	params := map[string]interface{}{
		"queryRunId": queryRunID,
		"format":     "json",
		"page": map[string]int{
			"number": 1,
			"size":   1000,
		},
	}

	resp, err := c.sendRequest("getQueryRunResults", params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, _ := io.ReadAll(resp.Body)

	var response GetQueryRunResultsResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
