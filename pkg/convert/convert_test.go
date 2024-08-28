package convert

import "testing"

func TestKelToCel(t *testing.T) {
	type args struct {
		kelvin float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test 1",
			args: args{
				kelvin: 273.15,
			},
			want: 0,
		},
		{
			name: "Test 2",
			args: args{
				kelvin: 323.1555555555,
			},
			want: 50,
		},
		{
			name: "Test 3",
			args: args{
				kelvin: 274.15,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KelToCel(tt.args.kelvin); got != tt.want {
				t.Errorf("KelToCel() = %v, want %v", got, tt.want)
			}
		})
	}
}
