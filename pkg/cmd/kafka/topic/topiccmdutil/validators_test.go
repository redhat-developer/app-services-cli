package topiccmdutil

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

// nolint:funlen
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
			name: "Should throw error when exceeds 249 characters",
			args: args{
				name: "verylongkafkanamef8d9dkf9dkc11dfsverylongkafkanamef8d9dkf9dkc11dfsverylongkafkanamef8d9dkf9dkc11dfsverylongkafkanamef8d9dkf9dkc11dfsverylongkafkanamef8d9dkf9dkc11dfsverylongkafkanamef8d9dkf9dkc11dfsverylongkafkanamef8d9dkf9dkc11dfsverylongkafkanamef8d9dkf9dkc11dfsverylongkafkanamef8d9dkf9dkc11dfsverylongkafkanamef8d9dkf9dkc11dfs",
			},
			wantErr: true,
		},
		{
			name: "Should be valid when name is exactly 32 characters",
			args: args{
				name: "verylongtopicnamef8d9dkf9dkc11dd",
			},
			wantErr: false,
		},
		{
			name: "Should be valid when containing dots",
			args: args{
				name: "k.afk.a",
			},
			wantErr: false,
		},
		{
			name: "Should be invalid when containing special characters",
			args: args{
				name: "r*maujj??",
			},
			wantErr: true,
		},
		{
			name: "Should be invalid when containing commas",
			args: args{
				name: "kaf,ka",
			},
			wantErr: true,
		},
		{
			name: "Should be invalid when name is exactly 0 characters",
			args: args{
				name: "",
			},
			wantErr: true,
		},
		{
			name: "Should be valid when using hyphens",
			args: args{
				name: "my-kafka-topic",
			},
			wantErr: false,
		},
		{
			name: "Should be valid when starts with number",
			args: args{
				name: "1my-kafka-instance",
			},
			wantErr: false,
		},
		{
			name: "Should be invalid when only two dots",
			args: args{
				name: "..",
			},
			wantErr: true,
		},
		{
			name: "Should be invalid when only one dot",
			args: args{
				name: ".",
			},
			wantErr: true,
		},
		{
			name: "Should be valid when only three dots",
			args: args{
				name: "...",
			},
			wantErr: false,
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

func TestValidateSearchInput(t *testing.T) {
	type args struct {
		search string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should be valid when containing dots",
			args: args{
				search: "k.afk.a",
			},
			wantErr: false,
		},
		{
			name: "Should be invalid when containing special characters",
			args: args{
				search: "r*maujj??",
			},
			wantErr: true,
		},
		{
			name: "Should be invalid when containing commas",
			args: args{
				search: "kaf,ka",
			},
			wantErr: true,
		},
		{
			name: "Should be invalid when length is zero",
			args: args{
				search: "",
			},
			wantErr: true,
		},
		{
			name: "Should be valid when using hyphens",
			args: args{
				search: "my-kafka-topic",
			},
			wantErr: false,
		},
		{
			name: "Should be valid when starts with number",
			args: args{
				search: "1my-kafka-instance",
			},
			wantErr: false,
		},
		{
			name: "Should be valid when only two dots",
			args: args{
				search: "..",
			},
			wantErr: false,
		},
		{
			name: "Should be valid when only one dot",
			args: args{
				search: ".",
			},
			wantErr: false,
		},
		{
			name: "Should be valid when only three dots",
			args: args{
				search: "...",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// nolint
			if err := validator.ValidateSearchInput(tt.args.search); (err != nil) != tt.wantErr {
				t.Errorf("ValidateSearchInput() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidatePartitionsN(t *testing.T) {
	type args struct {
		partitions string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should throw error when non-numeric value is provided",
			args: args{
				partitions: "kafka",
			},
			wantErr: true,
		},
		{
			name: "Should be valid when numeric value is provided",
			args: args{
				partitions: "1",
			},
			wantErr: false,
		},
		{
			name: "Should throw an error when 0 is passed",
			args: args{
				partitions: "0",
			},
			wantErr: true,
		},
		{
			name: "Should throw an error when negative number is passed",
			args: args{
				partitions: "-1",
			},
			wantErr: true,
		},
		{
			name: "Should be invalid when containing decimal point",
			args: args{
				partitions: "3.0",
			},
			wantErr: true,
		},
		{
			name: "Should be valid when equal to max allowed value(1000)",
			args: args{
				partitions: "1000",
			},
			wantErr: false,
		},
		{
			name: "Should be invalid when exceeds 1000",
			args: args{
				partitions: "1002",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// nolint
			if err := validator.ValidatePartitionsN(tt.args.partitions); (err != nil) != tt.wantErr {
				t.Errorf("ValidatePartitionsN() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateMessageRetentionPeriod(t *testing.T) {
	type args struct {
		retentionMs string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should throw error when non-numeric value is provided",
			args: args{
				retentionMs: "kafka",
			},
			wantErr: true,
		},
		{
			name: "Should be valid when numeric value is provided",
			args: args{
				retentionMs: "1",
			},
			wantErr: false,
		},
		{
			name: "Should be valid when 0 is passed",
			args: args{
				retentionMs: "0",
			},
			wantErr: false,
		},
		{
			name: "Should be valid when -1 is passed",
			args: args{
				retentionMs: "-1",
			},
			wantErr: false,
		},
		{
			name: "Should be invalid when value less than -1 is passed",
			args: args{
				retentionMs: "-2",
			},
			wantErr: true,
		},
		{
			name: "Should be invalid when containing decimal point",
			args: args{
				retentionMs: "3.0",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// nolint
			if err := validator.ValidateMessageRetentionPeriod(tt.args.retentionMs); (err != nil) != tt.wantErr {
				t.Errorf("ValidateMessageRetentionPeriod() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateMessageRetentionSize(t *testing.T) {
	type args struct {
		retentionBytes string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should throw error when non-numeric value is provided",
			args: args{
				retentionBytes: "kafka",
			},
			wantErr: true,
		},
		{
			name: "Should be valid when numeric value is provided",
			args: args{
				retentionBytes: "1",
			},
			wantErr: false,
		},
		{
			name: "Should be valid when 0 is passed",
			args: args{
				retentionBytes: "0",
			},
			wantErr: false,
		},
		{
			name: "Should be valid when -1 is passed",
			args: args{
				retentionBytes: "-1",
			},
			wantErr: false,
		},
		{
			name: "Should be invalid when value less than -1 is passed",
			args: args{
				retentionBytes: "-2",
			},
			wantErr: true,
		},
		{
			name: "Should be invalid when containing decimal point",
			args: args{
				retentionBytes: "3.0",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// nolint
			if err := validator.ValidateMessageRetentionPeriod(tt.args.retentionBytes); (err != nil) != tt.wantErr {
				t.Errorf("ValidateMessageRetentionSize() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
