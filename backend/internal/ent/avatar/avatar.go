// Code generated by entc, DO NOT EDIT.

package avatar

const (
	// Label holds the string label denoting the avatar type in the database.
	Label = "avatar"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// Table holds the table name of the avatar in the database.
	Table = "avatar"
)

// Columns holds all SQL columns for avatar fields.
var Columns = []string{
	FieldID,
	FieldAvatar,
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