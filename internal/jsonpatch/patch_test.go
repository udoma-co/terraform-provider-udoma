package jsonpatch

import (
	"testing"

	"github.com/spyzhov/ajson"
)

func TestValidations(t *testing.T) {

	// invalid operation
	patchSource := `[{ "op": "invalid", "path": "$.x" }]`
	_, err := DecodePatch([]byte(patchSource))
	if err == nil {
		t.Fatalf("Expected error")
	}

	// missing path
	patchSource = `[{ "op": "add" }]`
	_, err = DecodePatch([]byte(patchSource))
	if err == nil {
		t.Fatalf("Expected error")
	}

	// invalid path
	patchSource = `[{ "op": "add", "path": "$.x[" }]`
	_, err = DecodePatch([]byte(patchSource))
	if err == nil {
		t.Fatalf("Expected error")
	}

	// missing from
	patchSource = `[{ "op": "move", "path": "$.x" }]`
	_, err = DecodePatch([]byte(patchSource))
	if err == nil {
		t.Fatalf("Expected error")
	}

	// invalid from
	patchSource = `[{ "op": "move", "from": "$.x[", "path": "$.x" }]`
	_, err = DecodePatch([]byte(patchSource))
	if err == nil {
		t.Fatalf("Expected error")
	}
}

func TestPatchAddBasicAttribute(t *testing.T) {
	orig := `{"x":1}`
	patchSource := `[{ "op": "add", "path": "$.y", "value": 2 }]`

	patch, err := DecodePatch([]byte(patchSource))
	if err != nil {
		t.Fatalf("Could not decode patch, unexpected error: %v", err)
	}

	res, err := patch.Apply([]byte(orig))
	if err != nil {
		t.Fatalf("Could not apply patch, unexpected error: %v", err)
	}

	if !matches(res, `{"x":1,"y":2}`) {
		t.Fatalf("Unexpected result: %s", string(res))
	}
}

func TestPatchAddElementToArray(t *testing.T) {

	orig := `{"x":[{"y":"a"},{"y":"b"}]}`
	patchSource := `[{ "op": "add", "path": "$.x.[]", "value": {"y":"o"} }]`

	patch, err := DecodePatch([]byte(patchSource))
	if err != nil {
		t.Fatalf("Could not decode patch, unexpected error: %v", err)
	}

	res, err := patch.Apply([]byte(orig))
	if err != nil {
		t.Fatalf("Could not apply patch, unexpected error: %v", err)
	}
	if !matches(res, `{"x":[{"y":"a"},{"y":"b"},{"y":"o"}]}`) {
		t.Fatalf("Unexpected result: %s", string(res))
	}

	// add at position 1
	patchSource = `[{ "op": "add", "path": "$.x.1", "value": {"y":"o"} }]`

	patch, err = DecodePatch([]byte(patchSource))
	if err != nil {
		t.Fatalf("Could not decode patch, unexpected error: %v", err)
	}

	res, err = patch.Apply([]byte(orig))
	if err != nil {
		t.Fatalf("Could not apply patch, unexpected error: %v", err)
	}

	if !matches(res, `{"x":[{"y":"a"},{"y":"o"},{"y":"b"}]}`) {
		t.Fatalf("Unexpected result: %s", string(res))
	}

	// add at position -1
	patchSource = `[{ "op": "add", "path": "$.x.-1", "value": {"y":"o"} }]`

	patch, err = DecodePatch([]byte(patchSource))
	if err != nil {
		t.Fatalf("Could not decode patch, unexpected error: %v", err)
	}

	res, err = patch.Apply([]byte(orig))
	if err != nil {
		t.Fatalf("Could not apply patch, unexpected error: %v", err)
	}

	if !matches(res, `{"x":[{"y":"a"},{"y":"o"},{"y":"b"}]}`) {
		t.Fatalf("Unexpected result: %s", string(res))
	}

	// add to the end of the array
	patchSource = `[{ "op": "add", "path": "$.x.2", "value": {"y":"o"} }]`

	patch, err = DecodePatch([]byte(patchSource))
	if err != nil {
		t.Fatalf("Could not decode patch, unexpected error: %v", err)
	}

	res, err = patch.Apply([]byte(orig))
	if err != nil {
		t.Fatalf("Could not apply patch, unexpected error: %v", err)
	}

	if !matches(res, `{"x":[{"y":"a"},{"y":"b"},{"y":"o"}]}`) {
		t.Fatalf("Unexpected result: %s", string(res))
	}

}

