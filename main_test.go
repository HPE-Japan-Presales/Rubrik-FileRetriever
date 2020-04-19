package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	os.Setenv("rubrik_cdm_node_ip", "172.16.14.220")
	os.Setenv("rubrik_cdm_username", "tak")
	os.Setenv("rubrik_cdm_password", "P@ssw0rd")
	os.Setenv("RUBRICK_TARGET_HOST", "rubrik-sql01.hybrid-lab.local")
	os.Setenv("RUBRICK_TARGET_FILESET", "sdk01")
	os.Setenv("RUBRICK_TARGET_HOSTOS", "Linux")
	os.Setenv("RUBRICK_TARGET_DATE", "04-20-2020 12:02 AM")
	os.Setenv("RUBRICK_TARGET_FILEPATH", "/rubrik/bu-tar01/test01.txt")
	os.Setenv("DOWNLOAD_FILE_PATH", "/tmp/test01.txt")
	os.Setenv("POST_SCRIPT_PATH", "/tmp/post.sh")

	main()
}
