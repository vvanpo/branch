package titian

import (
	"bytes"
	"testing"
	"time"
)

// Tests the LabelFieldValue type.
func TestLabelFieldValue(t *testing.T) {
}

// Tests the Marshal/Unmarshal methods of DateFieldValue.
func TestDateFieldValueMarshalling(t *testing.T) {
	field := stubField()
	value := &DateFieldValue{newID(), field, time.Now()}
	buf, err := value.MarshalBinary()

	if err != nil {
		t.Fail()
	}

	err = value.UnmarshalBinary(buf)

	if err != nil {
		t.Fail()
	}

	buf2, err := value.MarshalBinary()

	if err != nil {
		t.Fail()
	}

	if !bytes.Equal(buf, buf2) {
		t.Fail()
	}
}

func stubField() *Field {
	app := New(Config{})
	category, _ := app.fields.NewCategory("top", nil)
	field, _ := category.NewField("label", &LabelField{})
	return field
}