func TestPatchAddElementToArrayViaFilter(t *testing.T) {

	orig := `{"x":[{"y":"a"},{"y":"b"}]}`
	patchSource := `[{ "op": "add", "path": "$.x[?(@.y=='a')].z", "value": 1 }]`

	patch, err := DecodePatch([]byte(patchSource))
	if err != nil {
		t.Fatalf("Could not decode patch, unexpected error: %v", err)
	}

	res, err := patch.Apply([]byte(orig))
	if err != nil {
		t.Fatalf("Could not apply patch, unexpected error: %v", err)
	}

	if !matches(res, `{"x":[{"y":"a","z":1},{"y":"b"}]}`) {
		t.Fatalf("Unexpected result: %s", string(res))
	}

	orig = `{"data":[{"id":"x","elements":[{"attr1":10,"attr2":true},{"attr1":20,"attr2":true}]}]}`
	patchSource = `[{ "op": "add", "path": "$.data[?(@.id=='x')].elements.0", "value": {"attr1":0,"attr2":false} }]`

	patch, err = DecodePatch([]byte(patchSource))
	if err != nil {
		t.Fatalf("Could not decode patch, unexpected error: %v", err)
	}

	res, err = patch.Apply([]byte(orig))
	if err != nil {
		t.Fatalf("Could not apply patch, unexpected error: %v", err)
	}

	if !matches(res, `{"data":[{"id":"x","elements":[{"attr1":0,"attr2":false},{"attr1":10,"attr2":true},{"attr1":20,"attr2":true}]}]}`) {
		t.Fatalf("Unexpected result: %s", string(res))
	}
}

func TestPatchRemoveBasicAttribute(t *testing.T) {
	orig := `{"x":1,"y":2}`
	patchSource := `[{ "op": "remove", "path": "$.y" }]`

	patch, err := DecodePatch([]byte(patchSource))
	if err != nil {
		t.Fatalf("Could not decode patch, unexpected error: %v", err)
	}

	res, err := patch.Apply([]byte(orig))
	if err != nil {
		t.Fatalf("Could not apply patch, unexpected error: %v", err)
	}

	if !matches(res, `{"x":1}`) {
		t.Fatalf("Unexpected result: %s", string(res))
	}
}

func TestPatchRemoveElementFromArray(t *testing.T) {

	orig := `{"x":[{"y":"a"},{"y":"b"},{"y":"c"}]}`

	// remove first item in array
	patchSource := `[{ "op": "remove", "path": "$.x.0" }]`

	patch, err := DecodePatch([]byte(patchSource))
	if err != nil {
		t.Fatalf("Could not decode patch, unexpected error: %v", err)
	}

	res, err := patch.Apply([]byte(orig))
	if err != nil {
		t.Fatalf("Could not apply patch, unexpected error: %v", err)
	}

	if !matches(res, `{"x":[{"y":"b"},{"y":"c"}]}`) {
		t.Fatalf("Unexpected result: %s", string(res))
	}

	// remove last item in array
	patchSource = `[{ "op": "remove", "path": "$.x.2"}]`

	patch, err = DecodePatch([]byte(patchSource))
	if err != nil {
		t.Fatalf("Could not decode patch, unexpected error: %v", err)
	}

	res, err = patch.Apply([]byte(orig))
	if err != nil {
		t.Fatalf("Could not apply patch, unexpected error: %v", err)
	}

	if !matches(res, `{"x":[{"y":"a"},{"y":"b"}]}`) {
		t.Fatalf("Unexpected result: %s", string(res))
	}
}

