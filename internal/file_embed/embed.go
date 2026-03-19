package fileembed

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/bluekeyes/go-gitdiff/gitdiff"
)

// ReadFileWithEmbeds reads the file with the given name and embeds any nested
// files, such as JSON, PNG, JPG, and JS files, into the content. For JS files,
// it applies any diff patches that were provided that match the file name.
func ReadFileWithEmbeds(fileName string, diffPatches []*gitdiff.File) ([]byte, error) {

	dirName := filepath.Dir(fileName)

	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %w", err)
	}

	if strings.HasSuffix(fileName, ".yaml") || strings.HasSuffix(fileName, ".yml") {
		data, err = embedFilesYaml(data, dirName, diffPatches)
		if err != nil {
			return nil, fmt.Errorf("could not embed nested files: %w", err)
		}
		return data, nil
	}

	if strings.HasSuffix(fileName, ".json") {
		data, err = embedFilesJson(data, dirName, diffPatches)
		if err != nil {
			return nil, fmt.Errorf("could not embed nested files: %w", err)
		}
		return data, nil
	}

	return nil, fmt.Errorf("unknown file format, only support .yaml, .yml, .json")
}

// embedFilesJson finds occurrences of "src:.*.json", "src:.*.png", "src:.*.jpg",
// and "src:.*.js" in the data and replaces them with the content of the
// respective files. For JSON files, it embeds the content as a string, for
// PNG and JPG  files, it embeds the content as base64 encoded strings, and
// for JS files, it embeds the content as a single line string with escaped
// characters, after applying all patches that match the file name. JSON files
// can also be embedded as a string, rather than a nested JSON structure by
// adding "|string" at the end of the file name. It also recursively embeds
// any nested files found within the JSON content. It returns the modified
// data or an error if any file could not be read.
func embedFilesJson(data []byte, dirName string, diffPatches []*gitdiff.File) ([]byte, error) {

	iteratePattern := func(pattern string, callback func(fileName string, start int, fileData []byte) ([]byte, error)) error {
		re := regexp.MustCompile(pattern)
		for {
			match := re.FindIndex(data)
			if match == nil {
				break
			}
			start := match[0]
			end := match[1]
			fileName := string(data[start+5 : end-1]) // +5 to strip '"src:' prefix and -1 to strip trailing quote
			filePath := filepath.Join(dirName, fileName)
			fileData, err := os.ReadFile(filePath)
			if err != nil {
				return fmt.Errorf("could not read file: %w", err)
			}

			embeddedContent, err := callback(fileName, match[0], fileData)
			if err != nil {
				return err
			}

			data = append(data[:start], append(embeddedContent, data[end:]...)...)
		}
		return nil
	}

	// find occurrences of "src:.*.json" and replace with the content of the file
	if err := iteratePattern(`"src:.*\.json"`, func(fileName string, _ int, fileData []byte) ([]byte, error) {
		return embedFilesJson(fileData, dirName, diffPatches)
	}); err != nil {
		return nil, err
	}

	// find occurrences of "src:.*.png" and replace with
	// the base64 encoded content of the file
	if err := iteratePattern(`"src:.*\.(png|jpg|jpeg)"`, func(fileName string, _ int, fileData []byte) ([]byte, error) {
		base64Data := base64.StdEncoding.EncodeToString(fileData)
		return []byte(`"` + base64Data + `"`), nil
	}); err != nil {
		return nil, err
	}

	// do the same for JS files, but make sure content is escaped and put in one line
	if err := iteratePattern(`"src:.*\.js"`, func(fileName string, _ int, fileData []byte) ([]byte, error) {

		// apply any diff patches that match the file name
		for _, patch := range diffPatches {
			if !strings.HasSuffix(fileName, patch.OldName) {
				// skip patches that do not match the file name
				continue
			}

			var output bytes.Buffer
			if err := gitdiff.Apply(&output, bytes.NewReader(fileData), patch); err != nil {
				return nil, fmt.Errorf("could not apply diff patch to file %s: %w", fileName, err)
			}
			fileData = output.Bytes()
		}

		sanitized := embedString(fileData, false)

		return sanitized, nil
	}); err != nil {
		return nil, err
	}

	return data, nil
}

