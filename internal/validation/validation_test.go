package validation

import "testing"

func TestValidator_AddError(t *testing.T) {
	v := New()
	v.AddError("title", "must be provided")

	if v.Errors["title"] != "must be provided" {
		t.Error("expected error for title field not found")
	}
}

func TestValidator_Check(t *testing.T) {
	v := New()
	v.Check(false, "title", "must be provided")

	if v.Errors["title"] != "must be provided" {
		t.Error("expected error for title field not found")
	}
}

func TestValidator_Valid(t *testing.T) {
	v := New()
	v.Check(false, "title", "must be provided")

	if v.Valid() {
		t.Error("expected invalid validator")
	}
}