func TestPatchRemoveElementFromArrayViaFilter(t *testing.T) {

	orig := `{"x":[{"y":"a"},{"y":"b"}] }`
	patchSource := `[{ "op": "remove", "path": "$.x[?(@.y=='a')]" }]`

	patch, err := DecodePatch([]byte(patchSource))
	if err != nil {
		t.Fatalf("Could not decode patch, unexpected error: %v", err)
	}

	res, err := patch.Apply([]byte(orig))
	if err != nil {
		t.Fatalf("Could not apply patch, unexpected error: %v", err)
	}

	if !matches(res, `{"x":[{"y":"b"}]}`) {
		t.Fatalf("Unexpected result: %s", string(res))
	}
}

func TestReplaceAttribute(t *testing.T) {
	orig := `{"x":{"a":1,"b":2}}`
	patchSource := `[{ "op": "replace", "path": "$.x.b", "value": 3 }]`

	patch, err := DecodePatch([]byte(patchSource))
	if err != nil {
		t.Fatalf("Could not decode patch, unexpected error: %v", err)
	}

	res, err := patch.Apply([]byte(orig))
	if err != nil {
		t.Fatalf("Could not apply patch, unexpected error: %v", err)
	}

	if !matches(res, `{"x":{"a":1,"b":3}}`) {
		t.Fatalf("Unexpected result: %s", string(res))
	}

	patchSource = `[{ "op": "replace", "path": "$.x", "value": {"a":"i","b":"e"} }]`

	patch, err = DecodePatch([]byte(patchSource))
	if err != nil {
		t.Fatalf("Could not decode patch, unexpected error: %v", err)
	}

	res, err = patch.Apply([]byte(orig))
	if err != nil {
		t.Fatalf("Could not apply patch, unexpected error: %v", err)
	}

	if !matches(res, `{"x":{"a":"i","b":"e"}}`) {
		t.Fatalf("Unexpected result: %s", string(res))
	}
}

func TestReplaceEllementInArray(t *testing.T) {

	orig := `{"x":[{"y":"a"},{"y":"b"},{"y":"c"}]}`

	patchSource := `[{ "op": "replace", "path": "$.x.0", "value": {"v":"d"} }]`

	patch, err := DecodePatch([]byte(patchSource))
	if err != nil {
		t.Fatalf("Could not decode patch, unexpected error: %v", err)
	}

	res, err := patch.Apply([]byte(orig))
	if err != nil {
		t.Fatalf("Could not apply patch, unexpected error: %v", err)
	}

	if !matches(res, `{"x":[{"v":"d"},{"y":"b"},{"y":"c"}]}`) {
		t.Fatalf("Unexpected result: %s", string(res))
	}

	patchSource = `[{ "op": "replace", "path": "$.x[?(@.y=='a')]", "value": {"v":"d"} }]`

	patch, err = DecodePatch([]byte(patchSource))
	if err != nil {
		t.Fatalf("Could not decode patch, unexpected error: %v", err)
	}

	res, err = patch.Apply([]byte(orig))
	if err != nil {
		t.Fatalf("Could not apply patch, unexpected error: %v", err)
	}

	if !matches(res, `{"x":[{"v":"d"},{"y":"b"},{"y":"c"}]}`) {
		t.Fatalf("Unexpected result: %s", string(res))
	}
}

func TestMoveAttribute(t *testing.T) {
	orig := `{"x":{"a":1,"b":2},"y":{"c":3}}`
	patchSource := `[{ "op": "move", "from": "$.x.b", "path": "$.y.b" }]`

	patch, err := DecodePatch([]byte(patchSource))
	if err != nil {
		t.Fatalf("Could not decode patch, unexpected error: %v", err)
	}

	res, err := patch.Apply([]byte(orig))
	if err != nil {
		t.Fatalf("Could not apply patch, unexpected error: %v", err)
	}

	if !matches(res, `{"x":{"a":1},"y":{"c":3,"b":2}}`) {
		t.Fatalf("Unexpected result: %s", string(res))
	}

	patchSource = `[{ "op": "replace", "path": "$.x", "value": {"a":"i","b":"e"} }]`

	patch, err = DecodePatch([]byte(patchSource))
	if err != nil {
		t.Fatalf("Could not decode patch, unexpected error: %v", err)
	}

	res, err = patch.Apply([]byte(orig))
	if err != nil {
		t.Fatalf("Could not apply patch, unexpected error: %v", err)
	}

	if !matches(res, `{"x":{"a":"i","b":"e"},"y":{"c":3}}`) {
		t.Fatalf("Unexpected result: %s", string(res))
	}
}

