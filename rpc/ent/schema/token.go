package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/saas-mingyang/mingyang-admin-common/orm/ent/mixins"
)

type Token struct {
	ent.Schema
}

func (Token) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").Optional().Default(1).
			Comment(" User's ID | 用户的ID"),
		field.String("username").
			Comment("Username | 用户名").
			Default("unknown"),
		field.String("token").
			Comment("Token string | Token 字符串"),
		field.String("source").
			Comment("Log in source such as GitHub | Token 来源 （本地为core, 第三方如github等）"),
		field.Time("expired_at").
			Comment(" Expire time | 过期时间"),
	}
}

func (Token) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IdSonyFlakeMixin{},
		mixins.StatusMixin{},
	}
}

func (Token) Edges() []ent.Edge {
	return nil
}

func (Token) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("expired_at"),
	}
}

func (Token) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		schema.Comment("Token Log Table | 令牌信息表"),
		entsql.Annotation{Table: "sys_tokens"},
	}
}
