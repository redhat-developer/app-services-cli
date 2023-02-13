package surveyjson

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/jackdelahunt/survey-json-schema/pkg/surveyjson/util"

	"github.com/pkg/errors"

	"github.com/iancoleman/orderedmap"
)

const (
	RefPathPrefixDefs        = "#/$defs/"
	RefPathPrefixDefinitions = "#/definitions/"
)

// JSONSchemaOptions are options for generating values from a schema
type JSONSchemaOptions struct {
	// If there are existingValues then those questions will be
	// ignored and the existing value used unless askExisting is true
	AskExisting bool
	// If AutoAcceptDefaults is true, then default values will be used automatically.
	AutoAcceptDefaults bool
	NoAsk              bool
	// If IgnoreMissingValues is false then any values which don't have an existing value
	// (or a default value if autoAcceptDefaults is true) will cause an error
	IgnoreMissingValues bool
	In                  terminal.FileReader
	Out                 terminal.FileWriter
	OutErr              io.Writer
	Overrides           map[string]func(o *JSONSchemaOptions, ctx SchemaContext) error
}

type SchemaContext struct {
	Name                 string
	Prefixes             []string
	RequiredFields       []string
	SchemaType           *JSONSchemaType
	ParentType           *JSONSchemaType
	Output               *orderedmap.OrderedMap
	AdditionalValidators []survey.Validator
	ExistingValues       map[string]interface{}
	Definitions          *map[string]*interface{}
	Required             bool
}

