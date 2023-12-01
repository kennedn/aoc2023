package common

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func advent_opener(filename string) ([]string, error) {
	f, err := os.Open(filename)

	lines := []string{}

	if err != nil {
		return lines, err
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

func advent_downloader(day string, year string) ([]string, error) {
	if year == "" {
		year = "2023"
	}

	filename := fmt.Sprintf("day_%s.txt", day)

	lines, err := advent_opener(filename)

	if err == nil {
		return lines, err
	}

	client := &http.Client{
		Timeout: time.Duration(1000) * time.Millisecond,
	}

	queryUrl := fmt.Sprintf("https://adventofcode.com/%s/day/%s/input", year, day)

	req, err := http.NewRequest("GET", queryUrl, nil)

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: os.Getenv("AOC_COOKIE"),
	})

	if err != nil {
		return lines, err
	}
	// Send the request and get the response
	resp, err := client.Do(req)
	if err != nil {
		return lines, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return lines, errors.New("Non 200 status code")
	}

	f, err := os.Create(filename)

	if err != nil {
		return lines, err
	}

	defer f.Close()

	io.Copy(f, resp.Body)

	lines, err = advent_opener(filename)

	return lines, err
}
