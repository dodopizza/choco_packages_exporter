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
			"2 elements list",
			args{"Chocolatey v0.10.11\n7zip.install 19.0"},
			[][]string{
				{"Chocolatey v0.10.11", "Chocolatey", "0.10.11"},
				{"7zip.install 19.0", "7zip.install", "19.0"},
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