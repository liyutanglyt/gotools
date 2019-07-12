/**
 * Created by Wangwei on 2019-03-20 14:49.
 */

package mapstruct

import (
	"bytes"
	"encoding/xml"
	"io"
	"goadmin/pkg/jsonutil"
	"strings"
)

func Map2Struct(input interface{}, output interface{}) {
	str := jsonutil.MarshalToString(input)
	if err := jsonutil.Unmarshal([]byte(str), output); err != nil {
		panic(err)
	}
}

func Struct2Map(obj interface{}) map[string]interface{} {
	m := map[string]interface{}{}
	str, _ := jsonutil.Json.Marshal(obj)
	if err := jsonutil.Unmarshal(str, &m); err != nil {
		panic(err)
	}

	return m
}

func XmlToMap(xmlData []byte) map[string]interface{} {
	decoder := xml.NewDecoder(bytes.NewReader(xmlData))
	m := make(map[string]interface{})
	var token xml.Token
	var err error
	var k string
	for token, err = decoder.Token(); err == nil; token, err = decoder.Token() {
		if v, ok := token.(xml.StartElement); ok {
			k = v.Name.Local
			continue
		}
		if v, ok := token.(xml.CharData); ok {
			data := string(v.Copy())
			if strings.TrimSpace(data) == "" {
				continue
			}
			m[k] = data
		}
	}

	if err != nil && err != io.EOF {
		panic(err)
	}
	return m
}
