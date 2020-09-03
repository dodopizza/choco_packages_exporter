package choco

import (
	"reflect"
	"testing"
)

func Test_extractPackageInfoFromPackagesMultilineString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			"3 elements list",
			args{"aspnetcore-runtimepackagestore|3.1.6\n7zip|8.1.0\nKB2919442|1.0.20160915"},
			[][]string{
				{"aspnetcore-runtimepackagestore|3.1.6", "aspnetcore-runtimepackagestore", "3.1.6"},
				{"7zip|8.1.0", "7zip", "8.1.0"},
				{"KB2919442|1.0.20160915", "KB2919442", "1.0.20160915"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractPackageInfoFromPackagesMultilineString(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractPackageInfoFromPackagesMultilineString() = %v, want %v", got, tt.want)
			}
		})
	}
}