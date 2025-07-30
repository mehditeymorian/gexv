package extractor

import (
	"encoding/csv"
	"errors"
	"io/ioutil"
	"os"

	"github.com/dlclark/regexp2"
	"github.com/mehditeymorian/gexv/config"
)

// GetSource returns the input text from a file or inline
func GetSource(filePath, text string) (string, error) {
	if filePath != "" {
		b, err := ioutil.ReadFile(filePath)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}
	if text != "" {
		return text, nil
	}
	return "", errors.New("either --file or --text must be specified")
}

// ExtractToCSV applies the regex and writes named groups to CSV
func ExtractToCSV(cfg *config.Config, src, outPath string, includeMatchedSection bool) error {
	// Compile regex
	opts := regexp2.None
	for _, f := range cfg.Flags {
		switch f {
		case 'i':
			opts |= regexp2.IgnoreCase
		case 'm':
			opts |= regexp2.Multiline
		case 's':
			opts |= regexp2.Singleline
		}
	}
	re, err := regexp2.Compile(cfg.Pattern, opts)
	if err != nil {
		return err
	}
	match, err := re.FindStringMatch(src)
	if err != nil {
		return err
	}
	if match == nil {
		return nil // no matches
	}
	// Headers = group names
	grp := match.Groups()
	var headers []string
	for _, g := range grp {
		if g.Name != "" && !(!includeMatchedSection && g.Name == "0") {
			headers = append(headers, g.Name)
		}
	}
	// Create CSV
	file, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer file.Close()
	w := csv.NewWriter(file)
	defer w.Flush()
	w.Write(headers)
	// Iterate matches
	for match != nil {
		row := make([]string, len(headers))
		idx := map[string]int{}
		for i, h := range headers {
			idx[h] = i
		}
		for _, g := range match.Groups() {
			if g.Name != "" && !(!includeMatchedSection && g.Name == "0") {
				row[idx[g.Name]] = g.String()
			}
		}
		w.Write(row)
		match, err = re.FindNextMatch(match)
		if err != nil {
			return err
		}
	}
	return nil
}
