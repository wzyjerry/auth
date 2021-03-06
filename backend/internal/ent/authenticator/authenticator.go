// Code generated by entc, DO NOT EDIT.

package authenticator

const (
	// Label holds the string label denoting the authenticator type in the database.
	Label = "authenticator"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldKind holds the string denoting the kind field in the database.
	FieldKind = "kind"
	// FieldUnique holds the string denoting the unique field in the database.
	FieldUnique = "unique"
	// Table holds the table name of the authenticator in the database.
	Table = "authenticator"
)

// Columns holds all SQL columns for authenticator fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldKind,
	FieldUnique,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
