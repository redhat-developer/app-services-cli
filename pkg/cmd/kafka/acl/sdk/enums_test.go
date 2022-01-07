package sdk

import (
	"reflect"
	"testing"
)

func TestGetMappedValidResourceOperations(t *testing.T) {
	resourceOperationsMap := map[string][]string{
		"cluster":          {"describe", "alter"},
		"group":            {"all", "delete", "describe", "read"},
		"transactional_id": {"all", "describe", "write"},
		"topic": {
			"all",
			"alter",
			"alter_configs",
			"create",
			"delete",
			"describe",
			"describe_configs",
			"read",
			"write",
		},
	}
	type args struct {
		resourceType string
		operations   []string
	}
	tests := []struct {
		name        string
		args        args
		wantIsValid bool
		wantOps     []string
	}{
		{
			name: "are invalid cluster operations",
			args: args{
				resourceType: "cluster",
				operations:   []string{"delete", "all"},
			},
			wantIsValid: false,
			wantOps:     []string{"describe", "alter"},
		},
		{
			name: "are valid cluster operations",
			args: args{
				resourceType: "cluster",
				operations:   []string{"alter"},
			},
			wantIsValid: true,
			wantOps:     nil,
		},
		{
			name: "are invalid group operations",
			args: args{
				resourceType: "group",
				operations:   []string{"alter"},
			},
			wantIsValid: false,
			wantOps:     []string{"all", "delete", "describe", "read"},
		},
		{
			name: "are valid group operations",
			args: args{
				resourceType: "group",
				operations:   []string{"all", "delete", "describe", "read"},
			},
			wantIsValid: true,
			wantOps:     nil,
		},
		{
			name: "are valid topic operations",
			args: args{
				resourceType: "topic",
				operations: []string{
					"all",
					"alter",
					"alter-configs",
					"create",
					"delete",
					"describe",
					"describe-configs",
					"read",
					"write",
				},
			},
			wantIsValid: true,
			wantOps:     nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, op := range tt.args.operations {
				t.Logf("Testing operation '%v' on resource type '%v'", op, tt.args.resourceType)
				isValid, validOperations := IsValidResourceOperation(tt.args.resourceType, op, resourceOperationsMap)
				if isValid != tt.wantIsValid {
					t.Errorf("IsValidResourceOperation() isValid = %v, want %v", isValid, tt.wantIsValid)
				}
				if !reflect.DeepEqual(validOperations, tt.wantOps) {
					t.Errorf("IsValidResourceOperation() = %v, want %v", validOperations, tt.wantOps)
				}
			}
		})
	}
}
