package main

import (
	"go/token"
	"go/types"
	"testing"
)

func TestParseEmbeddedFields(t *testing.T) {
	inner := types.NewStruct([]*types.Var{
		types.NewField(token.NoPos, nil, "CreatedAt", types.Typ[types.String], false),
		types.NewField(token.NoPos, nil, "Ignored", types.Typ[types.String], false),
	}, []string{
		`gorm:"column:created_at"`,
		`gorm:"-"`,
	})
	named := types.NewNamed(types.NewTypeName(token.NoPos, nil, "ModelTime", nil), inner, nil)

	fields := parseEmbeddedFields(named)
	if len(fields) != 1 {
		t.Fatalf("expected one persisted field, got %d", len(fields))
	}
	if fields[0].name != "CreatedAt" || fields[0].column != "created_at" {
		t.Fatalf("unexpected field: %#v", fields[0])
	}
}
