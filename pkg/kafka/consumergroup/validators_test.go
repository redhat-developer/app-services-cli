package consumergroup

import (
	"testing"

	"github.com/redhat-developer/app-services-cli/pkg/localize/goi18n"
)

var validator *Validator

func init() {
	localizer, _ := goi18n.New(nil)

	validator = &Validator{
		Localizer: localizer,
	}
}

func TestValidateOffset(t *testing.T) {
	type args struct {
		offset string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should be valid when 'absolute' is provided",
			args: args{
				offset: "absolute",
			},
			wantErr: false,
		},
		{
			name: "Should be valid when 'latest' is provided",
			args: args{
				offset: "latest",
			},
			wantErr: false,
		},
		{
			name: "Should be valid when 'earliest' is provided",
			args: args{
				offset: "earliest",
			},
			wantErr: false,
		},
		{
			name: "Should be valid when 'timestamp' is provided",
			args: args{
				offset: "timestamp",
			},
			wantErr: false,
		},
		{
			name: "Should throw error when a valid offset type is capitalized",
			args: args{
				offset: "Timestamp",
			},
			wantErr: true,
		},
		{
			name: "Should throw error when a valid offset type is contains uppercase letters",
			args: args{
				offset: "offsEt",
			},
			wantErr: true,
		},
		{
			name: "Should throw error when an invalid value for offset type is provided",
			args: args{
				offset: "random",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// nolint
			if err := validator.ValidateOffset(tt.args.offset); (err != nil) != tt.wantErr {
				t.Errorf("ValidateName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateOffsetValue(t *testing.T) {
	type args struct {
		offset string
		value  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should be valid when value for absolute offset is an integer",
			args: args{
				offset: "absolute",
				value:  "1",
			},
			wantErr: false,
		},
		{
			name: "Should throw error when value for absolute offset is a string",
			args: args{
				offset: "absolute",
				value:  "random-value",
			},
			wantErr: true,
		},
		{
			name: "Should throw error when value for absolute offset contains special characters",
			args: args{
				offset: "absolute",
				value:  "random-value-*",
			},
			wantErr: true,
		},
		{
			name: "Should throw error when value for absolute offset contains numerics and alphabets",
			args: args{
				offset: "absolute",
				value:  "1rt2",
			},
			wantErr: true,
		},
		{
			name: "Should be valid when value for timestamp offset is in format 'yyyy-MM-dd'T'HH:mm:ssz'",
			args: args{
				offset: "timestamp",
				value:  "2016-06-23T09:07:21-07:00",
			},
			wantErr: false,
		},
		{
			name: "Should throw an error when value for timestamp offset is not in ISO 8601 format",
			args: args{
				offset: "timestamp",
				value:  "2016-06-23TZ09:07:21",
			},
			wantErr: true,
		},
		{
			name: "Should throw an error when value for timestamp offset is a string",
			args: args{
				offset: "timestamp",
				value:  "random-string-value",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// nolint
			if err := validator.ValidateOffsetValue(tt.args.offset, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ValidateName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
