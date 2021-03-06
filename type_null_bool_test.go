package recurly

import (
	"bytes"
	"encoding/xml"
	"reflect"
	"testing"
)

func TestNullBool(t *testing.T) {
	given0 := NewBool(true)
	expected0 := NullBool{Bool: true, Valid: true}

	given1 := NewBool(false)
	expected1 := NullBool{Bool: false, Valid: true}

	given2 := NullBool{Bool: false, Valid: false}

	if !reflect.DeepEqual(expected0, given0) {
		t.Fatalf("unexpected value: %v", given0)
	} else if !given0.Is(true) {
		t.Fatalf("unexpected value")
	} else if !reflect.DeepEqual(expected1, given1) {
		t.Fatalf("unexpected value")
	} else if !given1.Is(false) {
		t.Fatalf("unexpected value")
	} else if given2.Is(false) {
		t.Fatalf("unexpected value")
	}

	type s struct {
		XMLName xml.Name `xml:"s"`
		Name    string   `xml:"name"`
		Exempt  NullBool `xml:"exempt,omitempty"`
	}

	tests := []struct {
		v        s
		expected string
	}{
		{v: s{XMLName: xml.Name{Local: "s"}, Name: "Bob", Exempt: given0}, expected: "<s><name>Bob</name><exempt>true</exempt></s>"},
		{v: s{XMLName: xml.Name{Local: "s"}, Name: "Bob", Exempt: given1}, expected: "<s><name>Bob</name><exempt>false</exempt></s>"},
		{v: s{XMLName: xml.Name{Local: "s"}, Name: "Bob", Exempt: given2}, expected: "<s><name>Bob</name></s>"},
	}

	for i, tt := range tests {
		var given bytes.Buffer
		if err := xml.NewEncoder(&given).Encode(tt.v); err != nil {
			t.Fatalf("unexpected value")
		} else if tt.expected != given.String() {
			t.Fatalf("(%d): unexpected value: %v", i, given.String())
		}

		buf := bytes.NewBufferString(tt.expected)
		var dest s
		if err := xml.NewDecoder(buf).Decode(&dest); err != nil {
			t.Fatalf("(%d): %v", i, err)
		}

		if !reflect.DeepEqual(tt.v, dest) {
			t.Fatalf("(%d): %v", i, dest)
		}
	}
}