func TestMoveEllementInArray(t *testing.T) {

	orig := `{"x":[{"y":"a"},{"y":"b"},{"y":"c"}]}`

	patchSource := `[{ "op": "move", "from": "$.x.-1", "path": "$.x.0" }]`

	patch, err := DecodePatch([]byte(patchSource))
	if err != nil {
		t.Fatalf("Could not decode patch, unexpected error: %v", err)
	}

	res, err := patch.Apply([]byte(orig))
	if err != nil {
		t.Fatalf("Could not apply patch, unexpected error: %v", err)
	}

	if !matches(res, `{"x":[{"y":"c"},{"y":"a"},{"y":"b"}]}`) {
		t.Fatalf("Unexpected result: %s", string(res))
	}
}

func TestCopyAttribute(t *testing.T) {
	orig := `{"x":{"a":1,"b":2},"y":{"c":3}}`
	patchSource := `[{ "op": "copy", "from": "$.x.b", "path": "$.y.b" }]`

	patch, err := DecodePatch([]byte(patchSource))
	if err != nil {
		t.Fatalf("Could not decode patch, unexpected error: %v", err)
	}

	res, err := patch.Apply([]byte(orig))
	if err != nil {
		t.Fatalf("Could not apply patch, unexpected error: %v", err)
	}

	if !matches(res, `{"x":{"a":1,"b":2},"y":{"c":3,"b":2}}`) {
		t.Fatalf("Unexpected result: %s", string(res))
	}
}

func TestCopyEllementInArray(t *testing.T) {

	orig := `{"x":[{"y":"a"},{"y":"b"},{"y":"c"}]}`

	patchSource := `[{ "op": "copy", "from": "$.x.-1", "path": "$.x.0" }]`

	patch, err := DecodePatch([]byte(patchSource))
	if err != nil {
		t.Fatalf("Could not decode patch, unexpected error: %v", err)
	}

	res, err := patch.Apply([]byte(orig))
	if err != nil {
		t.Fatalf("Could not apply patch, unexpected error: %v", err)
	}

	if !matches(res, `{"x":[{"y":"c"},{"y":"a"},{"y":"b"},{"y":"c"}]}`) {
		t.Fatalf("Unexpected result: %s", string(res))
	}
}

func TestTestAttribute(t *testing.T) {
	orig := `{"x":{"a":1,"b":2},"y":{"c":3}}`
	patchSource := `[{ "op": "test", "path": "$.x.b", "value": 2 }]`

	patch, err := DecodePatch([]byte(patchSource))
	if err != nil {
		t.Fatalf("Could not decode patch, unexpected error: %v", err)
	}

	_, err = patch.Apply([]byte(orig))
	if err != nil {
		t.Fatalf("Could not apply patch, unexpected error: %v", err)
	}

	patchSource = `[{ "op": "test", "path": "$.x.b", "value": 3 }]`

	patch, err = DecodePatch([]byte(patchSource))
	if err != nil {
		t.Fatalf("Could not decode patch, unexpected error: %v", err)
	}

	_, err = patch.Apply([]byte(orig))
	if err == nil {
		t.Fatalf("Expected error")
	}
}

func matches(a []byte, b string) bool {
	nodeA, err := ajson.Unmarshal(a)
	if err != nil {
		return false
	}

	nodeB, err := ajson.Unmarshal([]byte(b))
	if err != nil {
		return false
	}

	res, err := nodeA.Eq(nodeB)
	if err != nil {
		return false
	}

	return res
}
