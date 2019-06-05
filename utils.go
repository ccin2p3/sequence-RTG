package sequence

import (
	"encoding/json"
	"os"
	"strings"
)

func ScanMessage(scanner *Scanner, data string, format string) (Sequence, error) {
	var (
		seq Sequence
		err error
		pos []int
	)

	if testJson(data){
		seq, err = scanner.ScanJson(data)
	} else {
		switch format {
		case "json":
			seq, err = scanner.ScanJson(data)

		default:
			seq, err = scanner.Scan(data, false, pos)
		}
	}
	return seq, err
}

func testJson(data string)bool{
	data = strings.TrimSpace(data)
	var js string
	if data[:1] == "{" && data[len(data)-1:] == "}"{
		//try to marshall the json
		return json.Unmarshal([]byte(data), &js) == nil
	}
	return false
}


func BuildParser(patfile string) *Parser {
	parser := NewParser()

	if patfile == "" {
		return parser
	}

	var files []string
	var pos []int

	if fi, err := os.Stat(patfile); err != nil {
		logger.HandleFatal(err.Error())
	} else if fi.Mode().IsDir() {
		files, err = GetDirOfFiles(patfile)
	} else {
		files = append(files, patfile)
	}

	scanner := NewScanner()

	for _, file := range files {
		// Open pattern file
		pscan, pfile, err:= OpenInputFile(file)
		defer pfile.Close()
		if err != nil {
			logger.HandleFatal(err.Error())
		}

		for pscan.Scan() {
			line := pscan.Text()
			if len(line) == 0 || line[0] == '#' {
				continue
			}

			seq, err := scanner.Scan(line, true, pos)
			if err != nil {
				logger.HandleError(err.Error())
			}

			if err := parser.Add(seq); err != nil {
				logger.HandleError(err.Error())
			}
		}
	}

	return parser
}

func BuildParserFromDb(serviceid string) *Parser {
	parser := NewParser()
	scanner := NewScanner()
	db, ctx := OpenDbandSetContext()
	defer db.Close()
	//load all patterns from the database
	pmap := GetPatternsFromDatabaseByService(db, ctx, serviceid)
	for _, ar := range pmap {
		pos := SplitToInt(ar.TagPositions, ",")
		seq, err := scanner.Scan(ar.Pattern, true, pos)
		if err != nil {
			logger.HandleError(err.Error())
		}

		if err := parser.Add(seq); err != nil {
			logger.HandleError(err.Error())
		}
	}
	return parser
}