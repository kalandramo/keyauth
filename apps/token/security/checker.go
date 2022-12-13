package security

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/cache"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/kalandramo/keyauth/apps/domain"
	"github.com/kalandramo/keyauth/apps/token"
	"github.com/kalandramo/keyauth/apps/user"
)

// NewChecker todo
func NewChecker() (Checker, error) {
	c := cache.C()
	if c == nil {
		return nil, fmt.Errorf("depended cache service is nil")
	}

	return &checker{
		domain: app.GetGrpcApp(domain.AppName).(domain.ServiceServer),
		user:   app.GetGrpcApp(user.AppName).(user.ServiceServer),
		cache:  c,
		log:    zap.L().Named("Login Security"),
	}, nil
}

type checker struct {
	domain domain.ServiceServer
	user   user.ServiceServer
	cache  cache.Cache
	log    logger.Logger
}

func (c *checker) MaxFailedRetryCheck(ctx context.Context, req *token.IssueTokenRequest) error {
	ss := c.getOrDefaultSecuritySettingWithUser(ctx, req.Username)
	if !ss.LoginSecurity.RetryLock {
		c.log.Debugf("retry lock check disabled, don't check")
		return nil
	}
	c.log.Debugf("max failed retry lock check enabled, checking ...")

	var count int32
	err := c.cache.Get(req.AbnormalUserCheckKey(), &count)
	if err != nil {
		c.log.Errorf("get key %s from cache error, %s", req.AbnormalUserCheckKey(), err)
	}
	rlc := ss.LoginSecurity.RetryLockConfig
	c.log.Debugf("retry times: %d, retry limit: %d", count, rlc.RetryLimite)
	if count+1 >= int32(rlc.RetryLimite) {
		return fmt.Errorf("登录失败次数过多, 请%d分钟后重试", rlc.LockedMinite)
	}

	return nil
}

func (c *checker) UpdateFailedRetry(ctx context.Context, req *token.IssueTokenRequest) error {
	ss := c.getOrDefaultSecuritySettingWithUser(ctx, req.Username)
	if !ss.LoginSecurity.RetryLock {
		c.log.Debugf("retry lock check disabled, don't check")
		return nil
	}

	c.log.Debugf("update failed retry count, check key: %s", req.AbnormalUserCheckKey())

	var count int
	if err := c.cache.Get(req.AbnormalUserCheckKey(), &count); err == nil {
		// 之前已经登陆失败过
		err := c.cache.Put(req.AbnormalUserCheckKey(), count+1)
		if err != nil {
			c.log.Errorf("set key %s to cache error, %s", req.AbnormalUserCheckKey())
		}
	} else {
		// 首次登陆失败
		err := c.cache.PutWithTTL(
			req.AbnormalUserCheckKey(),
			count+1,
			ss.LoginSecurity.RetryLockConfig.LockedMiniteDuration(),
		)
		if err != nil {
			c.log.Errorf("set key %s to cache error, %s", req.AbnormalUserCheckKey())
		}
	}

	return nil
}

func (c *checker) getOrDefaultSecuritySettingWithUser(ctx context.Context, account string) *domain.SecuritySetting {
	ss := domain.NewDefaultSecuritySetting()
	descReq := user.NewDescriptAccountRequestWithAccount(account)
	u, err := c.user.DescribeAccount(ctx, descReq)
	if err != nil {
		c.log.Errorf("get user account error, %s, use default setting to check", err)
		return ss
	}

	return c.getOrDefaultSecuritySettingWithDomain(ctx, u.Account, u.Domain)
}

func (c *checker) getOrDefaultSecuritySettingWithDomain(ctx context.Context, account, domainName string) *domain.SecuritySetting {
	ss := domain.NewDefaultSecuritySetting()
	descReq := domain.NewDescribeDomainRequestWithName(domainName)
	d, err := c.domain.DescribeDomain(ctx, descReq)
	if err != nil {
		c.log.Errorf("get domain error, %s, use default setting to check", err)
		return ss
	}

	return d.SecuritySetting
}
