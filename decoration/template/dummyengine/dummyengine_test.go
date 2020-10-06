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

func Test_dummyEngine_ProcessTemplate_WithInvalidTemplate_ShouldReturnError(t *testing.T) {
	eng := NewEngine()
	_, err := eng.ProcessTemplate("", struct{}{})

	wantErr := "invalid template"
	if err == nil || err.Error() != wantErr {
		t.Errorf("Should return error message \"%s\" but return \"%s\"", wantErr, err)
	}
}

func Test_dummyEngine_ProcessTemplate_WithInvalidTemplate_ShouldNotReturnValue(t *testing.T) {
	eng := NewEngine()
	got, _ := eng.ProcessTemplate("", struct{}{})

	if got != "" {
		t.Errorf("Should return nil but empty \"%s\"", got)
	}
}
