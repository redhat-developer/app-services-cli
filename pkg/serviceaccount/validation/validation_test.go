package validation

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

func TestValidateShortDescription(t *testing.T) {
	type args struct {
		val interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "fails when length is 0",
			args:    args{""},
			wantErr: true,
		},
		{
			name:    "passes when length is 1",
			args:    args{"s"},
			wantErr: false,
		},
		{
			name:    "passes when length is 50 (max length)",
			args:    args{"ssssssssssssssssssssssssssssssssssssssssssssssssss"},
			wantErr: false,
		},
		{
			name:    "fails when length exceeds max length",
			args:    args{"sssssssssssssssssssssssssssssssssssssssssssssssssss"},
			wantErr: true,
		},
		{
			name:    "passes on valid name",
			args:    args{"svcacctone"},
			wantErr: false,
		},
		{
			name:    "passes on valid name with hyphens",
			args:    args{"svc-acct-one"},
			wantErr: false,
		},
		{
			name:    "passes on valid name with digits",
			args:    args{"svc-acct-1s"},
			wantErr: false,
		},
		{
			name:    "fails with capital letters",
			args:    args{"Svc-acct-one"},
			wantErr: true,
		},
		{
			name:    "fails number in first section",
			args:    args{"1svc-acct-one"},
			wantErr: true,
		},
	}
	// nolint:scopelint
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validator.ValidateShortDescription(tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("ValidateShortDescription() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateUUID(t *testing.T) {
	type args struct {
		val interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "fails when length is 5",
			args:    args{"kafka"},
			wantErr: true,
		},
		{
			name:    "fails for empty string",
			args:    args{""},
			wantErr: true,
		},
		{
			name:    "fails for special chars",
			args:    args{"9e4d1b1f-19d*-47c2-a334-e420c5e5bbce"},
			wantErr: true,
		},
		{
			name:    "passes for valid UUID",
			args:    args{"9e4d1b1f-19dd-47c2-a334-e420c5e5bbce"},
			wantErr: false,
		},
		{
			name:    "passes for numeric UUID",
			args:    args{"11111111-2222-3333-4444-555555555555"},
			wantErr: false,
		},
		{
			name:    "fails for ID containing capital letters",
			args:    args{"9e4d1b1f-19dd-47c2-A334-e420c5e5bbce"},
			wantErr: true,
		},
	}

	// nolint:scopelint
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validator.ValidateUUID(tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("ValidateUUID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