// GenerateValues examines the schema in schemaBytes, asks a series of questions using in, out and outErr,
func (o *JSONSchemaOptions) GenerateValues(schemaBytes []byte, existingValues map[string]interface{}) ([]byte, error) {
	t := JSONSchemaType{}
	err := json.Unmarshal(schemaBytes, &t)
	if err != nil {
		return nil, errors.Wrapf(err, "unmarshaling schema %s", schemaBytes)
	}

	definitions := make(map[string]*interface{})

	ctx := SchemaContext{
		Name:                 "",
		Prefixes:             make([]string, 0),
		RequiredFields:       make([]string, 0),
		SchemaType:           &t,
		ParentType:           nil,
		Output:               orderedmap.New(),
		AdditionalValidators: make([]survey.Validator, 0),
		ExistingValues:       existingValues,
		Definitions:          &definitions,
		Required:             false,
	}

	err = o.Recurse(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// move the output up a level
	if root, ok := ctx.Output.Get(""); ok {
		bytes, err := json.Marshal(root)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return bytes, nil
	}
	return make([]byte, 0), fmt.Errorf("unable to find root element in %v", ctx.Output)

}

func (o *JSONSchemaOptions) Recurse(ctx SchemaContext) error {

	if f, ok := o.Overrides[ctx.Name]; ok {
		return f(o, ctx)
	}

	ctx.Required = util.Contains(ctx.RequiredFields, ctx.Name)
	if ctx.Name != "" {
		ctx.Prefixes = append(ctx.Prefixes, ctx.Name)
	}
	if ctx.SchemaType.ContentEncoding != nil {
		return fmt.Errorf("contentEncoding is not supported for %s", ctx.Name)
	}
	if ctx.SchemaType.ContentMediaType != nil {
		return fmt.Errorf("contentMediaType is not supported for %s", ctx.Name)
	}

	if len(ctx.SchemaType.Definitions) > 0 {
		for key, schema := range ctx.SchemaType.Definitions {
			(*ctx.Definitions)[key] = schema
		}
	}

	if len(ctx.SchemaType.DefinitionsAlias) > 0 {
		for key, schema := range ctx.SchemaType.DefinitionsAlias {
			(*ctx.Definitions)[key] = schema
		}
	}

	var err error

	switch ctx.SchemaType.Type {
	case "null":
		err = o.RecurseNull(ctx)
	case "boolean":
		err = o.RecurseBoolean(ctx)
	case "object":
		err = o.RecurseObject(ctx)
	case "array":
		err = o.RecurseArray(ctx)
	case "number":
		err = o.RecurseNumber(ctx)
	case "string":
		err = o.RecurseString(ctx)
	case "integer":
		err = o.RecurseInteger(ctx)
	default:
		if len(ctx.SchemaType.OneOf) != 0 {
			err = o.RecurseOneOf(ctx)
		}
	}

	if err != nil {
		return err
	}

	if ctx.SchemaType.Ref != "" {
		refPath, err := parseRefPath(ctx.SchemaType.Ref)
		if err != nil {
			return err
		}

		currentObject := (*(*ctx.Definitions)[refPath[0]]).(map[string]interface{})
		for i := 1; i < len(refPath); i += 1 {
			object, ok := currentObject[refPath[i]]
			if !ok {
				return errors.New(fmt.Sprintf("Could not resolve ref path \"%v\" is not a key in the object", refPath[i]))
			}

			currentObject = object.(map[string]interface{})
		}

		var mainDefinition *JSONSchemaType

		nestedJSON, err := json.Marshal(currentObject)
		if err != nil {
			return errors.New(fmt.Sprintf("Cannot marshal to json object %v", currentObject))
		}

		err = json.Unmarshal(nestedJSON, &mainDefinition)
		if err != nil {
			return errors.New(fmt.Sprintf("Cannot unmarshal json object to JsonSchemaObject %v", nestedJSON))
		}

		subContext := SchemaContext{
			Name:                 ctx.Name,
			Prefixes:             make([]string, 0),
			RequiredFields:       make([]string, 0),
			SchemaType:           mainDefinition,
			ParentType:           nil,
			Output:               ctx.Output,
			AdditionalValidators: make([]survey.Validator, 0),
			ExistingValues:       ctx.ExistingValues,
			Definitions:          ctx.Definitions,
		}
		err = o.Recurse(subContext)
		if err != nil {
			return err
		}
	}

	ctx.RequiredFields = ctx.SchemaType.Required
	err = o.handleConditionals(ctx)
	return err
}

func (o *JSONSchemaOptions) RecurseArray(ctx SchemaContext) error {
	if ctx.SchemaType.Const != nil {
		return fmt.Errorf("const is not supported for %s", ctx.Name)
		// TODO support const
	}
	if ctx.SchemaType.Contains != nil {
		return fmt.Errorf("contains is not supported for %s", ctx.Name)
		// TODO support contains
	}
	if ctx.SchemaType.AdditionalItems != nil {
		return fmt.Errorf("additionalItems is not supported for %s", ctx.Name)
		// TODO support additonalItems
	}
	err := o.handleArrayProperty(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (o *JSONSchemaOptions) RecurseObject(ctx SchemaContext) error {
	if len(ctx.SchemaType.PatternProperties) > 0 {
		return fmt.Errorf("patternProperties is not supported for %s", ctx.Name)
	}
	if len(ctx.SchemaType.Dependencies) > 0 {
		return fmt.Errorf("dependencies is not supported for %s", ctx.Name)
	}
	if ctx.SchemaType.PropertyNames != nil {
		return fmt.Errorf("propertyNames is not supported for %s", ctx.Name)
	}
	if ctx.SchemaType.Const != nil {
		return fmt.Errorf("const is not supported for %s", ctx.Name)
		// TODO support const
	}
	if ctx.SchemaType.Properties != nil {
		for valid := false; !valid; {
			result := orderedmap.New()
			duringValidators := make([]survey.Validator, 0)
			postValidators := []survey.Validator{
				// These validators are run after the processing of the properties
				MinPropertiesValidator(ctx.SchemaType.MinProperties, result, ctx.Name),
				EnumValidator(ctx.SchemaType.Enum),
				MaxPropertiesValidator(ctx.SchemaType.MaxProperties, result, ctx.Name),
			}
			for _, n := range ctx.SchemaType.Properties.Keys() {
				v, _ := ctx.SchemaType.Properties.Get(n)
				property := v.(*JSONSchemaType)
				var nestedExistingValues map[string]interface{}
				if ctx.Name == "" {
					// This is the root element
					nestedExistingValues = ctx.ExistingValues
				} else if v, ok := ctx.ExistingValues[ctx.Name]; ok {
					var err error
					nestedExistingValues, err = util.AsMapOfStringsIntefaces(v)
					if err != nil {
						return errors.Wrapf(err, "converting key %s from %v to map[string]interface{}", ctx.Name, ctx.ExistingValues)
					}
				}

				subContext := SchemaContext{
					Name:                 n,
					Prefixes:             ctx.Prefixes,
					RequiredFields:       ctx.SchemaType.Required,
					ParentType:           ctx.SchemaType,
					SchemaType:           property,
					Output:               result,
					AdditionalValidators: duringValidators,
					ExistingValues:       nestedExistingValues,
					Definitions:          ctx.Definitions,
					Required:             false,
				}

				err := o.Recurse(subContext)
				if err != nil {
					return err
				}
			}
			valid = true
			for _, v := range postValidators {
				err := v(result)
				if err != nil {
					str := fmt.Sprintf("Sorry, your reply was invalid: %s", err.Error())
					_, err1 := o.Out.Write([]byte(str))
					if err1 != nil {
						return err1
					}
					valid = false
				}
			}
			if valid {
				ctx.Output.Set(ctx.Name, result)
			}
		}
	} else {
		// if there are no properties then just insert an empty body.. "name": {}
		ctx.Output.Set(ctx.Name, make(map[string]*JSONSchemaType))
	}

	return nil
}

func (o *JSONSchemaOptions) RecurseString(ctx SchemaContext) error {
	validators := []survey.Validator{
		EnumValidator(ctx.SchemaType.Enum),
		MinLengthValidator(ctx.SchemaType.MinLength),
		MaxLengthValidator(ctx.SchemaType.MaxLength),
		RequiredValidator(ctx.Required),
		PatternValidator(ctx.SchemaType.Pattern),
	}
	// Defined Format validation
	if ctx.SchemaType.Format != nil {
		format := util.DereferenceString(ctx.SchemaType.Format)
		switch format {
		case "date-time":
			validators = append(validators, DateTimeValidator())
		case "date":
			validators = append(validators, DateValidator())
		case "time":
			validators = append(validators, TimeValidator())
		case "email", "idn-email":
			validators = append(validators, EmailValidator())
		case "hostname", "idn-hostname":
			validators = append(validators, HostnameValidator())
		case "ipv4":
			validators = append(validators, Ipv4Validator())
		case "ipv6":
			validators = append(validators, Ipv6Validator())
		case "uri":
			validators = append(validators, URIValidator())
		case "uri-reference":
			validators = append(validators, URIReferenceValidator())
		case "iri":
			return fmt.Errorf("iri defined format not supported")
		case "iri-reference":
			return fmt.Errorf("iri-reference defined format not supported")
		case "uri-template":
			return fmt.Errorf("uri-template defined format not supported")
		case "json-pointer":
			validators = append(validators, JSONPointerValidator())
		case "relative-json-pointer":
			return fmt.Errorf("relative-json-pointer defined format not supported")
		case "regex":
			return fmt.Errorf("regex defined format not supported, use pattern keyword")
		}
	}

	subContext := ctx
	subContext.AdditionalValidators = append(validators, ctx.AdditionalValidators...)

	err := o.handleBasicProperty(subContext, ctx.Required)
	if err != nil {
		return err
	}

	return nil
}

func (o *JSONSchemaOptions) RecurseNumber(ctx SchemaContext) error {
	validators := ctx.AdditionalValidators

	subContext := ctx
	subContext.AdditionalValidators = numberValidator(ctx.Required, append(validators, FloatValidator()), ctx.SchemaType)

	err := o.handleBasicProperty(ctx, ctx.Required)
	if err != nil {
		return err
	}
	return nil
}

func (o *JSONSchemaOptions) RecurseInteger(ctx SchemaContext) error {
	subContext := ctx
	subContext.AdditionalValidators = append(ctx.AdditionalValidators, IntegerValidator())

	err := o.handleBasicProperty(ctx, ctx.Required)
	if err != nil {
		return err
	}

	return nil
}

func (o *JSONSchemaOptions) RecurseOneOf(ctx SchemaContext) error {
	options := make([]string, len(ctx.SchemaType.OneOf))

	for i := 0; i < len(options); i++ {
		options[i] = ctx.SchemaType.OneOf[i].ToString()
	}

	prompt := &survey.Select{
		Message: fmt.Sprintf("Select one of these types to use for %v", ctx.Name),
		Options: options,
	}

	var answer int

	err := survey.AskOne(prompt, &answer)
	if err != nil {
		return err
	}

	result := orderedmap.New()

	subContext := SchemaContext{
		Name:                 strconv.Itoa(answer),
		Prefixes:             ctx.Prefixes,
		RequiredFields:       ctx.SchemaType.Required,
		ParentType:           ctx.SchemaType,
		SchemaType:           ctx.SchemaType.OneOf[answer],
		Output:               result,
		AdditionalValidators: ctx.AdditionalValidators,
		ExistingValues:       ctx.ExistingValues,
		Definitions:          ctx.Definitions,
		Required:             false,
	}

	err = o.Recurse(subContext)
	if err != nil {
		return err
	}

	if len(result.Keys()) > 0 {
		rootValue, ok := result.Get(strconv.Itoa(answer))
		if !ok {
			return errors.New("Cannot get root value from one of result")
		}

		ctx.Output.Set(ctx.Name, rootValue)
	} else {
		ctx.Output.Set(ctx.Name, make(map[string]*JSONSchemaType))
	}

	return nil
}

func (o *JSONSchemaOptions) RecurseNull(ctx SchemaContext) error {
	ctx.Output.Set(ctx.Name, nil)
	return nil
}

func (o *JSONSchemaOptions) RecurseBoolean(ctx SchemaContext) error {
	validators := []survey.Validator{
		EnumValidator(ctx.SchemaType.Enum),
		RequiredValidator(ctx.Required),
		BoolValidator(),
	}
	ctx.AdditionalValidators = append(validators, ctx.AdditionalValidators...)
	err := o.handleBasicProperty(ctx, ctx.Required)
	if err != nil {
		return err
	}
	return nil
}

func (o *JSONSchemaOptions) handleConditionals(ctx SchemaContext) error {
	if ctx.ParentType != nil {
		err := o.handleIf(ctx)
		if err != nil {
			return errors.WithStack(err)
		}
		err = o.handleAllOf(ctx)
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

func (o *JSONSchemaOptions) handleAllOf(ctx SchemaContext) error {
	if ctx.ParentType.AllOf != nil && len(ctx.ParentType.AllOf) > 0 {
		for _, allType := range ctx.ParentType.AllOf {
			ctx.ParentType = allType
			err := o.handleIf(ctx)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *JSONSchemaOptions) handleIf(ctx SchemaContext) error {
	if ctx.ParentType.If != nil {
		if len(ctx.ParentType.If.Properties.Keys()) > 1 {
			return fmt.Errorf("Please specify a single property condition when using If in your schema")
		}
		detypedCondition, conditionFound := ctx.ParentType.If.Properties.Get(ctx.Name)
		selectedValue, selectedValueFound := ctx.Output.Get(ctx.Name)
		if conditionFound && selectedValueFound {
			desiredState := true
			if detypedCondition != nil {
				condition := detypedCondition.(*JSONSchemaType)
				if condition.Const != nil {

					switch ctx.SchemaType.Type {
					case "boolean":
						tConst, err := util.AsBool(*condition.Const)
						if err != nil {
							return err
						}
						if tConst != selectedValue {
							desiredState = false
						}
					default:
						stringConst, err := util.AsString(*condition.Const)
						if err != nil {
							return errors.Wrapf(err, "converting %s to string", condition.Type)
						}
						typedConst, err := convertAnswer(stringConst, ctx.SchemaType.Type)
						if typedConst != selectedValue {
							desiredState = false
						}
					}
				}
			}
			result := orderedmap.New()
			if desiredState {
				if ctx.ParentType.Then != nil {
					ctx.ParentType.Then.Type = "object"
					err := o.processThenElse(result, ctx.ParentType.Then, ctx)
					if err != nil {
						return err
					}
				}
			} else {
				if ctx.ParentType.Else != nil {
					ctx.ParentType.Else.Type = "object"
					err := o.processThenElse(result, ctx.ParentType.Else, ctx)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

func (o *JSONSchemaOptions) processThenElse(result *orderedmap.OrderedMap, conditionalType *JSONSchemaType, ctx SchemaContext) error {

	subContext := ctx
	subContext.SchemaType = conditionalType
	subContext.Output = result

	err := o.Recurse(subContext)
	if err != nil {
		return err
	}
	resultSet, found := result.Get("")
	if found {
		resultMap := resultSet.(*orderedmap.OrderedMap)
		for _, key := range resultMap.Keys() {
			value, foundValue := resultMap.Get(key)
			if foundValue {
				ctx.Output.Set(key, value)
			}
		}
	}
	return nil
}

func parseRefPath(path string) ([]string, error) {

	if strings.HasPrefix(path, RefPathPrefixDefs) {
		path = strings.TrimPrefix(path, RefPathPrefixDefs)
	} else if strings.HasPrefix(path, RefPathPrefixDefinitions) {
		path = strings.TrimPrefix(path, RefPathPrefixDefinitions)
	} else {
		return nil, errors.New(fmt.Sprintf("Reference path must start with definition prefix: %v or %v", RefPathPrefixDefs, RefPathPrefixDefinitions))
	}

	return strings.Split(path, "/"), nil
}

// According to the spec, "An instance validates successfully against this keyword if its value
// is equal to the value of the keyword." which we interpret for questions as "this is the value of this keyword"
func (o *JSONSchemaOptions) handleConst(ctx SchemaContext) error {
	message := fmt.Sprintf("Set %s to %v", ctx.Name, *ctx.SchemaType.Const)
	if ctx.SchemaType.Title != "" {
		message = ctx.SchemaType.Title
	}
	// These are console output, not logging - DO NOT CHANGE THEM TO log statements
	fmt.Fprint(o.Out, message)
	if ctx.SchemaType.Description != "" {
		fmt.Fprint(o.Out, ctx.SchemaType.Description)
	}

	switch ctx.SchemaType.Type {
	case "boolean":
		tConst, err := util.AsBool(*ctx.SchemaType.Const)
		if err != nil {
			return err
		}
		ctx.Output.Set(ctx.Name, tConst)
	default:
		stringConst, err := util.AsString(*ctx.SchemaType.Const)
		if err != nil {
			return errors.Wrapf(err, "converting %s to string", *ctx.SchemaType.Const)
		}
		typedConst, err := convertAnswer(stringConst, ctx.SchemaType.Type)
		ctx.Output.Set(ctx.Name, typedConst)
	}

	return nil
}

func (o *JSONSchemaOptions) handleArrayProperty(ctx SchemaContext) error {
	results := make([]interface{}, 0)

	validators := []survey.Validator{
		MaxItemsValidator(ctx.SchemaType.MaxItems, results),
		UniqueItemsValidator(results),
		MinItemsValidator(ctx.SchemaType.MinItems, results),
		EnumValidator(ctx.SchemaType.Enum),
	}
	if ctx.SchemaType.Items.Type != nil && ctx.SchemaType.Items.Type.Enum != nil {
		// Arrays can used to create a multi-select list
		// Note that this only supports basic types at the moment
		if ctx.SchemaType.Items.Type.Type == "null" {
			ctx.Output.Set(ctx.Name, nil)
			return nil
		} else if !util.Contains([]string{"string", "boolean", "number", "integer"}, ctx.SchemaType.Items.Type.Type) {
			return fmt.Errorf("type %s is not supported for array %s", ctx.SchemaType.Items.Type.Type, ctx.Name)
			// TODO support other types
		}
		message := fmt.Sprintf("Select values for %s", ctx.Name)
		help := ""
		ask := true
		var defaultValue []string
		autoAcceptMessage := ""
		if value, ok := ctx.ExistingValues[ctx.Name]; ok {
			if !o.AskExisting {
				ask = false
			}
			existingString, err := util.AsString(value)
			existingArray, err1 := util.AsSliceOfStrings(value)
			if err != nil && err1 != nil {
				v := reflect.ValueOf(value)
				v = reflect.Indirect(v)
				return fmt.Errorf("Cannot convert %v (%v) to string or []string", v.Type(), value)
			}
			if existingString != "" {
				defaultValue = []string{
					existingString,
				}
			} else {
				defaultValue = existingArray
			}
			autoAcceptMessage = "Automatically accepted existing value"
		} else if ctx.SchemaType.Default != nil {
			if o.AutoAcceptDefaults {
				ask = false
				autoAcceptMessage = "Automatically accepted default value"
			}
			defaultString, err := util.AsString(ctx.SchemaType.Default)
			defaultArray, err1 := util.AsSliceOfStrings(ctx.SchemaType.Default)
			if err != nil && err1 != nil {
				v := reflect.ValueOf(ctx.SchemaType.Default)
				v = reflect.Indirect(v)
				return fmt.Errorf("Cannot convert %value (%value) to string or []string", v.Type(), ctx.SchemaType.Default)
			}
			if defaultString != "" {
				defaultValue = []string{
					defaultString,
				}
			} else {
				defaultValue = defaultArray
			}
		}
		if o.NoAsk {
			ask = false
		}

		options := make([]string, 0)
		if ctx.SchemaType.Title != "" {
			message = ctx.SchemaType.Title
		}
		if ctx.SchemaType.Description != "" {
			help = ctx.SchemaType.Description
		}
		for _, e := range ctx.SchemaType.Items.Type.Enum {
			options = append(options, fmt.Sprintf("%v", e))
		}

		answer := make([]string, 0)
		surveyOpts := survey.WithStdio(o.In, o.Out, o.OutErr)
		validator := survey.ComposeValidators(validators...)

		if ask {
			prompt := &survey.MultiSelect{
				Default: defaultValue,
				Help:    help,
				Message: message,
				Options: options,
			}
			err := survey.AskOne(prompt, &answer, survey.WithValidator(validator), surveyOpts)
			if err != nil {
				return err
			}
		} else {
			answer = defaultValue
			msg := fmt.Sprintf("%s %s [%s]\n", message, util.ColorInfo(answer), autoAcceptMessage)
			_, err := fmt.Fprint(terminal.NewAnsiStdout(o.Out), msg)
			if err != nil {
				return errors.Wrapf(err, "writing %s to console", msg)
			}
		}

		for _, a := range answer {
			v, err := convertAnswer(a, ctx.SchemaType.Items.Type.Type)
			// An error is a genuine error as we've already done type validation
			if err != nil {
				return err
			}
			results = append(results, v)
		}
	}

	ctx.Output.Set(ctx.Name, results)
	return nil
}

func convertAnswer(answer string, t string) (interface{}, error) {
	if t == "number" {
		return strconv.ParseFloat(answer, 64)
	} else if t == "integer" {
		return strconv.Atoi(answer)
	} else if t == "boolean" {
		return strconv.ParseBool(answer)
	} else {
		return answer, nil
	}
}

func (o *JSONSchemaOptions) handleBasicProperty(ctx SchemaContext, required bool) error {
	if ctx.SchemaType.Const != nil {
		return o.handleConst(ctx)
	}

	ask := true
	defaultValue := ""
	autoAcceptMessage := ""
	if v, ok := ctx.ExistingValues[ctx.Name]; ok {
		if !o.AskExisting {
			ask = false
		}
		defaultValue = fmt.Sprintf("%v", v)
		autoAcceptMessage = "Automatically accepted existing value"
	} else if ctx.SchemaType.Default != nil {
		if o.AutoAcceptDefaults {
			ask = false
			autoAcceptMessage = "Automatically accepted default value"
		}
		defaultValue = fmt.Sprintf("%v", ctx.SchemaType.Default)
	}
	if o.NoAsk {
		ask = false
	}

	var result interface{}
	message := fmt.Sprintf("Enter a value for %s", ctx.Name)
	help := ""
	if ctx.SchemaType.Title != "" {
		message = ctx.SchemaType.Title
	}
	if ctx.SchemaType.Description != "" {
		help = ctx.SchemaType.Description
	}

	if !ask {
		envVar := strings.ToUpper("SURVEY_VALUE_" + strings.Join(ctx.Prefixes, "_"))
		if defaultValue == "" {
			defaultValue = os.Getenv(envVar)
			if defaultValue != "" {
				fmt.Fprintf(os.Stderr, "defaulting value from $%s\n", envVar)
			}
		}
		if !o.IgnoreMissingValues && defaultValue == "" {
			// lets not fail if in batch mode for non-required fields
			if !o.NoAsk || required {
				return fmt.Errorf("no existing or default value in answer to question %s and no value for $%s", message, envVar)
			}
		}
	}

	surveyOpts := survey.WithStdio(o.In, o.Out, o.OutErr)
	validator := survey.ComposeValidators(ctx.AdditionalValidators...)
	// Ask the question
	// Custom format support for passwords
	dereferencedFormat := strings.TrimSuffix(util.DereferenceString(ctx.SchemaType.Format), "-passthrough")
	if dereferencedFormat == "password" || dereferencedFormat == "token" {
		// the default value for a password is just the path, so clear those values
		if _, ok := ctx.ExistingValues[ctx.Name]; ok {
			defaultValue = ""
			ask = true
		}

		secret, err := handlePasswordProperty(message, help, dereferencedFormat, ask, validator, surveyOpts, defaultValue,
			autoAcceptMessage, o.Out, ctx.SchemaType.Type)
		if err != nil {
			return errors.WithStack(err)
		}
		if secret != nil {
			value, err := util.AsString(secret)
			if err != nil {
				return err
			}
			// TODO passwords etc. should be stored in a secret store instead
			ctx.Output.Set(ctx.Name, value)
		}
	} else if ctx.SchemaType.Enum != nil {
		var enumResult string
		// Support for selects
		names := make([]string, 0)
		for _, e := range ctx.SchemaType.Enum {
			names = append(names, fmt.Sprintf("%v", e))
		}
		prompt := &survey.Select{
			Message: message,
			Options: names,
			Default: defaultValue,
			Help:    help,
		}
		if ask {
			err := survey.AskOne(prompt, &enumResult, survey.WithValidator(validator), surveyOpts)
			if err != nil {
				return err
			}
			result = enumResult
		} else {
			result = defaultValue
			msg := fmt.Sprintf("%s %s [%s]\n", message, util.ColorInfo(result), autoAcceptMessage)
			_, err := fmt.Fprint(terminal.NewAnsiStdout(o.Out), msg)
			if err != nil {
				return errors.Wrapf(err, "writing %s to console", msg)
			}
		}
	} else if ctx.SchemaType.Type == "boolean" {
		// Confirm dialog
		var d bool
		var err error
		if defaultValue != "" {
			d, err = strconv.ParseBool(defaultValue)
			if err != nil {
				return err
			}
		}

		var answer bool
		prompt := &survey.Confirm{
			Message: message,
			Help:    help,
			Default: d,
		}

		if ask {
			err = survey.AskOne(prompt, &answer, survey.WithValidator(validator), surveyOpts)
			if err != nil {
				return errors.Wrapf(err, "error asking user %s using validators %v", message, ctx.AdditionalValidators)
			}
		} else {
			answer = d
			msg := fmt.Sprintf("%s %s [%s]\n", message, util.ColorInfo(answer), autoAcceptMessage)
			_, err := fmt.Fprint(terminal.NewAnsiStdout(o.Out), msg)
			if err != nil {
				return errors.Wrapf(err, "writing %s to console", msg)
			}
		}
		result = answer
	} else {
		// Basic input
		prompt := &survey.Input{
			Message: message,
			Default: defaultValue,
			Help:    help,
		}
		var answer string
		var err error
		if ask {
			err = survey.AskOne(prompt, &answer, survey.WithValidator(validator), surveyOpts)
			if err != nil {
				return errors.Wrapf(err, "error asking user %s using validators %v", message, ctx.AdditionalValidators)
			}
		} else {
			answer = defaultValue
			msg := fmt.Sprintf("%s %s [%s]\n", message, util.ColorInfo(answer), autoAcceptMessage)
			_, err := fmt.Fprint(terminal.NewAnsiStdout(o.Out), msg)
			if err != nil {
				return errors.Wrapf(err, "writing %s to console", msg)
			}
		}
		if answer != "" {
			result, err = convertAnswer(answer, ctx.SchemaType.Type)
		}
		if err != nil {
			return errors.Wrapf(err, "error converting result %s to type %s", answer, ctx.SchemaType.Type)
		}
	}

	if result != nil {
		// Write the value to the output
		ctx.Output.Set(ctx.Name, result)
	}
	return nil
}

func handlePasswordProperty(message string, help string, kind string, ask bool, validator survey.Validator,
	surveyOpts survey.AskOpt, defaultValue string, autoAcceptMessage string, out terminal.FileWriter,
	t string) (interface{}, error) {
	// Secret input
	prompt := &survey.Password{
		Message: message,
		Help:    help,
	}

	var answer string
	if ask {
		err := survey.AskOne(prompt, &answer, survey.WithValidator(validator), surveyOpts)
		if err != nil {
			return nil, err
		}
	} else {
		answer = defaultValue
		msg := fmt.Sprintf("%s *** [%s]\n", message, autoAcceptMessage)
		_, err := fmt.Fprint(terminal.NewAnsiStdout(out), msg)
		if err != nil {
			return nil, errors.Wrapf(err, "writing %s to console", msg)
		}
	}
	if answer != "" {
		result, err := convertAnswer(answer, t)
		if err != nil {
			return nil, errors.Wrapf(err, "error converting answer %s to type %s", answer, t)
		}
		return result, nil
	}
	return nil, nil
}

func numberValidator(required bool, additonalValidators []survey.Validator, t *JSONSchemaType) []survey.Validator {
	validators := []survey.Validator{EnumValidator(t.Enum),
		MultipleOfValidator(t.MultipleOf),
		RequiredValidator(required),
		MinValidator(t.Minimum, false),
		MinValidator(t.ExclusiveMinimum, true),
		MaxValidator(t.Maximum, false),
		MaxValidator(t.ExclusiveMaximum, true),
	}
	return append(validators, additonalValidators...)
}
