package surveyjson

import (
	"encoding/json"
	"fmt"

	"github.com/iancoleman/orderedmap"
	"github.com/jackdelahunt/survey-json-schema/pkg/surveyjson/util"
)

// Based of https://www.ietf.org/archive/id/draft-handrews-json-schema-validation-01.txt
type JSONSchemaType struct {
	Version              string                     `json:"$schema,omitempty"`
	Ref                  string                     `json:"$ref,omitempty"`
	MultipleOf           *float64                   `json:"multipleOf,omitempty"`
	Maximum              *float64                   `json:"maximum,omitempty"`
	ExclusiveMaximum     *float64                   `json:"exclusiveMaximum,omitempty"`
	Minimum              *float64                   `json:"minimum,omitempty"`
	ExclusiveMinimum     *float64                   `json:"exclusiveMinimum,omitempty"`
	MaxLength            *int                       `json:"maxLength,omitempty"`
	MinLength            *int                       `json:"minLength,omitempty"`
	Pattern              *string                    `json:"pattern,omitempty"`
	AdditionalItems      *JSONSchemaType            `json:"additionalItems,omitempty"`
	Items                Items                      `json:"items,omitempty"`
	MaxItems             *int                       `json:"maxItems,omitempty"`
	MinItems             *int                       `json:"minItems,omitempty"`
	UniqueItems          bool                       `json:"uniqueItems,omitempty"`
	MaxProperties        *int                       `json:"maxProperties,omitempty"`
	MinProperties        *int                       `json:"minProperties,omitempty"`
	Required             []string                   `json:"required,omitempty"`
	Properties           *Properties                `json:"properties,omitempty"`
	PatternProperties    map[string]*JSONSchemaType `json:"patternProperties,omitempty"`
	AdditionalProperties *interface{}               `json:"additionalProperties,omitempty"`
	Dependencies         map[string]Dependency      `json:"dependencies,omitempty"`
	PropertyNames        *JSONSchemaType            `json:"propertyNames,omitempty"`
	Enum                 []interface{}              `json:"enum,omitempty"`
	Type                 string                     `json:"type,omitempty"`
	If                   *JSONSchemaType            `json:"if,omitempty"`
	Then                 *JSONSchemaType            `json:"then,omitempty"`
	Else                 *JSONSchemaType            `json:"else,omitempty"`

	AllOf []*JSONSchemaType `json:"allOf,omitempty"`
	AnyOf []*JSONSchemaType `json:"anyOf,omitempty"`

	OneOf []*JSONSchemaType `json:"oneOf,omitempty"`

	Not *JSONSchemaType `json:"not,omitempty"`

	Definitions      Definitions `json:"$defs,omitempty"`
	DefinitionsAlias Definitions `json:"definitions,omitempty"`

	Contains         *JSONSchemaType `json:"contains,omitempty"`
	Const            *interface{}    `json:"const,omitempty"`
	Title            string          `json:"title,omitempty"`
	Description      string          `json:"description,omitempty"`
	Default          interface{}     `json:"default,omitempty"`
	Format           *string         `json:"format,omitempty"`
	ContentMediaType *string         `json:"contentMediaType,omitempty"`
	ContentEncoding  *string         `json:"contentEncoding,omitempty"`
}

// Definitions hold schema definitions.
type Definitions map[string]*interface{}

// Dependency is either a Type or an array of strings, and so requires special unmarshaling from JSON
type Dependency struct {
	Type  *JSONSchemaType `json:-`
	Array []string        `json:-`
}

// UnmarshalJSON performs unmarshals Dependency from JSON, required as the json field can be one of two types
func (d *Dependency) UnmarshalJSON(b []byte) error {
	if b[0] == '[' {
		return json.Unmarshal(b, &d.Array)
	}
	return json.Unmarshal(b, d.Type)
}

// Items is a either a Type or a array of types, and so requires special unmarshaling from JSON
type Items struct {
	Types []*JSONSchemaType `json:-`
	Type  *JSONSchemaType   `json:-`
}

// Properties is a set of ordered key-value pairs, as it is ordered it requires special marshaling to/from JSON
type Properties struct {
	*orderedmap.OrderedMap
}

// UnmarshalJSON performs custom Unmarshaling for Properties allowing us to preserve order,
// which is not a standard JSON feature
func (p *Properties) UnmarshalJSON(b []byte) error {
	m := orderedmap.New()
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	if p.OrderedMap == nil {
		p.OrderedMap = orderedmap.New()
	}
	for _, k := range m.Keys() {
		v, _ := m.Get(k)
		t := JSONSchemaType{}
		om, ok := v.(orderedmap.OrderedMap)
		if !ok {
			return fmt.Errorf("Unable to cast nested data structure to OrderedMap")
		}
		values := make(map[string]interface{}, 0)
		for _, k1 := range om.Keys() {
			v1, _ := om.Get(k1)
			values[k1] = v1
		}
		err := util.ToStructFromMapStringInterface(values, &t)
		if err != nil {
			return err
		}
		p.Set(k, &t)
	}
	return nil
}

// UnmarshalJSON performs unmarshals Items from JSON, required as the json field can be one of two types
func (t *Items) UnmarshalJSON(b []byte) error {
	if b[0] == '[' {
		return json.Unmarshal(b, &t.Types)
	}
	err := json.Unmarshal(b, &t.Type)
	if err != nil {
		return err
	}
	return nil
}
