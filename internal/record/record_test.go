
package record

import "testing"

func TestEncodeDecode(t *testing.T) {
	original := Record{
		Key:     "name",
		Value:   "Kader",
		Deleted: false,
	}

	data, err := Encode(original)
	if err != nil {
		t.Fatal(err)
	}

	decoded, err := Decode(data)
	if err != nil {
		t.Fatal(err)
	}

	if decoded != original {
		t.Fatalf("expected %+v, got %+v", original, decoded)
	}
}

func TestDeletedRecord(t *testing.T) {
	original := Record{
		Key:     "name",
		Value:   "",
		Deleted: true,
	}

	data, err := Encode(original)
	if err != nil {
		t.Fatal(err)
	}

	decoded, err := Decode(data)
	if err != nil {
		t.Fatal(err)
	}

	if decoded != original {
		t.Fatalf("expected %+v, got %+v", original, decoded)
	}
}
