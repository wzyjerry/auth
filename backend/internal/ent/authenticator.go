// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/wzyjerry/auth/internal/ent/authenticator"
	"github.com/wzyjerry/auth/internal/ent/schema/authenticatorNested"
)

// Authenticator is the model entity for the Authenticator schema.
type Authenticator struct {
	config `json:"-"`
	// ID of the ent.
	// 主键
	ID string `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	// 用户ID
	UserID *string `json:"user_id,omitempty"`
	// Kind holds the value of the "kind" field.
	// 认证器类型
	Kind *int32 `json:"kind,omitempty"`
	// Unique holds the value of the "unique" field.
	// 唯一值
	Unique *authenticatorNested.Unique `json:"unique,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Authenticator) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case authenticator.FieldUnique:
			values[i] = new([]byte)
		case authenticator.FieldKind:
			values[i] = new(sql.NullInt64)
		case authenticator.FieldID, authenticator.FieldUserID:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Authenticator", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Authenticator fields.
func (a *Authenticator) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case authenticator.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				a.ID = value.String
			}
		case authenticator.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				a.UserID = new(string)
				*a.UserID = value.String
			}
		case authenticator.FieldKind:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field kind", values[i])
			} else if value.Valid {
				a.Kind = new(int32)
				*a.Kind = int32(value.Int64)
			}
		case authenticator.FieldUnique:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field unique", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &a.Unique); err != nil {
					return fmt.Errorf("unmarshal field unique: %w", err)
				}
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Authenticator.
// Note that you need to call Authenticator.Unwrap() before calling this method if this Authenticator
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Authenticator) Update() *AuthenticatorUpdateOne {
	return (&AuthenticatorClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Authenticator entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Authenticator) Unwrap() *Authenticator {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Authenticator is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Authenticator) String() string {
	var builder strings.Builder
	builder.WriteString("Authenticator(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	if v := a.UserID; v != nil {
		builder.WriteString(", user_id=")
		builder.WriteString(*v)
	}
	if v := a.Kind; v != nil {
		builder.WriteString(", kind=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", unique=")
	builder.WriteString(fmt.Sprintf("%v", a.Unique))
	builder.WriteByte(')')
	return builder.String()
}

// Authenticators is a parsable slice of Authenticator.
type Authenticators []*Authenticator

func (a Authenticators) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}