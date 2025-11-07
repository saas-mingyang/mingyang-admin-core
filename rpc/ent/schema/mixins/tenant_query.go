package mixins

import (
	"context"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/mixin"
	"github.com/saas-mingyang/mingyang-admin-common/orm/ent/entctx/tenantctx"
	"github.com/saas-mingyang/mingyang-admin-core/rpc/ent/intercept"
	"log"
)

// TenantQueryMixin defines the tenant query mixin.
type TenantQueryMixin struct {
	mixin.Schema
}

// Interceptors SoftDeleteMixin 的拦截器。
func (d TenantQueryMixin) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		intercept.TraverseFunc(func(ctx context.Context, q intercept.Query) error {
			tenantId := tenantctx.GetTenantIDFromCtx(ctx)
			log.Println("tenantId: ", tenantId)
			q.WhereP(func(s *sql.Selector) {
				s.Where(sql.EQ(s.C("tenant_id"), tenantId))
			})
			return nil
		}),
	}

}
