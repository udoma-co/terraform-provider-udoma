package jsonpatch

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/spyzhov/ajson"
)

type operation struct {
	Op    string          `json:"op"`
	From  string          `json:"from"`
	Path  string          `json:"path"`
	Value json.RawMessage `json:"value"`
}

type Patch []operation

// DecodePatch decodes a JSON patch document.
func DecodePatch(patch []byte) (Patch, error) {
	var ops Patch
	if err := json.Unmarshal(patch, &ops); err != nil {
		return nil, err
	}

	if err := ops.validate(); err != nil {
		return nil, err
	}

	return ops, nil
}

// validate checks if the patch is valid.
func (p Patch) validate() error {
	for _, op := range p {
		// check if operation value is correct
		switch op.Op {
		case "add", "remove", "replace", "move", "copy", "test":
		default:
			return fmt.Errorf("unsupported operation: %s", op.Op)
		}

		// check if path is valid
		if op.Path == "" {
			return fmt.Errorf("path is missing")
		}
		_, err := ajson.ParseJSONPath(op.Path)
		if err != nil {
			return fmt.Errorf("could not parse path %q: %w", op.Path, err)
		}

		// check if from is valid
		if op.Op == "move" || op.Op == "copy" {
			if op.From == "" {
				return fmt.Errorf("from is missing")
			}
			_, err := ajson.ParseJSONPath(op.From)
			if err != nil {
				return fmt.Errorf("could not parse from %q: %w", op.From, err)
			}
		}

		// check if value is valid
		if op.Op == "add" || op.Op == "replace" || op.Op == "test" {
			if len(op.Value) == 0 {
				return fmt.Errorf("value is missing")
			}
		}
	}
	return nil
}

// Apply applies the patch to the provided document.
func (p Patch) Apply(doc []byte) ([]byte, error) {
	root, err := ajson.UnmarshalSafe(doc)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal document: %w", err)
	}

	var path []string
	for _, op := range p {

		if op.Path != "" {
			path, err = ajson.ParseJSONPath(op.Path)
			if err != nil {
				return nil, fmt.Errorf("could not parse path %q: %w", op.Path, err)
			}
		}

		switch op.Op {
		case "add":
			if err := add(root, path, op.Value); err != nil {
				return nil, fmt.Errorf("could not apply operation %q: %w", op.Op, err)
			}
		case "remove":
			if err := remove(root, path); err != nil {
				return nil, fmt.Errorf("could not apply operation %q: %w", op.Op, err)
			}
		case "move":
			from, err := ajson.ParseJSONPath(op.From)
			if err != nil {
				return nil, fmt.Errorf("could not parse from %q: %w", op.Path, err)
			}

			val, err := getVal(root, from)
			if err != nil {
				return nil, fmt.Errorf("could not apply operation %q: %w", op.Op, err)
			}
			if err = add(root, path, val); err != nil {
				return nil, fmt.Errorf("could not apply operation %q: %w", op.Op, err)
			}
			if err := remove(root, from); err != nil {
				return nil, fmt.Errorf("could not apply operation %q: %w", op.Op, err)
			}
		case "copy":
			from, err := ajson.ParseJSONPath(op.From)
			if err != nil {
				return nil, fmt.Errorf("could not parse from %q: %w", op.Path, err)
			}

			val, err := getVal(root, from)
			if err != nil {
				return nil, fmt.Errorf("could not apply operation %q: %w", op.Op, err)
			}
			if err = add(root, path, val); err != nil {
				return nil, fmt.Errorf("could not apply operation %q: %w", op.Op, err)
			}
		case "replace":
			if err := replace(root, path, op.Value); err != nil {
				return nil, fmt.Errorf("could not apply operation %q: %w", op.Op, err)
			}
		case "test":
			if err := test(root, path, op.Value); err != nil {
				return nil, fmt.Errorf("could not apply operation %q: %w", op.Op, err)
			}
		default:
			return nil, fmt.Errorf("unsupported operation: %s", op.Op)
		}
	}

	result, err := ajson.Marshal(root)
	if err != nil {
		return nil, fmt.Errorf("could not marshal document: %w", err)
	}

	return result, nil
}

