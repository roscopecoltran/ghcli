package common

import (
	"fmt"
	"io"

	"github.com/palantir/pkg/cli"
	"github.com/palantir/pkg/cli/flag"
	"github.com/pkg/errors"
)

func RateLimit() cli.Command {
	return cli.Command{
		Name:  "rate-limit",
		Usage: "print the rate limit for the authenticated user",
		Flags: []flag.Flag{
			GitHubTokenFlag,
		},
		Action: func(ctx cli.Context) error {
			return doRateLimit(NewGitHubParams(ctx), ctx.App.Stdout)
		},
	}
}

func doRateLimit(params GitHubParams, stdout io.Writer) error {
	client := params.CachingOAuthGitHubClient()
	limits, _, err := client.RateLimits()
	if err != nil {
		return errors.Wrapf(err, "failed to get rate limits")
	}
	fmt.Fprintf(stdout, "Remaining requests: %d/%d\n", limits.Core.Remaining, limits.Core.Limit)
	fmt.Fprintf(stdout, "Rate limit resets:  %v\n", limits.Core.Reset.Format("15:04:05 MST Mon Jan 2 2006"))
	return nil
}
