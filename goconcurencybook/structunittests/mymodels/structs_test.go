package mymodels

import "testing"

func Test_my_operations(t *testing.T) {
	type fields struct {
		Branch bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"name", fields{false},false},
		{"name", fields{true},true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &my{
				Branch: tt.fields.Branch,
			}
			if err := s.operations(); (err != nil) != tt.wantErr {
				t.Errorf("my.operations() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
