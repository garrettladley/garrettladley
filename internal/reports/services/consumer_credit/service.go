package consumer_credit

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/garrettladley/garrettladley/internal/reports/constants"
	"github.com/garrettladley/garrettladley/internal/reports/types"
	"github.com/gofiber/fiber/v2"
)

func Query(ctx context.Context, start time.Time, end time.Time) ([]Data, error) {
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
			return []Data{}, ctx.Err()
		case err := <-errCh:
			return []Data{}, err
		case body = <-resultCh:
			done = true
			break
		}
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		return []Data{}, err
	}

	now := time.Now()
	data := make([]Data, (now.Year()-constants.CONSUMER_CREDIT_START)*4+types.QuarterFromTime(now).Int()-1)
	doc.Find("#content > div.data-table > table > tbody > tr").Each(func(i int, s *goquery.Selection) {
		data[i] = parse(s.Text())
	})

	return data, nil
}
