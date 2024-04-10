package expr

import "testing"

func TestNumberExpr(t *testing.T) {

	tests := []struct {
		e AssignExpr
		s string
		r string
	}{
		{AssignExpr{"key", "456"}, "key = 123", "key = 456"},
		{AssignExpr{"key", "456"}, "key=123", "key=456"},
		{AssignExpr{"key", "456"}, "key :=123", "key :=456"},
		{AssignExpr{"key", "456"}, "var key =123", "var key =456"},
		{AssignExpr{"key", "456"}, " key := 123 ", " key := 456 "},
	}

	for _, test := range tests {
		r, ok := test.e.Assign(test.s)
		if !ok {
			t.Errorf("expected true, got false")
		}
		if r != test.r {
			t.Errorf("expected %s, got %s", test.r, r)
		}
	}
}

func TestStringExpr(t *testing.T) {

	tests := []struct {
		e AssignExpr
		s string
		r string
	}{
		{AssignExpr{"key", "\"abc\""}, "key = \"aaa\"", "key = \"abc\""},
		{AssignExpr{"key", "\"abc\""}, "key=\"aaa\"", "key=\"abc\""},
		{AssignExpr{"key", "\"abc\""}, "key =\"aaa\"", "key =\"abc\""},
		{AssignExpr{"key", "\"abc\""}, "var key =\"aaa\"", "var key =\"abc\""},
		{AssignExpr{"key", "\"abc\""}, " key := \"aaa\" ", " key := \"abc\" "},
		{AssignExpr{"key", "\"abc\""}, "key=\"aaa bbb\"", "key=\"abc\""},
		{AssignExpr{"key", "\"abc\""}, "key=\"aaa \\\"bbb\"", "key=\"abc\""},
	}

	for _, test := range tests {
		r, ok := test.e.Assign(test.s)
		if !ok {
			t.Errorf("expected true, got false")
		}
		if r != test.r {
			t.Errorf("expected %s, got %s", test.r, r)
		}
	}
}

func TestNotExpr(t *testing.T) {

	tests := []struct {
		e AssignExpr
		s string
		r string
	}{
		{AssignExpr{"key", "456"}, "akey = 123", "key = 456"},
		{AssignExpr{"key", "456"}, "keyb=123", "key=456"},
		{AssignExpr{"key", "456"}, "k ey =123", "key =456"},
	}

	for _, test := range tests {
		_, ok := test.e.Assign(test.s)
		if ok {
			t.Errorf("expected true, got false")
		}
	}
}
