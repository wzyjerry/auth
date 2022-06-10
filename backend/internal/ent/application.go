// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/wzyjerry/auth/internal/ent/application"
	"github.com/wzyjerry/auth/internal/ent/schema/applicationNested"
)

// Application is the model entity for the Application schema.
type Application struct {
	config `json:"-"`
	// ID of the ent.
	// 主键
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	// 应用名称
	Name *string `json:"name,omitempty"`
	// Homepage holds the value of the "homepage" field.
	// 应用主页
	Homepage *string `json:"homepage,omitempty"`
	// Description holds the value of the "description" field.
	// 应用简介
	Description *string `json:"description,omitempty"`
	// Callback holds the value of the "callback" field.
	// 授权回调地址
	Callback *string `json:"callback,omitempty"`
	// Admin holds the value of the "admin" field.
	// 管理员ID
	Admin *string `json:"admin,omitempty"`
	// ClientID holds the value of the "client_id" field.
	// 应用ID
	ClientID *string `json:"client_id,omitempty"`
	// ClientSecrets holds the value of the "client_secrets" field.
	// 应用密钥
	ClientSecrets []*applicationNested.ClientSecret `json:"client_secrets,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Application) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case application.FieldClientSecrets:
			values[i] = new([]byte)
		case application.FieldID, application.FieldName, application.FieldHomepage, application.FieldDescription, application.FieldCallback, application.FieldAdmin, application.FieldClientID:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Application", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Application fields.
func (a *Application) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case application.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				a.ID = value.String
			}
		case application.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = new(string)
				*a.Name = value.String
			}
		case application.FieldHomepage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field homepage", values[i])
			} else if value.Valid {
				a.Homepage = new(string)
				*a.Homepage = value.String
			}
		case application.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				a.Description = new(string)
				*a.Description = value.String
			}
		case application.FieldCallback:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field callback", values[i])
			} else if value.Valid {
				a.Callback = new(string)
				*a.Callback = value.String
			}
		case application.FieldAdmin:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field admin", values[i])
			} else if value.Valid {
				a.Admin = new(string)
				*a.Admin = value.String
			}
		case application.FieldClientID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field client_id", values[i])
			} else if value.Valid {
				a.ClientID = new(string)
				*a.ClientID = value.String
			}
		case application.FieldClientSecrets:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field client_secrets", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &a.ClientSecrets); err != nil {
					return fmt.Errorf("unmarshal field client_secrets: %w", err)
				}
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Application.
// Note that you need to call Application.Unwrap() before calling this method if this Application
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Application) Update() *ApplicationUpdateOne {
	return (&ApplicationClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Application entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Application) Unwrap() *Application {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Application is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Application) String() string {
	var builder strings.Builder
	builder.WriteString("Application(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	if v := a.Name; v != nil {
		builder.WriteString(", name=")
		builder.WriteString(*v)
	}
	if v := a.Homepage; v != nil {
		builder.WriteString(", homepage=")
		builder.WriteString(*v)
	}
	if v := a.Description; v != nil {
		builder.WriteString(", description=")
		builder.WriteString(*v)
	}
	if v := a.Callback; v != nil {
		builder.WriteString(", callback=")
		builder.WriteString(*v)
	}
	if v := a.Admin; v != nil {
		builder.WriteString(", admin=")
		builder.WriteString(*v)
	}
	if v := a.ClientID; v != nil {
		builder.WriteString(", client_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", client_secrets=")
	builder.WriteString(fmt.Sprintf("%v", a.ClientSecrets))
	builder.WriteByte(')')
	return builder.String()
}

// Applications is a parsable slice of Application.
type Applications []*Application

func (a Applications) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
