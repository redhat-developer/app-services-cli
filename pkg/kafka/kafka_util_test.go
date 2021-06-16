package kafka

import (
	"testing"

	"github.com/redhat-developer/app-services-cli/pkg/localize/goi18n"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
)

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
			name: "Should throw error when exceeds 32 characters",
			args: args{
				name: "verylongkafkanamef8d9dkf9dkc11dfs",
			},
			wantErr: true,
		},
		{
			name: "Should be valid when name is exactly 32 characters",
			args: args{
				name: "verylongkafkanamef8d9dkf9dkc11dd",
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
			name: "Should be invalid when using hyphens",
			args: args{
				name: "my-kafka-instance",
			},
			wantErr: false,
		},
		{
			name: "Should be invalid when starts with number",
			args: args{
				name: "1my-kafka-instance",
			},
			wantErr: true,
		},
		{
			name: "Should be invalid when includes uppercase letter",
			args: args{
				name: "my-Kafka-instance",
			},
			wantErr: true,
		},
		{
			name: "Should be invalid when includes a space",
			args: args{
				name: "my-kafka instance",
			},
			wantErr: true,
		},
		{
			name: "Should be invalid when includes a '.'",
			args: args{
				name: "my-kafka.instance",
			},
			wantErr: true,
		},
		{
			name: "Should be invalid when starts with a '-'",
			args: args{
				name: "-my-kafka-instance",
			},
			wantErr: true,
		},
		{
			name: "Should be invalid when ends with a '-'",
			args: args{
				name: "my-kafka-instance-",
			},
			wantErr: true,
		},
	}

	localizer, _ := goi18n.New(nil)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// nolint
			if err := ValidateName(localizer)(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("ValidateName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTransformRequest(t *testing.T) {
	hostWithSSLPort := "my-kafka-url:443"
	hostWithNoPort := "my-kafka-url"

	type args struct {
		kafkaInstance *kafkamgmtclient.KafkaRequest
	}
	tests := []struct {
		name                    string
		args                    args
		wantBootstrapServerHost string
	}{
		{
			name: "bootstrapServerHost should be transformed to empty string when nil",
			args: args{
				kafkaInstance: &kafkamgmtclient.KafkaRequest{
					BootstrapServerHost: nil,
				},
			},
			wantBootstrapServerHost: "",
		},
		{
			name: "bootstrapServerHost should be the same when SSL port already exists",
			args: args{
				kafkaInstance: &kafkamgmtclient.KafkaRequest{
					BootstrapServerHost: &hostWithSSLPort,
				},
			},
			wantBootstrapServerHost: hostWithSSLPort,
		},
		{
			name: "bootstrapServerHost should get SSL port when none exists",
			args: args{
				kafkaInstance: &kafkamgmtclient.KafkaRequest{
					BootstrapServerHost: &hostWithNoPort,
				},
			},
			wantBootstrapServerHost: hostWithSSLPort,
		},
	}
	for _, tt := range tests {
		// nolint
		t.Run(tt.name, func(t *testing.T) {
			transformedInstance := TransformKafkaRequest(tt.args.kafkaInstance)

			if transformedInstance == nil {
				t.Errorf("Expected KafkaRequest type, but got nil")
			}

			transformedHost := transformedInstance.GetBootstrapServerHost()
			if tt.wantBootstrapServerHost != transformedHost {
				t.Errorf("Expected bootstrapServerHost: %v, got %v", tt.wantBootstrapServerHost, transformedHost)
			}
		})
	}
}

func TestTransformKafkaRequestListItems(t *testing.T) {
	hostWithSSLPort := "my-kafka-url:443"
	hostWithNoPort := "my-kafka-url"
	emptyHost := ""

	type args struct {
		items []kafkamgmtclient.KafkaRequest
	}
	tests := []struct {
		name string
		args args
		want []kafkamgmtclient.KafkaRequest
	}{
		{
			name: "Should return empty list when no kafkas",
			args: args{
				items: []kafkamgmtclient.KafkaRequest{},
			},
			want: []kafkamgmtclient.KafkaRequest{},
		},
		{
			name: "Should update all bootstrapServerHost ports",
			args: args{
				items: []kafkamgmtclient.KafkaRequest{
					{
						BootstrapServerHost: &emptyHost,
					},
					{
						BootstrapServerHost: &hostWithNoPort,
					},
					{
						BootstrapServerHost: &hostWithSSLPort,
					},
				},
			},
			want: []kafkamgmtclient.KafkaRequest{
				{
					BootstrapServerHost: &emptyHost,
				},
				{
					BootstrapServerHost: &hostWithSSLPort,
				},
				{
					BootstrapServerHost: &hostWithSSLPort,
				},
			},
		},
	}
	for _, tt := range tests {
		// nolint
		t.Run(tt.name, func(t *testing.T) {
			gotItems := TransformKafkaRequestListItems(tt.args.items)

			if len(gotItems) != len(tt.want) {
				t.Fatalf("Expected number of items to be %v, got %v", len(gotItems), len(tt.want))
				return
			}

			for j, wantKafka := range tt.want {
				gotKafka := gotItems[j]

				gotBootstrapHost := gotKafka.GetBootstrapServerHost()
				wantBootstrapHost := wantKafka.GetBootstrapServerHost()

				if gotBootstrapHost != wantBootstrapHost {
					t.Fatalf("Expected BootstrapServerHost = %v, got %v", wantBootstrapHost, gotBootstrapHost)
				}
			}
		})
	}
}
