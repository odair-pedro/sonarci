package dummy

import (
	"sonarci/decoration/template/engine/dummy/resources"
	"testing"
)

func Test_dummyEngine_GetQualityReportTemplate_CheckReportTitleWithTag(t *testing.T) {
	eng := NewEngine()

	want := resources.QualityReportTemplateTitleWithTag + resources.QualityReportTemplate
	got := eng.GetQualityReportTemplate(true)

	if got != want {
		t.Fail()
	}
}

func Test_dummyEngine_GetQualityReportTemplate_CheckReportTitleWithoutTag(t *testing.T) {
	eng := NewEngine()

	want := resources.QualityReportTemplateTitle + resources.QualityReportTemplate

	got := eng.GetQualityReportTemplate()
	if got != want {
		t.Fail()
	}

	got = eng.GetQualityReportTemplate(false)
	if got != want {
		t.Fail()
	}

	got = eng.GetQualityReportTemplate("anything")
	if got != want {
		t.Fail()
	}
}

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

func Test_dummyEngine_ProcessTemplate_WithEscape(t *testing.T) {
	eng := NewEngine()
	got, _ := eng.ProcessTemplate("Value1 = \"${field1}\"\nValue2 = \"${field2}\"", struct {
		field1 string `dummy:"field1"`
		field2 string `dummy:"field2" escape:"true"`
	}{
		field1: "value1",
		field2: "value2 test with escape",
	})

	want := "Value1 = \"value1\"\nValue2 = \"value2%20test%20with%20escape\""
	if got != want {
		t.Errorf("Result invalid, want \"%s\" but got \"%s\"", want, got)
	}
}

func Test_dummyEngine_ProcessTemplate_WithoutEscape(t *testing.T) {
	eng := NewEngine()
	got, _ := eng.ProcessTemplate("Value1 = \"${field1}\"\nValue2 = \"${field2}\"", struct {
		field1 string `dummy:"field1"`
		field2 string `dummy:"field2"`
	}{
		field1: "value1",
		field2: "value2 test without escape",
	})

	want := "Value1 = \"value1\"\nValue2 = \"value2 test without escape\""
	if got != want {
		t.Errorf("Result invalid, want \"%s\" but got \"%s\"", want, got)
	}
}

func Test_dummyEngine_ProcessTemplate_WithEmptyTemplate_ShouldReturnError(t *testing.T) {
	eng := NewEngine()
	_, err := eng.ProcessTemplate("", struct{}{})

	wantErr := "invalid template"
	if err == nil || err.Error() != wantErr {
		t.Errorf("Should return error message \"%s\" but return \"%s\"", wantErr, err)
	}
}

func Test_dummyEngine_ProcessTemplate_WithEmptyTemplate_ShouldNotReturnValue(t *testing.T) {
	eng := NewEngine()
	got, _ := eng.ProcessTemplate("", struct{}{})

	if got != "" {
		t.Errorf("Should return nil but empty \"%s\"", got)
	}
}

func Test_dummyEngine_ProcessTemplate_WithNilDataSource_ShouldReturnError(t *testing.T) {
	eng := NewEngine()
	_, err := eng.ProcessTemplate("test", nil)

	wantErr := "invalid data source"
	if err == nil || err.Error() != wantErr {
		t.Errorf("Should return error message \"%s\" but return \"%s\"", wantErr, err)
	}
}

func Test_dummyEngine_ProcessTemplate_WithNilDataSource_ShouldNotReturnValue(t *testing.T) {
	eng := NewEngine()
	got, _ := eng.ProcessTemplate("test", nil)

	if got != "" {
		t.Errorf("Should return nil but empty \"%s\"", got)
	}
}

func Test_dummyEngine_ProcessTemplate_WithInvalidDataSource_ShouldReturnError(t *testing.T) {
	eng := NewEngine()
	_, err := eng.ProcessTemplate("test", "test")

	wantErr := "invalid data source, it is not a struct"
	if err == nil || err.Error() != wantErr {
		t.Errorf("Should return error message \"%s\" but return \"%s\"", wantErr, err)
	}
}

func Test_dummyEngine_ProcessTemplate_WithInvalidDataSource_ShouldNotReturnValue(t *testing.T) {
	eng := NewEngine()
	got, _ := eng.ProcessTemplate("test", "test")

	if got != "" {
		t.Errorf("Should return nil but empty \"%s\"", got)
	}
}

func Test_processDataSourceField_WithEscape_ShouldReturnExpectedValue(t *testing.T) {
	got := processDataSourceField("test", "Hi, i'm testing here", "Testing message: \"${test}\"", true)
	want := "Testing message: \"Hi%2C%20i%27m%20testing%20here\""

	if got != want {
		t.Errorf("Should process right value. Got \"%s\" but want \"%s\"", got, want)
	}
}

func Test_processDataSourceField_WithoutEscape_ShouldReturnExpectedValue(t *testing.T) {
	got := processDataSourceField("test", "Hi, i'm testing here", "Testing message: \"${test}\"", false)
	want := "Testing message: \"Hi, i'm testing here\""

	if got != want {
		t.Errorf("Should process right value. Got \"%s\" but want \"%s\"", got, want)
	}
}

func Test_escapeValue_ShouldEscabeCharacters(t *testing.T) {
	got := escapeValue("98.0%")
	want := "98.0%25"

	if got != want {
		t.Errorf("Should escape special characters. Got \"%s\" but want \"%s\"", got, want)
	}
}
