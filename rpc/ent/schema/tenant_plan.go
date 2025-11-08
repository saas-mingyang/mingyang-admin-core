package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/saas-mingyang/mingyang-admin-common/orm/ent/mixins"
	mixins2 "github.com/saas-mingyang/mingyang-admin-core/rpc/ent/schema/mixins"
)

type TenantPlan struct {
	ent.Schema
}

func (TenantPlan) Fields() []ent.Field {
	return []ent.Field{
		field.String("package_name").Unique().Comment("套餐名称"),
		field.Strings("menu_ids").Comment("菜单ID"),
		field.Strings("remark").Comment("备注"),
		field.Int("menu_check_strictly").Comment("菜单树选择项是否关联显示"),
	}
}

func (TenantPlan) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
		mixins2.SoftDeleteMixin{},
	}
}

func (TenantPlan) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (TenantPlan) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		schema.Comment("tenant Table | 租户套餐计划表"),
		entsql.Annotation{Table: "sy_tenant_plans"},
	}
}

func (TenantPlan) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id", "status"),
	}
}
