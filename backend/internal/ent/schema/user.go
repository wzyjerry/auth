// Code generated by windranger, DO NOT EDIT.
package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/dialect/entsql"
    "entgo.io/ent/schema"
    "entgo.io/ent/schema/field"
)

// User 保存User实体的结构定义
type User struct {
    ent.Schema
}

// Annotations 配置使用单数表名
func (User)Annotations() []schema.Annotation {
    return []schema.Annotation{
        entsql.Annotation{
            Table: "user",
        },
    }
}

// Fields 定义User实体的字段
func (User) Fields() []ent.Field{
    return []ent.Field{
        field.String("id").Comment("主键"),
        field.String("ancestor_id").Optional().Nillable().Comment("祖先ID"),
        field.String("password").Optional().Nillable().Comment("密码"),
        field.String("nickname").Optional().Nillable().Comment("昵称"),
        field.String("ip").Optional().Nillable().Comment("注册IP"),
    }
}
