package converter

import (
	"encoding/json"
	"os"
	"strings"
	"strconv"
	log "github.com/sirupsen/logrus"
)

type JsonConverter struct{
	ConverterBase
}

func (c JsonConverter) Convert() []byte{
	// convert to JSON
	log.Debug("Converting to JSON format")
	data, err := json.MarshalIndent(c.Root.TxsdIodefDocument, "", "    ")

	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	// Unescape unicode characters (\uXXXX) in JSON output
	log.Debug("Unescape unicode character in JSON output")
	data, err = unescapeUnicodeCharactersInJSON(data)
	
	if err != nil{
		log.Error(err)
		os.Exit(1)
	}

	return data
}


/*
* Unescape unicode characters in JSON data
*
*/
func unescapeUnicodeCharactersInJSON(jsonRaw json.RawMessage) (json.RawMessage, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(jsonRaw)), `\\u`, `\u`, -1))
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	return []byte(str), nil
}