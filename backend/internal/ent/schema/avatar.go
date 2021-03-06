// Code generated by windranger, DO NOT EDIT.
package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/dialect/entsql"
    "entgo.io/ent/schema"
    "entgo.io/ent/schema/field"
)

// Avatar 保存Avatar实体的结构定义
type Avatar struct {
    ent.Schema
}

// Annotations 配置使用单数表名
func (Avatar)Annotations() []schema.Annotation {
    return []schema.Annotation{
        entsql.Annotation{
            Table: "avatar",
        },
    }
}

// Fields 定义Avatar实体的字段
func (Avatar) Fields() []ent.Field{
    return []ent.Field{
        field.String("id").Comment("主键"),
        field.Int32("kind").Optional().Nillable().Comment("头像类型"),
        field.String("rel_id").Optional().Nillable().Comment("关联ID"),
        field.String("avatar").Optional().Nillable().Comment("头像base64串"),
    }
}
