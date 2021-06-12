package ritmo

import (
	"reflect"
	"testing"
)

func TestStart(t *testing.T) {
	type args struct {
		requestsPerSecond int
		maxBurst          int
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Start(tt.args.requestsPerSecond, tt.args.maxBurst); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Start() = %v, want %v", got, tt.want)
			}
		})
	}
}