// embedFilesJson finds occurrences of "src:.*.json", "src:.*.png", "src:.*.jpg",
// and "src:.*.js" in the data and replaces them with the content of the
// respective files. For JSON files, it embeds the content as a string, for
// PNG and JPG  files, it embeds the content as base64 encoded strings, and
// for JS files, it embeds the content as a single line string with escaped
// characters, after applying all patches that match the file name. JSON files
// can also be embedded as a string, rather than a nested JSON structure by
// adding "|string" at the end of the file name. It also recursively embeds
// any nested files found within the JSON content. It returns the modified
// data or an error if any file could not be read.
func embedFilesYaml(data []byte, dirName string, diffPatches []*gitdiff.File) ([]byte, error) {

	// for yaml files we need to remove comments before processing embeds
	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		if before, _, ok := strings.Cut(line, "#"); ok {
			lines[i] = before
		}
	}
	data = []byte(strings.Join(lines, "\n"))

	// make sure there are no json files embedded in yaml files, as we don't support this
	pattern := regexp.MustCompile(`src:.*\.json`)
	if jsonEmbed := pattern.Find(data); jsonEmbed != nil {
		return nil, fmt.Errorf("embedding JSON files in YAML files is not supported: %s", string(jsonEmbed))
	}

	iteratePattern := func(pattern string, callback func(fileName string, start int, fileData []byte) ([]byte, error)) error {
		re := regexp.MustCompile(pattern)
		for {
			match := re.FindIndex(data)
			if match == nil {
				break
			}
			start := match[0]
			end := match[1]
			fileName := string(data[start+4 : end]) // +4 to strip "src:" prefix
			filePath := filepath.Join(dirName, fileName)
			fileData, err := os.ReadFile(filePath)
			if err != nil {
				return fmt.Errorf("could not read file: %w", err)
			}

			embeddedContent, err := callback(fileName, match[0], fileData)
			if err != nil {
				return err
			}

			data = append(data[:start], append(embeddedContent, data[end:]...)...)
		}
		return nil
	}

	// find occurrences of "src:.*.png" and replace with
	// the base64 encoded content of the file
	if err := iteratePattern(`src:.*\.(png|jpg|jpeg)`, func(fileName string, _ int, fileData []byte) ([]byte, error) {
		base64Data := base64.StdEncoding.EncodeToString(fileData)
		return []byte(base64Data), nil
	}); err != nil {
		return nil, err
	}

	// do the same for JS files, but make sure content is escaped and put in one line
	if err := iteratePattern(`src:.*\.js`, func(fileName string, pos int, fileData []byte) ([]byte, error) {

		// apply any diff patches that match the file name
		for _, patch := range diffPatches {
			if !strings.HasSuffix(fileName, patch.OldName) {
				// skip patches that do not match the file name
				continue
			}

			var output bytes.Buffer
			if err := gitdiff.Apply(&output, bytes.NewReader(fileData), patch); err != nil {
				return nil, fmt.Errorf("could not apply diff patch to file %s: %w", fileName, err)
			}
			fileData = output.Bytes()
		}

		if sourceContent := string(fileData); strings.Contains(sourceContent, "\n") {
			// for multi-line content, use YAML literal block scalar (|) to preserve newlines
			// extract only the leading whitespace of the current line (not the key)
			indentation := getIndentation(data, pos)
			indentedContent := indentContent(sourceContent, indentation)
			fileData = []byte("|\n" + indentedContent + "\n")
		}

		return fileData, nil
	}); err != nil {
		return nil, err
	}

	if err := iteratePattern(`src:.*\.yaml`, func(fileName string, pos int, fileData []byte) ([]byte, error) {

		newDirName := filepath.Dir(filepath.Join(dirName, fileName))
		fileData, err := embedFilesYaml(fileData, newDirName, diffPatches)
		if err != nil {
			return nil, fmt.Errorf("could not embed nested files: %w", err)
		}

		content := strings.TrimRight(string(fileData), "\n")

		if strings.Contains(content, "\n") {
			// for multi-line content, use YAML literal block scalar (|) to preserve newlines
			// extract only the leading whitespace of the current line (not the key)
			indentation := getIndentation(data, pos)
			indentedContent := indentContent(content, indentation)
			fileData = []byte("\n" + indentedContent + "\n")
		} else {
			fileData = []byte(content)
		}

		return fileData, nil
	}); err != nil {
		return nil, err
	}

	return data, nil
}

func getIndentation(data []byte, offset int) string {
	lineStart := strings.LastIndex(string(data[:offset]), "\n") + 1
	var leadingWhitespace strings.Builder
	for _, c := range string(data[lineStart:offset]) {
		if c == ' ' || c == '\t' {
			leadingWhitespace.WriteString(string(c))
		} else {
			break
		}
	}
	contentIndent := leadingWhitespace.String() + "  "

	return contentIndent
}

func indentContent(content string, indent string) string {
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		if line != "" {
			lines[i] = indent + line
		}
	}
	return strings.Join(lines, "\n")
}

// embedString takes a byte slice, replaces newlines with \n, escapes quotes,
// and optionally replaces tabs with \t. It returns the modified byte slice
func embedString(in []byte, replaceTabs bool) []byte {

	ret := string(in)
	ret = regexp.MustCompile(`\r?\n`).ReplaceAllString(ret, `\n`)
	ret = regexp.MustCompile(`"`).ReplaceAllString(ret, `\"`)
	if replaceTabs {
		ret = regexp.MustCompile(`\\t`).ReplaceAllString(ret, `\\t`)
	}
	ret = `"` + ret + `"`

	return []byte(ret)
}
