package token

import (
	"fmt"
)

func (issue *IssueTokenRequest) Validate() error {
	switch issue.GrantType {
	case GrantType_PASSWORD:
		if issue.Username == "" || issue.Password == "" {
			return fmt.Errorf("password grant required username and password")
		}
	default:
		return fmt.Errorf("grant type %s not implemented", issue.GrantType)
	}

	return nil
}
