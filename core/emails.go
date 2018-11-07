package titian

import ()

type Emails struct {
	app        *Container
	unverified map[EmailAddress]struct{}
}
