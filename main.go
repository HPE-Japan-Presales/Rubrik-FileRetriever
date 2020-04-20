package main

import (
	"crypto/tls"
	"io"
	"net/http"
	"os"
	"os/exec"

	"github.com/rubrikinc/rubrik-sdk-for-go/rubrikcdm"
	utils "github.com/rubrik-filesets-restore/utils"
)

var logging = &utils.Logging{}

func main() {
	logging.Infoln("Start Rubrik Retriever")
	logging.Infof("rubrik_cdm_node_ip: %s", os.Getenv("rubrik_cdm_node_ip"))
	logging.Infof("rubrik_cdm_username: %s", os.Getenv("rubrik_cdm_username"))
	logging.Infoln("rubrik_cdm_password: ************")
	logging.Infof("RUBRICK_TARGET_HOST: %s", os.Getenv("RUBRICK_TARGET_HOST"))
	logging.Infof("RUBRICK_TARGET_FILESET: %s", os.Getenv("RUBRICK_TARGET_FILESET"))
	logging.Infof("RUBRICK_TARGET_HOSTOS: %s", os.Getenv("RUBRICK_TARGET_HOSTOS"))
	logging.Infof("RUBRICK_TARGET_DATE: %s", os.Getenv("RUBRICK_TARGET_DATE"))
	logging.Infof("RUBRICK_TARGET_FILEPATH: %s", os.Getenv("RUBRICK_TARGET_FILEPATH"))
	logging.Infof("DOWNLOAD_FILE_PATH: %s", os.Getenv("DOWNLOAD_FILE_PATH"))
	logging.Infof("POST_SCRIPT_PATH: %s", os.Getenv("POST_SCRIPT_PATH"))

	logging.Infoln("Trying to connect Rubrik Cluster")
	rubrik, err := rubrikcdm.ConnectEnv()
	if err != nil {
		logging.Fatalln("Failed to connect Rubrik Cluster")
	}

	summary, err := rubrik.Get("v1", "/cluster/me")
	if err != nil {
		logging.Fatalln("Failed to retrieve Rubrik infomation")
	}
	s := summary.(map[string]interface{})
	logging.Infof("Accessed %v Ver %v", s["name"], s["version"])

	jobUrl, err := rubrik.RecoverFileDownload(
		os.Getenv("RUBRICK_TARGET_HOST"),
		os.Getenv("RUBRICK_TARGET_FILESET"),
		os.Getenv("RUBRICK_TARGET_HOSTOS"),
		os.Getenv("RUBRICK_TARGET_FILEPATH"),
		os.Getenv("RUBRICK_TARGET_DATE"))

	if err != nil {
		logging.Fatalf("Failed to execute file download job: %v", err)
	}
	logging.Infoln(jobUrl)
	resp, err := rubrik.JobStatus(jobUrl)
	if err != nil {
		logging.Fatalf("Job failed: %v", err)
	}
	d_path := resp.(map[string]interface{})["links"].([]interface{})[1].(map[string]interface{})["href"].(string)
	d_url := "https://" + os.Getenv("rubrik_cdm_node_ip") + "/" + d_path
	logging.Infof(d_url)
	err = downloadFile(d_url, os.Getenv("DOWNLOAD_FILE_PATH"))
	if err != nil {
		logging.Fatalf("Download failed: %v", err)
	}

	ps_path := os.Getenv("POST_SCRIPT_PATH")
	if ps_path != "" {
		err := doPostScript(ps_path)
		if err != nil {
			logging.Fatalf("Post script failed: %v", err)
		}
	}
	logging.Infoln("All tasks completed!!")
}

func downloadFile(url string, filepath string) error {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func doPostScript(path string) error {
	out, err := exec.Command(path).Output()

	if err != nil {
		return err
	}

	logging.Infof("Post script output\n%s", string(out))
	return nil
}
