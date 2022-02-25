package fetch

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"os"
)

func Data(url string, filename string) error {
	f, err := os.Open(filename)
	defer f.Close()
	if os.IsNotExist(err) {
		resp, err := makeReq(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		file, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			return err
		}

		fmt.Printf("successfully fetched data from %s and saved it to %s\n", url, filename)
		return nil
	} else if err != nil {
		return err
	} else {
		fmt.Printf("file already exists %s\n", filename)
		return nil
	}
}

func makeReq(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	client := http.Client{
		Jar: jar,
	}
	cookie := &http.Cookie{
		Name:  "session",
		Value: os.Getenv("AOC_COOKIE"),
	}
	req.AddCookie(cookie)
	resp, err := client.Do(req)
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Data: request failed with status code %d", resp.StatusCode)
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}
