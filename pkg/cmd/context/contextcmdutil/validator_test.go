package contextcmdutil

import (
	"testing"

	"github.com/redhat-developer/app-services-cli/pkg/core/localize/goi18n"
)

var validator *Validator

func init() {
	localizer, _ := goi18n.New(nil)

	validator = &Validator{
		Localizer: localizer,
	}
}

func TestValidateName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should throw error when name is null",
			args: args{
				name: "",
			},
			wantErr: true,
		},
		{
			name: "Should pass when name contains alphanumeric characters",
			args: args{
				name: "context1",
			},
			wantErr: false,
		},
		{
			name: "Should pass when name contains alphanumeric characters and hyphens",
			args: args{
				name: "context-1",
			},
			wantErr: false,
		},
		{
			name: "Should throw error if name contains special characters",
			args: args{
				name: "context**1",
			},
			wantErr: true,
		},
		{
			name: "Should throw error if name contains trailing hyphen",
			args: args{
				name: "context-1-",
			},
			wantErr: true,
		},
		{
			name: "Should throw error if name contains leading hyphen",
			args: args{
				name: "-context-1",
			},
			wantErr: true,
		},
		{
			name: "Should throw error if name contains white spaces",
			args: args{
				name: "context  1",
			},
			wantErr: true,
		},
		{
			name: "Should throw error if name contains underscore",
			args: args{
				name: "context_1",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// nolint
			if err := validator.ValidateName(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("ValidateName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
