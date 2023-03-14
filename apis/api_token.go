package apis

import (
	"context"
	"errors"
	"github.com/cenkalti/backoff/v4"
	"log"
	"sync"
	"time"
)

// 分布式access_token：获取和设置access_token的值，自行实现该接口的具体逻辑，比如使用redis方案
type DcsToken interface {
	Get(cacheKey string) TokenInfo                                     // 获取缓存
	Set(cacheKey string, tokenInfo TokenInfo, ttl time.Duration) error // 设置缓存，ttl：缓存生存时间
	Del(cacheKey string) error                                         // 删除缓存
	Lock(cacheKey string, ttl time.Duration) bool                      // 加锁，返回成功或失败
	Unlock(cacheKey string) error                                      // 释放锁
}

type TokenInfo struct {
	Token       string        `json:"token"`        // access_token/jsapi_ticket
	ExpiresIn   time.Duration `json:"expires_in"`   // 过期时间
	LastRefresh time.Time     `json:"last_refresh"` // 上次刷新access_token时间
}

type token struct {
	mutex *sync.RWMutex
	TokenInfo
	getTokenFunc  func() (TokenInfo, error)
	dcsToken      DcsToken
	tokenCacheKey string
}

func (c *ApiClient) GetToken() (token string, err error) {
	token = c.accessToken.getToken()
	if token == "" {
		err = errors.New("access_token获取失败")
		return
	}
	return
}

func (t *token) setGetTokenFunc(f func() (TokenInfo, error)) {
	t.getTokenFunc = f
}

func (t *token) getToken() string {
	if err := Retry(t.syncToken); err != nil {
		log.Printf("retry getting access Token failed, err=%+v\n", err)
	}

	t.mutex.RLock()
	tokenToUse := t.Token
	t.mutex.RUnlock()

	return tokenToUse
}

func (t *token) syncToken() error {
	var refreshHour int64 = 3600 // access_token刷新时间间隔，单位秒
	var now = time.Now()

	var tokenInfo TokenInfo
	if t.dcsToken != nil {
		tokenInfo = t.dcsToken.Get(t.tokenCacheKey)

		if tokenInfo.Token == "" || tokenInfo.LastRefresh.Unix()+refreshHour <= now.Unix() {
			lockCacheKey := t.tokenCacheKey + "#lock"

			// 抢锁
			if ok := t.dcsToken.Lock(lockCacheKey, 10*time.Second); ok {
				defer func() {
					_ = t.dcsToken.Unlock(lockCacheKey)
				}()

				get, err := t.getTokenFunc()
				if err != nil {
					return err
				}

				tokenInfo.Token = get.Token
				tokenInfo.ExpiresIn = get.ExpiresIn
				tokenInfo.LastRefresh = now

				if err := t.dcsToken.Set(t.tokenCacheKey, tokenInfo, tokenInfo.ExpiresIn); err != nil {
					return err
				}
			} else {
				// 抢锁失败则等待
				time.Sleep(time.Second * 2)
				tokenInfo = t.dcsToken.Get(t.tokenCacheKey)
			}
		}
	} else {
		if t.Token == "" || t.LastRefresh.Unix()+refreshHour <= now.Unix() {
			get, err := t.getTokenFunc()
			if err != nil {
				return err
			}
			tokenInfo.Token = get.Token
			tokenInfo.ExpiresIn = get.ExpiresIn
			tokenInfo.LastRefresh = now
		}
	}

	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.Token = tokenInfo.Token
	t.ExpiresIn = tokenInfo.ExpiresIn
	t.LastRefresh = tokenInfo.LastRefresh

	return nil
}

func Retry(o backoff.Operation) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancelFunc()
	retryer := backoff.WithContext(backoff.NewExponentialBackOff(), ctx)
	return backoff.Retry(o, retryer)
}

// 获取access_token
func (c *ApiClient) getAccessToken() (TokenInfo, error) {
	req := ReqGetAccessToken{
		GrantType: "client_credential",
		Appid:     c.AppId,
		Secret:    c.AppSecret,
	}
	get, err := c.ExecGetAccessToken(req)
	if err != nil {
		c.logger.Errorf(c.accessTokenName+": req=%+v, err=%+v\n", req, err)
		return TokenInfo{}, err
	}
	return TokenInfo{Token: get.AccessToken, ExpiresIn: time.Duration(get.ExpiresIn) * time.Second}, nil
}
