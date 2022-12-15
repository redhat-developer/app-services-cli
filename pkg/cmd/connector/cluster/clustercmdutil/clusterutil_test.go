package clustercmdutil

import "testing"

func TestConvertToMap(t *testing.T) {

	type args struct {
		keyValPairs []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should be valid when alphanumeric key and values are provided",
			args: args{
				keyValPairs: []string{"environment=production", "app=nginx"},
			},
			wantErr: false,
		},
		{
			name: "Should be valid when a \"/\" is provided in the key",
			args: args{
				keyValPairs: []string{"cos.io/environment=production", "app=nginx"},
			},
			wantErr: false,
		},
		{
			name: "Should throw error when key contains special characters",
			args: args{
				keyValPairs: []string{"cos.io/environment??=production", "app=nginx"},
			},
			wantErr: true,
		},
		{
			name: "Should throw error when key starts with \".\"",
			args: args{
				keyValPairs: []string{".environment=production", "app=nginx"},
			},
			wantErr: true,
		},
		{
			name: "Should be valid when key contains \".\"",
			args: args{
				keyValPairs: []string{"envi.ronment=production", "app=nginx"},
			},
			wantErr: false,
		},
		{
			name: "Should be valid when value contains \".\"",
			args: args{
				keyValPairs: []string{"environment=production.1", "app=nginx.2"},
			},
			wantErr: false,
		},
		{
			name: "Should throw an error when value ends with \".\"",
			args: args{
				keyValPairs: []string{"environment=production.", "app=nginx."},
			},
			wantErr: true,
		},
		{
			name: "Should throw an error when key contains a reserved domain",
			args: args{
				keyValPairs: []string{"openshift.io/env=minishift", "app=nginx"},
			},
			wantErr: true,
		},
		{
			name: "Should be valid when value is empty",
			args: args{
				keyValPairs: []string{"environment=", "app"},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// nolint
			if _, err := BuildAnnotationsMap(tt.args.keyValPairs); (err != nil) != tt.wantErr {
				t.Errorf("ConvertAnnotationsToMap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
