/*
 * `cleanup` is a utility for deleting orphaned Rollbar projects from failed
 * acceptance test runs.
 */
package main

import (
	"github.com/rollbar/terraform-provider-rollbar/client"
	"github.com/rs/zerolog/log"
	"os"
	"strings"
)

func main() {
	log.Info().Msg("Cleaning up orphaned Rollbar projects from failed acceptance test runs.")

	token := os.Getenv("ROLLBAR_TOKEN")
	c := client.NewClient(token)

	projects, err := c.ListProjects()
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	for _, p := range projects {
		l := log.With().
			Str("name", p.Name).
			Int("id", p.Id).
			Logger()
		if strings.HasPrefix(p.Name, "tf-acc-test-") {
			err = c.DeleteProject(p.Id)
			if err != nil {
				l.Fatal().Err(err).Send()
			}
			l.Info().Msg("Deleted project")
		}
	}

	log.Info().Msg("Cleanup complete")
}
