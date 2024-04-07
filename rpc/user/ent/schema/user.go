package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Immutable(),
		field.String("account").
			Comment("账号").
			Annotations(entsql.WithComments(true)),
		field.String("password").
			Comment("密码").
			Annotations(entsql.WithComments(true)),
		field.String("name").
			Comment("名称").
			Annotations(entsql.WithComments(true)),
		field.String("real_name").
			Comment("真实姓名").
			Annotations(entsql.WithComments(true)),
		field.String("phone").
			Comment("电话").
			Annotations(entsql.WithComments(true)),
		field.String("address").
			Comment("电话").
			Annotations(entsql.WithComments(true)),
		field.Time("create_at").
			Comment("create time | 创建时间").
			Default(time.Now).
			Annotations(entsql.WithComments(true)),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

// 用户实体的注解
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "storm_user"},
	}
}
