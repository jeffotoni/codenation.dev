package rpg

import "testing"

func TestSelect(t *testing.T) {
	type args struct {
		codemp int64
	}
	tests := []struct {
		name    string
		args    args
		wantOk  bool
		wantMsg string
	}{
		// TODO: Add test cases.
		{"test_select_1", args{1}, true, "ok"},
		{"test_select_2", args{2}, true, "ok"},
		{"test_select_3", args{3}, true, "ok"},
		{"test_select_4", args{4}, true, "ok"},
		{"test_select_5", args{5}, true, "ok"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOk, gotMsg := Select(tt.args.codemp)
			if gotOk != tt.wantOk {
				t.Errorf("Select() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
			if gotMsg != tt.wantMsg {
				t.Errorf("Select() gotMsg = %v, want %v", gotMsg, tt.wantMsg)
			}
		})
	}
}
