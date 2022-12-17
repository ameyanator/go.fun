package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
If a large file is being downloaded and a timeout occurs, starting the download from
the beginning isn’t ideal. This is becoming truer with the growth of file sizes. In many
cases, files are gigabytes or larger. It’d be nice to avoid the extra bandwidth use and
time to redownload data.
*/

func main() {
	file, err := os.Create("file.zip")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	location := "https://example.com/file.zip"
	err = download(location, file, 100)
	if err != nil {
		fmt.Println(err)
		return
	}
	fi, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Got it with", fi.Size(), "bytes downloaded")
}

func download(location string, file *os.File, retries int64) error {
	req, err := http.NewRequest(http.MethodGet, location, nil)
	if err != nil {
		return err
	}
	fi, err := file.Stat()
	if err != nil {
		return err
	}
	current := fi.Size()
	if current > 0 {
		start := strconv.FormatInt(current, 10)
		req.Header.Set("Range", "bytes="+start+"-")
	}
	cc := &http.Client{Timeout: 5 * time.Minute}
	res, err := cc.Do(req)
	if err != nil && hasTimedOut(err) {
		if retries > 0 {
			return download(location, file, retries-1)
		}
		return err
	} else if err != nil {
		return err
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		errFmt := "Unsuccess HTTP request. Status: %s"
		return fmt.Errorf(errFmt, res.Status)
	}
	if res.Header.Get("Accept-Ranges") != "bytes" {
		retries = 0
	}
	_, err = io.Copy(file, res.Body)
	if err != nil && hasTimedOut(err) {
		if retries > 0 {
			return download(location, file, retries-1)
		}
		return err
	} else if err != nil {
		return err
	}
	return nil
}

func hasTimedOut(err error) bool {
	switch err := err.(type) {
	case *url.Error:
		//url.Error may be caused by an underlying net error that can be checked for a timeout
		if err, ok := err.Err.(net.Error); ok && err.Timeout() {
			return true
		}
	case net.Error:
		// Look for timeout detected by net package
		if err.Timeout() {
			return true
		}
	case *net.OpError:
		// Look for timeout detected by net package
		if err.Timeout() {
			return true
		}
	}
	errTxt := "use of closed network connection:"
	// check without custom Timeout set

	if err != nil && strings.Contains(err.Error(), errTxt) {
		return true
	}
	return false
}

/*
The timeout is set to five minutes. This can be tuned for your application. A
shorter or longer timeout may provide better performance in your environ-
ment. For example, if you’re downloading files that typically take longer than
five minutes, a timeout longer than most files take will limit the number of
HTTP requests needed for a normal download.
*/

/*
If a hash of a file is easily available, a check could be put in to make sure that
the final download matches the hash. This integrity check can improve trust in
the final download, even if it takes multiple attempts to download the file.
*/
