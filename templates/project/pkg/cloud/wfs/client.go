/**
 * Created by Wangwei on 2019-03-27 11:20.
 */

package wfs

// web应用
type Config struct {
	WfsHost string `toml:"wfs_host"`
}

func Upload(fileData []byte, filename string, config *Config) (url string) {
	uploadPath := config.WfsHost + "/thrift"
	serverPath := config.WfsHost + "/r/"

	wfsClient := WfsClient{uploadPath}
	wfsClient.PostFile(fileData, filename, "")

	return serverPath + filename
}
