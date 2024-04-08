package serve

import (
	"bytes"
	"fkclaude/utls"
	"fmt"
	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/gofiber/fiber/v2"
	"io"
	"log"
	"strings"
)

func forwardRequest(c *fiber.Ctx, url string) error {
	method := string(c.Request().Header.Method())
	ua := c.Request().Header.Peek("User-Agent")
	body := c.Request().Body()
	jar := tls_client.NewCookieJar()
	profile, _ := utls.GetProfileUA(string(ua))
	options := []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(120),
		tls_client.WithClientProfile(profile),
		tls_client.WithNotFollowRedirects(),
		tls_client.WithCookieJar(jar),
	}
	client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	if err != nil {
		log.Printf("Failed to create HTTP client: %v. Retrying...\n", err)
		return err
	}
	contentType := string(c.Request().Header.ContentType())
	var req *http.Request
	if strings.HasPrefix(contentType, "application/octet-stream") ||
		strings.HasPrefix(contentType, "video/") ||
		strings.HasPrefix(contentType, "audio/") ||
		strings.HasPrefix(contentType, "text/event-stream") {
		req, err = http.NewRequest(method, url, bytes.NewReader(body))
	} else {
		req, err = http.NewRequest(method, url, io.NopCloser(bytes.NewReader(body)))
	}
	if err != nil {
		fmt.Printf("Failed to create request: %v\n\n", err)
		return err
	}
	req.Header = utls.GetBrowserFrom(c)
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Failed to send request: %v\n", err)
		return err
	}
	defer res.Body.Close()

	for _, cookie := range res.Cookies() {
		existingCookie := c.Cookies(cookie.Name)
		if len(existingCookie) == 0 || cookie.Name != "__cf_bm" {
			c.Cookie(&fiber.Cookie{
				Name:     cookie.Name,
				Value:    cookie.Value,
				Expires:  cookie.Expires,
				Domain:   cookie.Domain,
				Path:     cookie.Path,
				Secure:   cookie.Secure,
				HTTPOnly: cookie.HttpOnly,
			})
		}
	}

	responseContentType := res.Header.Get("Content-Type")
	if strings.HasPrefix(responseContentType, "application/octet-stream") ||
		strings.HasPrefix(responseContentType, "video/") ||
		strings.HasPrefix(responseContentType, "audio/") ||
		strings.HasPrefix(responseContentType, "text/event-stream") {
		c.Set("Content-Type", responseContentType)
		log.Print("Content-Type: ", responseContentType)
		c.Status(res.StatusCode)
		io.Copy(c, res.Body)
		res.Body.Close()
		return nil

	} else {
		c.Status(res.StatusCode)
		io.Copy(c, res.Body)
		res.Body.Close()
		return nil
	}
}

func APIHandler(app *fiber.App) {
	app.All("/*", func(c *fiber.Ctx) error {
		path := c.Path()
		url := "https://claude.ai" + path
		return forwardRequest(c, url)
	})
}
