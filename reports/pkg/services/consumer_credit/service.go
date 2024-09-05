package consumer_credit

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/garrettladley/garrettladley/reports/pkg/constants"
	"github.com/gofiber/fiber/v2"
)

func Query(ctx context.Context, start time.Time, end time.Time) error {
	resultCh := make(chan []byte)
	errCh := make(chan error)

	go func() {
		defer close(resultCh)
		defer close(errCh)

		statusCode, body, errs := fiber.Get(constants.CONSUMER_CREDIT).Bytes()
		if errs != nil {
			errCh <- fmt.Errorf("failed to get consumer credit endpoint: %w ", errors.Join(errs...))
			return
		}

		if statusCode != http.StatusOK {
			errCh <- fmt.Errorf("received status code: %d", statusCode)
			return
		}

		resultCh <- body
	}()

	var body []byte
	var done bool
	for !done {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-errCh:
			return err
		case body = <-resultCh:
			done = true
			break
		}
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		return err
	}

	// 1943
	now := time.Now()
	data := make([]Data, (now.Year()-constants.CONSUMER_CREDIT_START)*4+whatQuarter(now).into())
	fmt.Println(data) // remove unused
	doc.Find("#content > div.data-table > table > tbody > tr").Each(func(i int, s *goquery.Selection) {
		text := s.Text() // text is a string
		fmt.Printf("Row %d: '%v'\n", i+1, text)

		// for i, r := range text {
		// 	fmt.Printf("Rune at %d: %c\n", i, r)
		// }

	})

	return nil
}
