package listclusters

import (
	"testing"

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
				Items: []kafkamgmtclient.EnterpriseCluster{},
			},
			},
			expected: "",
		},
		{
			name: "single clusters",
			opts: &options{kfmClusterList: &kafkamgmtclient.EnterpriseClusterList{
				Items: []kafkamgmtclient.EnterpriseCluster{
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
				Items: []kafkamgmtclient.EnterpriseCluster{
					{Id: "abc"},
					{Id: "def"},
				}}},
			expected: "id = 'abc' or id = 'def'",
		},
		{
			name: "multiple clusters",
			opts: &options{kfmClusterList: &kafkamgmtclient.EnterpriseClusterList{
				Items: []kafkamgmtclient.EnterpriseCluster{
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
			actual := createSearchString(tc.opts)

			// Compare the result with the expected value
			if actual != tc.expected {
				t.Errorf("unexpected result - \ngot: %v, \nwant: %v", actual, tc.expected)
			}
		})
	}
}
