package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/casbin/casbin/v2"
	"github.com/redis/go-redis/v9"
	"github.com/saas-mingyang/mingyang-admin-common/config"
	"github.com/saas-mingyang/mingyang-admin-common/orm/ent/entctx/rolectx"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/saas-mingyang/mingyang-admin-common/utils/jwt"
)

type AuthorityMiddleware struct {
	Cbn *casbin.Enforcer
	Rds redis.UniversalClient
}

func NewAuthorityMiddleware(cbn *casbin.Enforcer, rds redis.UniversalClient) *AuthorityMiddleware {
	return &AuthorityMiddleware{
		Cbn: cbn,
		Rds: rds,
	}
}

func (m *AuthorityMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the path
		obj := r.URL.Path
		// get the method
		act := r.Method

		// get the role id with enhanced type handling
		var roleIds []string
		var err error

		if roleIdVal := r.Context().Value("roleId"); roleIdVal != nil {
			switch v := roleIdVal.(type) {
			case string:
				roleIds = strings.Split(v, ",")
			case json.Number:
				roleIds = []string{v.String()}
			case int64:
				roleIds = []string{fmt.Sprintf("%d", v)}
			case uint64:
				roleIds = []string{fmt.Sprintf("%d", v)}
			case int:
				roleIds = []string{fmt.Sprintf("%d", v)}
			default:
				roleIds = []string{fmt.Sprintf("%v", v)}
			}
		} else {
			// Fallback to original function
			roleIds, err = rolectx.GetRoleIDFromCtx(r.Context())
			if err != nil {
				httpx.Error(w, err)
				return
			}
		}

		// check jwt blacklist
		jwtResult, err := m.Rds.Get(context.Background(), config.RedisTokenPrefix+jwt.StripBearerPrefixFromToken(r.Header.Get("Authorization"))).Result()
		if err != nil && !errors.Is(err, redis.Nil) {
			logx.Errorw("redis error in jwt", logx.Field("detail", err.Error()))
			httpx.Error(w, errorx.NewApiError(http.StatusInternalServerError, err.Error()))
			return
		}
		if jwtResult == "1" {
			logx.Errorw("token in blacklist", logx.Field("detail", r.Header.Get("Authorization")))
			httpx.Error(w, errorx.NewApiErrorWithoutMsg(http.StatusUnauthorized))
			return
		}

		result := batchCheck(m.Cbn, roleIds, act, obj)

		if result {
			// Handle different types for userId (string, json.Number, int64, uint64)
			var userIdStr string
			if userIdVal := r.Context().Value("userId"); userIdVal != nil {
				switch v := userIdVal.(type) {
				case string:
					userIdStr = v
				case json.Number:
					userIdStr = v.String()
				case int64:
					userIdStr = fmt.Sprintf("%d", v)
				case uint64:
					userIdStr = fmt.Sprintf("%d", v)
				case int:
					userIdStr = fmt.Sprintf("%d", v)
				default:
					userIdStr = fmt.Sprintf("%v", v)
				}
			}

			logx.Infow("HTTP/HTTPS Request", logx.Field("UUID", userIdStr),
				logx.Field("path", obj), logx.Field("method", act))
			next(w, r)
			return
		} else {
			logx.Errorw("the role is not permitted to access the API", logx.Field("roleId", roleIds),
				logx.Field("path", obj), logx.Field("method", act))
			httpx.Error(w, errorx.NewApiForbiddenError("You do not have permission to access the API"))
			return
		}
	}
}

func batchCheck(cbn *casbin.Enforcer, roleIds []string, act, obj string) bool {
	var checkReq [][]any
	for _, v := range roleIds {
		checkReq = append(checkReq, []any{v, obj, act})
	}

	result, err := cbn.BatchEnforce(checkReq)
	if err != nil {
		logx.Errorw("Casbin enforce error", logx.Field("detail", err.Error()))
		return false
	}

	for _, v := range result {
		if v {
			return true
		}
	}

	return false
}
