package downloader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/gorilla/mux"
)

type DownloadConfigs struct {
	URL     string `json:"url,omitempty"`
	Threads int    `json:"threads,omitempty"`
}

var wg sync.WaitGroup
var defaultThreadCount = 2

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", Download).Methods("POST")
	return router
}

// Download function accepts the requests for the resource to be downloaded
func Download(w http.ResponseWriter, r *http.Request) {
	// parse query params
	decoder := json.NewDecoder(r.Body)
	var downloadConfigs DownloadConfigs
	err := decoder.Decode(&downloadConfigs)
	if err != nil {
		log.Fatal(err)
	}

	resourceURL := downloadConfigs.URL
	threadCount := downloadConfigs.Threads

	params, _ := json.Marshal(downloadConfigs)
	fmt.Printf("\npid:[%d]\t\t%s params: %s\n", os.Getpid(), r.RemoteAddr, string(params))
	log.Printf("pid:[%d]\t\t%s params: %s\n", os.Getpid(), r.RemoteAddr, string(params))

	// Handle cases like thread = 0 or param not present
	if threadCount == 0 {
		threadCount = defaultThreadCount
	}

	// using url for the name of the file
	splitURL := strings.Split(resourceURL, "/")
	fileName := splitURL[len(splitURL)-1]

	res, err := http.Head(resourceURL)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	fileSize, err := strconv.Atoi(res.Header.Get("content-length"))
	if err != nil {
		log.Fatal(err)
	}
	// create empty file with the same name in append mode
	createEmptyFile(fileName)

	// Calculate the chunk size based on number of threads
	subLen := fileSize / threadCount
	remainingChunk := fileSize % threadCount

	body := make([]string, threadCount)
	for i := 0; i < threadCount; i++ {
		wg.Add(1)

		start := subLen * i
		if i != 0 {
			start = (subLen * i) + 1
		}
		end := subLen * (i + 1)
		// Add the remaining bytes in the last request
		if i == threadCount-1 {
			end += remainingChunk
		}

		go func(start int, end int, i int) {
			client := &http.Client{}
			req, _ := http.NewRequest("GET", resourceURL, nil)
			rangeHeader := "bytes=" + strconv.Itoa(start) + "-" + strconv.Itoa(end)
			req.Header.Add("Range", rangeHeader)
			resp, err := client.Do(req)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()
			reader, _ := ioutil.ReadAll(resp.Body)
			body[i] = string(reader)

			wg.Done()
			writeToFile(body, fileName, i)
		}(start, end, i)
	}
	wg.Wait()
	fmt.Fprintf(w, "File is downloaded from url: %s", resourceURL)
}

func createEmptyFile(fileName string) {
	filePath := fmt.Sprintf("./downloads/%s", fileName)
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := file.Write([]byte("")); err != nil {
		log.Fatal(err)
	}
}

func writeToFile(body []string, fileName string, i int) {
	combinedString := strings.Join(body, "")
	filePath := fmt.Sprintf("./downloads/%s", fileName)
	ioutil.WriteFile(filePath, []byte(string(combinedString)), 0x777)
}