func getVal(root *ajson.Node, path []string) (json.RawMessage, error) {
	if len(path) == 0 {
		return ajson.Marshal(root)
	}

	nodes, err := ajson.ApplyJSONPath(root, path)
	if err != nil {
		return nil, fmt.Errorf("could not find node: %w", err)
	}

	if len(nodes) == 0 {
		return nil, fmt.Errorf("could not find node")
	}

	return ajson.Marshal(nodes[0])
}

func getNode(root *ajson.Node, path []string) (*ajson.Node, error) {
	if len(path) == 0 {
		return root, nil
	}

	nodes, err := ajson.ApplyJSONPath(root, path)
	if err != nil {
		return nil, fmt.Errorf("could not find node: %w", err)
	}

	if len(nodes) == 0 {
		return nil, fmt.Errorf("could not find node")
	}

	return nodes[0], nil
}

func add(root *ajson.Node, path []string, value json.RawMessage) error {

	newValue, err := ajson.Unmarshal(value)
	if err != nil {
		return fmt.Errorf("could not unmarshal value: %w", err)
	}

	if len(path) == 0 {
		root.AppendObject("", newValue)
		return nil
	} else if len(path) == 1 {
		root.AppendObject(path[0], newValue)
		return nil
	}

	nodes, err := ajson.ApplyJSONPath(root, path[:len(path)-1])
	if err != nil {
		return fmt.Errorf("could not find node: %w", err)
	}

	if len(nodes) == 0 {
		return fmt.Errorf("could not find node")
	}

	node := nodes[0]

	if node == nil {
		return fmt.Errorf("node is nil")
	}

	key := path[len(path)-1]

	// append to object
	if node.IsObject() {
		return node.AppendObject(key, newValue)
	}

	// append to array
	if node.IsArray() {
		// "-" and empty string mean append to end of array
		if key == "-" || key == "" {
			return node.AppendArray(newValue)
		}

		// else, try to parse key as index
		idx, err := strconv.Atoi(key)
		if err != nil {
			return fmt.Errorf("could not parse key as index: %w", err)
		}

		l := node.Size()
		if idx < 0 {
			idx = l + idx
		}

		if idx < 0 || idx > l {
			return fmt.Errorf("index out of bounds")
		}

		if idx == l {
			return node.AppendArray(newValue)
		}

		arr := make([]*ajson.Node, l+1)
		for i := 0; i < idx; i++ {
			arr[i], err = node.GetIndex(i)
			if err != nil {
				return fmt.Errorf("could not get index: %w", err)
			}
		}
		arr[idx] = newValue
		for i := idx; i < l; i++ {
			arr[i+1], err = node.GetIndex(i)
			if err != nil {
				return fmt.Errorf("could not get index: %w", err)
			}
		}

		return node.SetArray(arr)
	}

	return fmt.Errorf("node is not an object or array")
}

func remove(root *ajson.Node, path []string) error {

	if len(path) == 0 {
		return fmt.Errorf("cannot remove root node")
	} else {
		nodes, err := ajson.ApplyJSONPath(root, path)
		if err != nil {
			return fmt.Errorf("could not find node: %w", err)
		}
		for _, node := range nodes {
			node.Delete()
		}
	}

	return nil
}

func replace(root *ajson.Node, path []string, value json.RawMessage) error {
	node, err := getNode(root, path)
	if err != nil {
		return fmt.Errorf("could not apply replace, could not find node: %w", err)
	}

	newValue, err := ajson.Unmarshal(value)
	if err != nil {
		return fmt.Errorf("could not unmarshal value: %w", err)
	}

	if err := node.Set(newValue); err != nil {
		return fmt.Errorf("could not apply replace, could not update node: %w", err)
	}

	return nil
}

func test(root *ajson.Node, path []string, value json.RawMessage) error {

	node, err := getNode(root, path)
	if err != nil {
		return fmt.Errorf("could not check test, could not find node: %w", err)
	}

	expectedValue, err := ajson.Unmarshal(value)
	if err != nil {
		return fmt.Errorf("could not unmarshal value: %w", err)
	}

	res, err := node.Eq(expectedValue)
	if err != nil {
		return fmt.Errorf("could not check test, could not compare node: %w", err)
	}

	if !res {
		return fmt.Errorf("value test failed, node does not match expected value")
	}

	return nil
}
