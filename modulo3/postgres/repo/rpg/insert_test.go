// Back-End in Go server
// @jeffotoni
// 2019-01-04

package rpg

import "testing"

func TestInsert(t *testing.T) {
	type args struct {
		cnpj        string
		razaosocial string
		ativo       bool
	}
	tests := []struct {
		name    string
		args    args
		wantOk  bool
		wantMsg string
	}{
		// TODO: Add test cases.
		{"tes_insert_1", args{"08776968001271", "Empresa 10", true}, true, "ok"},
		{"tes_insert_2", args{"08776968001372", "Empresa 11", true}, true, "ok"},
		{"tes_insert_3", args{"08776968001372", "Empresa 12", true}, true, "ok"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOk, gotMsg := Insert(tt.args.cnpj, tt.args.razaosocial, tt.args.ativo)
			if gotOk != tt.wantOk {
				t.Errorf("Insert() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
			if gotMsg != tt.wantMsg {
				t.Errorf("Insert() gotMsg = %v, want %v", gotMsg, tt.wantMsg)
			}
		})
	}
}
