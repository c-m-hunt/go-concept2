package downloader

import (
	"context"
	"fmt"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	log "github.com/sirupsen/logrus"
)

type Downloader struct {
	Username string
	Password string
	Path     string
}

func (dl Downloader) GetSeasons(seasons []string) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
	)
	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()
	url := "https://log.concept2.com/login"
	err := chromedp.Run(ctx, chromedp.Navigate(url))
	if err != nil {
		log.Error("Could not open browser")
	}

	fmt.Print(dl)

	err = chromedp.Run(ctx,
		chromedp.Tasks{
			chromedp.WaitVisible(".form-signin .btn-primary"),
			chromedp.SendKeys(`#username`, dl.Username),
			chromedp.SendKeys(`#password`, dl.Password),
			chromedp.Click(".form-signin .btn-primary", chromedp.NodeVisible),
			chromedp.WaitReady("#last_30_days_time"),
			chromedp.Navigate("https://log.concept2.com/history"),
			chromedp.WaitReady(".history"),
		},
	)

	if err != nil {
		log.Error("Could not log in to Concept2")
	}

	for _, s := range seasons {
		err = chromedp.Run(ctx,
			page.SetDownloadBehavior(page.SetDownloadBehaviorBehaviorAllow).WithDownloadPath(dl.Path),
			chromedp.Navigate(fmt.Sprintf("https://log.concept2.com/season/%v/export", s)),
		)
	}

	time.Sleep(1 * time.Second)
}
