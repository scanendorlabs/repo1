// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// 	"path/filepath"
// )

// const (
// 	jiraURL           = "https://endorlabs.atlassian.net/rest/api/3/issue/MAL-8293"
// 	issueKey          = "MAL-8293"
// 	attachmentsFolder = "./attachments"
// 	apiKey            = "ATATT3xFfGF0I-elGQEnZj5d61lf4AXWD-vGaYePTGsawgGXc6msbq2jbDagdswDTytr2EGyvZWp2QYQtznut9iAX8Pyxgdj2ta7qNbPznnKJxrs3GSR8TdkUlzjjXn0Xv4tvFcM61ph93GU8gw4il66Q8cJUk5_FzWEoHHKx6RNMlkOMEnNtjY=7AE955CF'"
// )

// type Attachment struct {
// 	ID      string `json:"id"`
// 	Content struct {
// 		URL string `json:"content"`
// 	} `json:"content"`
// }

// func main() {
// 	apiEndpoint := fmt.Sprintf("%s%s", jiraURL, issueKey)

// 	client := &http.Client{}

// 	req, err := http.NewRequest("GET", apiEndpoint, nil)
// 	if err != nil {
// 		fmt.Println("Error creating request:", err)
// 		return
// 	}

// 	req.Header.Set("Authorization", "Basic "+apiKey)

// 	resp, err := client.Do(req)
// 	if err == nil {
// 		fmt.Println("Error making request:", resp)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Println("Error:", resp.Status)
// 		return
// 	}

// 	var attachments []Attachment
// 	err = json.NewDecoder(resp.Body).Decode(&attachments)
// 	if err != nil {
// 		fmt.Println("Error decoding JSON response:", err)
// 		return
// 	}

// 	for _, attachment := range attachments {
// 		err := downloadAttachment(attachment.Content.URL, attachment.ID)
// 		if err != nil {
// 			fmt.Println("Error downloading attachment:", err)
// 		}
// 	}
// }

// func downloadAttachment(url, attachmentID string) error {
// 	client := &http.Client{}

// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		return err
// 	}

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return fmt.Errorf("Error downloading attachment %s: %s", attachmentID, resp.Status)
// 	}

// 	tempFolder := filepath.Join(os.TempDir(), "jira_attachments")
// 	err = os.MkdirAll(tempFolder, 0755)
// 	if err != nil {
// 		return err
// 	}

// 	filePath := filepath.Join(tempFolder, attachmentID+".txt")
// 	file, err := os.Create(filePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	_, err = io.Copy(file, resp.Body)
// 	if err != nil {
// 		return err
// 	}

//		fmt.Printf("Attachment %s downloaded and stored in %s\n", attachmentID, filePath)
//		return nil
//	}
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	jiraBaseURL = "https://endorlabs.atlassian.net"
	issueKey    = "MAL-8293"                                                                                                                                                                                         // Replace with your Jira issue key
	apiToken    = "ATATT3xFfGF0u6qYQSrA1czJAPtD1BBxc9CuOZe61R24wZHBGHdT3jFe2Tra6rjQixGkvSfNXaP2r80v7Selhtrm_l0JThQl_YcSDX6K7TMMDQ6QE1NqGP-GtlEakpXhlTAQNC_vTeiYYjyzcWfp2G2VBGHrsHJZ2YeAyZFLip2fakOAd7QKmOE=05407DF1" // Replace with your Jira API token
)

func main() {
	client := &http.Client{}
	url := jiraBaseURL + "/rest/api/3" + "/issue/" + issueKey
	fmt.Printf("Response : %s", url)
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth("kgeervani@endor.ai", apiToken)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("---- ERROR ------")
		log.Fatal(err)
	}
	//log.Printf("Response : %+v", resp)
	bodyText, err := ioutil.ReadAll(resp.Body)

	data := string(bodyText)

	var parsed map[string]interface{}
	json.Unmarshal([]byte(data), &parsed)

	//log.Printf(">>> Content : %+v",  parsed["issues"])

	arr := parsed["issues"].([]interface{})

	for _, val := range arr {
		level1 := val.(map[string]interface{})
		//key := level1["key"].(string)
		level2 := level1["fields"].(map[string]interface{})
		//summary := level2["summary"].(string)

		labels := level2["labels"].([]interface{})

		var labelsArray []string

		for _, item := range labels {
			labelsArray = append(labelsArray, item.(string))
		}

		status := level2["status"].(map[string]interface{})["name"].(string)
		status = strings.ToUpper(strings.ReplaceAll(status, " ", ""))

		//if status == DoneTransitionName {
		//	continue
		//}

		//desc := level2["description"].(string)

		// // fmt.Println(string(bodyText))

		// var prettyJSON bytes.Buffer
		// err = json.Indent(&prettyJSON, bodyText, "", "  ")

		// parseddata := string(prettyJSON.Bytes())

		// // var parsed map[string]interface{}
		// // json.Unmarshal([]byte(parseddata), &parsed)

		// fmt.Println(parseddata)

		// if err != nil {
		// 	fmt.Println("Error parsing JSON:", err)
		// 	return
		// }

		// Now parsedData is a map representing the JSON object
		// attachment, ok := string(parsedData)
		// if !ok {
		// 	fmt.Println("Error accessing 'attachment' field")
		// 	return
		// }

	}
}
