/**
 * Created by Wangwei on 2019-04-09 17:25.
 */

package utils

import "goadmin/pkg/jsonutil"

func PrettyStruct(v interface{}) string {
	b, _ := jsonutil.Json.MarshalIndent(v, "", "  ")
	return string(b)
}
