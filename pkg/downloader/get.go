// Package downloader manages grabbing CSV from Concept2 website
// Simply grab a downloader and either call GetSeasons or GetAllSeasons
//   path, _ := filepath.Abs("./data")
//   dl := downloader.NewDownloader("myuser", "mypassword", path)
//   dl.SetHeadless(false)
//   dl.GetSeasons([]string{"2021"})
package downloader

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	log "github.com/sirupsen/logrus"
)

// Downloader manages the download options
type Downloader struct {
	// Username Concept 2 username
	Username string
	// Password Concept 2 password
	Password string
	// Path The location where data files will be saved to
	Path     string
	headless bool
}

// NewDownloader generates a new downloader with default options
func NewDownloader(username string, password string, path string) Downloader {
	return Downloader{
		Username: username,
		Password: password,
		Path:     path,
		headless: true,
	}
}

// SetHeadless sets whether the browser should be headless or not
func (dl *Downloader) SetHeadless(headless bool) {
	dl.headless = headless
}

// GetAllSeasons downloads all of the available seasons
func (dl Downloader) GetAllSeasons() {
	ctx, cancel := (&dl).login()
	defer cancel()
	navigateToHistory(ctx)

	var foundSeasons []*cdp.Node
	err := chromedp.Run(ctx,
		chromedp.Nodes(".history-section>table tbody tr td:nth-child(1)", &foundSeasons, chromedp.NodeVisible),
	)

	if err != nil {
		log.Error("Can't get list of seasons")
	}

	seasons := []string{}
	for _, s := range foundSeasons {
		seasons = append(seasons, strings.Split(s.Children[0].NodeValue, "/")[1])
	}

	dl.downloadSeasons(ctx, seasons)
}

// GetSeasons downloads the seasons CSV data from C2 site
// into the Downloader struct's path
func (dl Downloader) GetSeasons(seasons []string) {
	ctx, cancel := (&dl).login()
	defer cancel()
	navigateToHistory(ctx)
	dl.downloadSeasons(ctx, seasons)
}

func (dl *Downloader) login() (context.Context, context.CancelFunc) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", dl.headless),
	)
	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, cancel := chromedp.NewContext(allocCtx)
	url := "https://log.concept2.com/login"
	err := chromedp.Run(ctx, chromedp.Navigate(url))
	if err != nil {
		log.Error("Could not open browser")
	}
	err = chromedp.Run(ctx,
		chromedp.Tasks{
			chromedp.SendKeys(`#username`, dl.Username),
			chromedp.SendKeys(`#password`, dl.Password),
			chromedp.Click(".form-signin .btn-primary", chromedp.NodeVisible),
			chromedp.WaitReady("#last_30_days_time"),
		},
	)

	if err != nil {
		log.Error("Could not log in to Concept2")
	}
	return ctx, cancel
}

func navigateToHistory(ctx context.Context) {
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://log.concept2.com/history"),
		chromedp.WaitReady(".history"),
	)
	if err != nil {
		log.Error("Could not find history")
	}
}

func (dl Downloader) downloadSeasons(ctx context.Context, seasons []string) {
	files := []string{}
	for _, s := range seasons {
		filename := fmt.Sprintf("concept2-season-%v.csv", s)
		os.Remove(filepath.Join(dl.Path, filename))
		chromedp.Run(ctx,
			page.SetDownloadBehavior(page.SetDownloadBehaviorBehaviorAllow).WithDownloadPath(dl.Path),
			chromedp.Navigate(fmt.Sprintf("https://log.concept2.com/season/%v/export", s)),
		)
		files = append(files, filename)
	}

	err := waitForFilesToExist(dl.Path, files, time.Second*3)
	if err != nil {
		log.Error(err)
	}
}

func waitForFilesToExist(dir string, files []string, timeout time.Duration) error {
	retryTime := 100 * time.Millisecond
	retryCount := 0
	for {
		exists := 0
		for _, f := range files {
			_, err := os.Stat(filepath.Join(dir, f))
			if err == nil {
				exists++
			}
		}
		if exists == len(files) {
			return nil
		}
		retryCount++
		if retryCount*int(retryTime) > int(timeout) {
			return errors.New("Files don't seem to have downloaded within timeout")
		}
		time.Sleep(retryTime)
	}
}
