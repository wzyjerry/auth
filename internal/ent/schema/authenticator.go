// Code generated by windranger, DO NOT EDIT.
package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/dialect/entsql"
    "entgo.io/ent/schema"
    "entgo.io/ent/schema/field"
    "github.com/wzyjerry/auth/internal/ent/schema/authenticatorNested"
)

// Authenticator 保存Authenticator实体的结构定义
type Authenticator struct {
    ent.Schema
}

// Annotations 配置使用单数表名
func (Authenticator)Annotations() []schema.Annotation {
    return []schema.Annotation{
        entsql.Annotation{
            Table: "authenticator",
        },
    }
}

// Fields 定义Authenticator实体的字段
func (Authenticator) Fields() []ent.Field{
    return []ent.Field{
        field.String("id").Comment("主键"),
        field.String("user_id").Optional().Nillable().Comment("用户ID"),
        field.Int32("kind").Optional().Nillable().Comment("认证器类型"),
        field.JSON("anchor", new(authenticatorNested.Anchor)).Optional().Comment("锚"),
    }
}