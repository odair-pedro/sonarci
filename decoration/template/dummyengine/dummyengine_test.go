package dummyengine

import "testing"

func Test_dummyEngine_ProcessTemplate(t *testing.T) {
	eng := NewEngine()
	got, _ := eng.ProcessTemplate("Value1 = \"${field1}\"\nValue2 = \"${field2}\"", struct {
		field1 string `dummy:"field1"`
		field2 string `dummy:"field2"`
	}{
		field1: "value1",
		field2: "value2",
	})

	want := "Value1 = \"value1\"\nValue2 = \"value2\""
	if got != want {
		t.Errorf("Result invalid, want \"%s\" but got \"%s\"", want, got)
	}
}
