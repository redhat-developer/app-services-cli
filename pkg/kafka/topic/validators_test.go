package topic

import "testing"

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
			wantErr: false,
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
			if err := ValidateName(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("ValidateName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
