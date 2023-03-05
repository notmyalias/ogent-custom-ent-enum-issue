package schema

import (
	"database/sql/driver"
	"fmt"

	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Enum("item_type_example_1").
			GoType(ItemType(0)),
		field.Enum("item_type_example_2").
			GoType(ItemType(0)).
			Annotations(
				entoas.Schema(ogen.Int()),
			),
		field.Enum("item_type_example_3").
			GoType(StructItemType{}),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return nil
}

type ItemType int

const (
	Physical ItemType = iota + 1
	Digital
	Service
)

func (t ItemType) String() string {
	switch t {
	case Physical:
		return "PHYSICAL"

	case Digital:
		return "DIGITAL"

	case Service:
		return "SERVICE"

	default:
		return ""
	}
}

// Values implements the field.EnumValues interface
func (ItemType) Values() []string {
	return []string{
		Physical.String(),
		Digital.String(),
		Service.String(),
	}
}

func (t ItemType) Value() (driver.Value, error) {
	if s := t.String(); s != "" {
		return s, nil
	}
	return nil, fmt.Errorf("unknown ItemType: %d", t)
}

func (t *ItemType) Scan(val any) error {
	var s string

	switch v := val.(type) {
	case nil:
		return nil

	case []byte:
		s = string(v)

	case string:
		s = v

	default:
		return fmt.Errorf("unexpected type %T", v)
	}

	var err error
	*t, err = itemTypeFromStr(s)
	if err != nil {
		return err
	}

	return nil
}

func itemTypeFromStr(s string) (ItemType, error) {
	switch s {
	case Physical.String():
		return Physical, nil

	case Digital.String():
		return Digital, nil

	case Service.String():
		return Service, nil

	default:
		return ItemType(0), fmt.Errorf("invalid ItemType value: %s", s)
	}
}

type StructItemType struct {
	name  string
	value int
}

var (
	StructUnknown = StructItemType{
		name:  "UNKNOWN",
		value: 0,
	}

	StructPhysical = StructItemType{
		name:  "PHYSICAL",
		value: 1,
	}

	StructDigital = StructItemType{
		name:  "DIGITAL",
		value: 2,
	}

	StructService = StructItemType{
		name:  "SERVICE",
		value: 3,
	}
)

func (StructItemType) Values() []string {
	return []string{
		StructPhysical.String(),
		StructDigital.String(),
		StructService.String(),
	}
}

func (t StructItemType) Value() (driver.Value, error) {
	if t == StructUnknown {
		return nil, fmt.Errorf("unknown StructItemType: %s", t)
	}
	return t.String(), nil
}

func (t *StructItemType) Scan(val any) error {
	var s string

	switch v := val.(type) {
	case nil:
		return nil

	case []byte:
		s = string(v)

	case string:
		s = v

	default:
		return fmt.Errorf("unexpected type %T", v)
	}

	var err error
	*t, err = structItemTypeFromStr(s)
	if err != nil {
		return err
	}

	return nil
}

func (t StructItemType) String() string {
	return t.name
}

func structItemTypeFromStr(s string) (StructItemType, error) {
	switch s {
	case StructPhysical.String():
		return StructPhysical, nil

	case StructDigital.String():
		return StructDigital, nil

	case StructService.String():
		return StructService, nil

	default:
		return StructUnknown, fmt.Errorf("invalid ItemType value: %s", s)
	}
}
