package list

import (
	"bytes"
	"errors"
	"reflect"
	"testing"
)

func TestNewListCommand(t *testing.T) {
	tests := []struct {
		name    string
		args    [][]string
		wantErr error
	}{
		{
			name: "return error for invalid output format",
			args: [][]string{
				{"--output", "json2"},
			},
			wantErr: errors.New("Invalid output format 'json2'"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := NewListCommand()
			cmd.Help()
			b := bytes.NewBufferString("")
			cmd.SetOut(b)
			for _, arg := range tt.args {
				cmd.SetArgs(arg)
			}

			if got := cmd.Execute(); !reflect.DeepEqual(got, tt.wantErr) {
				t.Fatalf("NewListCommand() error = %v, want %v", got, tt.wantErr)
				return
			}
		})
	}
}
