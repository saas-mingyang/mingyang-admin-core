package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/saas-mingyang/mingyang-admin-common/orm/ent/mixins"
)

type Tenant struct {
	ent.Schema
}

func (Tenant) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().NotEmpty().MaxLen(128).Comment("名称，不能为空"),
		field.String("code").Unique().NotEmpty().MaxLen(128).Comment("编码，不能为空"),
		field.String("contact_phone").Comment("联系方式"),
		field.String("contact_email").Comment("联系邮箱"),
		field.String("company_name").Comment("企业名称"),
		field.String("license_number").Comment("统一社会信用代码"),
		field.String("address").Comment("地址"),
		field.String("intro").Comment("企业简介"),
		field.String("domain").Comment("域名"),
		field.Int("level").Comment("租户级别"),
		field.Int64("plan_id").Optional().Comment("套餐计划Id，外键关联 tenant_plan.id"),
		field.Int64("admin_id").Positive().Comment("管理员id"),
		field.Int64("parent_id").Positive().Comment("父级id,0为第一级"),
	}
}

func (Tenant) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Tenant) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IdSonyFlakeMixin{},
		mixins.StatusMixin{},
	}
}

func (Tenant) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		schema.Comment("tenant Table | 租户表"),
		entsql.Annotation{Table: "sys_tenants"},
	}
}

func (Tenant) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id", "status"),
		index.Fields("parent_id", "status"),
	}
}
