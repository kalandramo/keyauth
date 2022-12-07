package security

import (
	"context"

	"github.com/infraboard/mcube/logger"
	"github.com/kalandramo/keyauth/apps/token"
	"github.com/kalandramo/keyauth/apps/user"
)

type checker struct {
	user user.ServiceServer
	log  logger.Logger
}

func (c *checker) MaxFailedRetryCheck(ctx context.Context, req *token.IssueTokenRequest) error {
	return nil
}
