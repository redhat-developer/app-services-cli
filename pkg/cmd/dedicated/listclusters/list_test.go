package listclusters

import (
	"testing"

	"github.com/redhat-developer/app-services-cli/pkg/shared/kafkautil"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/kafkamgmt/apiv1/client"
)

func TestCreateSearchString(t *testing.T) {
	// Define test cases
	tests := []struct {
		name     string
		opts     *options
		expected string
	}{
		{
			name: "empty list",
			opts: &options{kfmClusterList: &kafkamgmtclient.EnterpriseClusterList{
				Items: []kafkamgmtclient.EnterpriseClusterListItem{},
			},
			},
			expected: "",
		},
		{
			name: "single clusters",
			opts: &options{kfmClusterList: &kafkamgmtclient.EnterpriseClusterList{
				Items: []kafkamgmtclient.EnterpriseClusterListItem{
					{
						Id: "abc",
					},
				},
			}},

			expected: "id = 'abc'",
		},
		{
			name: "two clusters",
			opts: &options{kfmClusterList: &kafkamgmtclient.EnterpriseClusterList{
				Items: []kafkamgmtclient.EnterpriseClusterListItem{
					{Id: "abc"},
					{Id: "def"},
				}}},
			expected: "id = 'abc' or id = 'def'",
		},
		{
			name: "multiple clusters",
			opts: &options{kfmClusterList: &kafkamgmtclient.EnterpriseClusterList{
				Items: []kafkamgmtclient.EnterpriseClusterListItem{
					{Id: "abc"},
					{Id: "def"},
					{Id: "ghi"},
				}}},
			expected: "id = 'abc' or id = 'def' or id = 'ghi'",
		},
	}

	// Run test cases
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Call the function
			actual := kafkautil.CreateClusterSearchStringFromKafkaList(tc.opts.kfmClusterList)

			// Compare the result with the expected value
			if actual != tc.expected {
				t.Errorf("unexpected result - \ngot: %v, \nwant: %v", actual, tc.expected)
			}
		})
	}
}
