	import "encoding/json"

	type QueryResult struct {
		Errors           []json.RawMessage
		Warnings         []json.RawMessage
		TotalResultCount int64
		StartIndex       int64
		PageSize         int64
		Results          []interface{}
	}
