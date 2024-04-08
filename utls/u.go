package utls

import (
	http "github.com/bogdanfinn/fhttp"
	"github.com/bogdanfinn/tls-client/profiles"
	"github.com/gofiber/fiber/v2"
	"regexp"
	"strings"
)

func GetProfileUA(ua string) (profiles.ClientProfile, bool) {
	re := regexp.MustCompile(`(?i)(chrome|firefox|safari|opera)/(\d+)`)
	match := re.FindStringSubmatch(ua)
	if len(match) > 2 {
		key := strings.ToLower(match[1]) + "_" + match[2]
		profile, ok := profiles.MappedTLSClients[key]
		if ok {
			return profile, ok
		}
	}
	return profiles.Chrome_117, false
}

func GetBrowserFrom(c *fiber.Ctx) http.Header {
	ua := c.Request().Header.Peek("User-Agent")
	platform := c.Request().Header.Peek("Sec-ch-ua-platform")
	language := c.Request().Header.Peek("Accept-Language")
	accept := c.Request().Header.Peek("Accept")
	encoding := c.Request().Header.Peek("Accept-Encoding")
	sa := c.Request().Header.Peek("Sec-ch-ua")
	ct := c.Request().Header.Peek("Content-Type")
	ck := c.Request().Header.Peek("Cookie")
	clientIP := c.IP()
	return http.Header{
		"accept":             {string(accept)},
		"accept-language":    {string(language)},
		"user-agent":         {string(ua)},
		"sec-ch-ua":          {string(sa)},
		"Content-Type":       {string(ct)},
		"cookie":             {string(ck)},
		"X-Real-IP":          {clientIP},
		"X-Forwarded-For":    {clientIP},
		"sec-ch-ua-mobile":   {"?0"},
		"Sec-ch-ua-platform": {string(platform)},
		"accept-encoding":    {string(encoding)},
		http.HeaderOrderKey: {
			"accept",
			"accept-language",
			"content-type",
			"cookie",
			"user-agent",
			"accept-encoding",
			"sec-ch-ua",
			"sec-ch-ua-mobile",
			"Sec-ch-ua-platform",
			"X-Real-IP",
			"X-Forwarded-For",
		},
	}
}
