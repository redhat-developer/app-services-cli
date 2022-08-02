package topiccmdutil

import (
	"reflect"
	"testing"

	kafkainstanceclient "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1/client"
)

func TestCreateConfigEntries(t *testing.T) {
	keyOne := "key1"
	keyTwo := "key2"

	valOne := "1000"

	type args struct {
		entryMap map[string]*string
	}
	tests := []struct {
		name string
		args args
		want *[]kafkainstanceclient.ConfigEntry
	}{
		{
			name: "should convert config entry map to an array with the same values",
			args: args{
				entryMap: map[string]*string{
					keyOne: &valOne,
					keyTwo: nil,
				},
			},
			want: &[]kafkainstanceclient.ConfigEntry{
				{
					Key:   keyOne,
					Value: valOne,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// nolint:scopelint
			if got := CreateConfigEntries(tt.args.entryMap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateConfigEntries() = %v, want %v", got, tt.want)
			}
		})
	}
}
