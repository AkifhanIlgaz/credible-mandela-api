package mande

type queryRequest struct {
	OperationName string         `json:"operationName"`
	Query         string         `json:"query"`
	Variables     map[string]any `json:"variables"`
}

const (
	credQueryOperationName string = "CredQuery"
)

const credQuery string = `
query CredQuery($address: ID!) {
  user(id: $address) {
    credScoreAccrued
  }
}
`

type user struct {
	CredScoreAccrued string `json:"credScoreAccrued"`
}

type data struct {
	User user `json:"user"`
}

type credQueryResponse struct {
	Data data `json:"data"`
}

func generateCredQueryRequest(address string) queryRequest {
	return queryRequest{
		OperationName: credQueryOperationName,
		Query:         credQuery,
		Variables: map[string]any{
			"address": address,
		},
	}
}
