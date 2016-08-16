package status

import (
	"fmt"
	"strconv"

	"github.com/mholt/caddy"
	"github.com/mholt/caddy/caddyhttp/httpserver"
)

// init registers Status plugin
func init() {
	caddy.RegisterPlugin("status", caddy.Plugin{
		ServerType: "http",
		Action:     setup,
	})
}

// setup configures new Status middleware instance.
func setup(c *caddy.Controller) error {
	rules, err := statusParse(c)
	if err != nil {
		return err
	}

	cfg := httpserver.GetConfig(c)
	mid := func(next httpserver.Handler) httpserver.Handler {
		return Status{Rules: rules, Next: next}
	}
	cfg.AddMiddleware(mid)

	return nil
}

// statusParse parses status directive
func statusParse(c *caddy.Controller) (map[string]int, error) {
	res := make(map[string]int)
	hadBlock := false

	saveStatusRule := func(path, statusCodeString string) error {
		statusCode, err := strconv.Atoi(statusCodeString)
		if err != nil {
			return c.Err(fmt.Sprintf("Expecting a numeric status code, got '%s'", statusCodeString))
		}

		if _, exists := res[path]; exists {
			return c.Errf("Duplicate path: '%s'", path)
		}
		res[path] = statusCode

		return nil
	}

	for c.Next() {
		hadBlock = false

		for c.NextBlock() {
			hadBlock = true

			path := c.Val()

			if !c.NextArg() {
				return res, c.ArgErr()
			}

			if err := saveStatusRule(path, c.Val()); err != nil {
				return res, err
			}

			if c.NextArg() {
				return res, c.ArgErr()
			}
		}

		if !hadBlock {
			args := c.RemainingArgs()
			if len(args) != 2 {
				return res, c.ArgErr()
			}

			if err := saveStatusRule(args[0], args[1]); err != nil {
				return res, err
			}
		}
	}

	return res, nil
}
