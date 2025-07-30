# GexV (REGEX CSV)

Extract named regex groups from text or files and save as CSV.

## Installation

```bash
go install github.com/mehditeymorian/gexv@latest
```

## Usage

```bash
gexv --config=pattern.json --file=input.txt --output=results.csv
```

Flags:
- `-c, --config string`             Path to JSON config file
- `-f, --file string`               Input file path
- `-g, --flags string`              list of flags to pass to regex. such as <gm> for global and multiline matching
- `-h, --help`                      help for regex-extractor
- `-i, --include-matched-section`   Include matched section with whole regex
- `-o, --output string`             Output CSV file path (default "output.csv")
- `-p, --pattern string`            Pattern to match against content
- `-t, --text string`               Inline input text
- `-v, --version`                   version for regex-extractor

Config:
```json
{
	"pattern": "",
	"flags": ""
}
```