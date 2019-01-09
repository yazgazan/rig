package validators

import (
	"fmt"
	"net/url"
)

type URL func(*url.URL) error

func URLScheme(scheme string) URL {
	return func(u *url.URL) error {
		if u.Scheme != scheme {
			return fmt.Errorf("url should use %q scheme", scheme)
		}

		return nil
	}
}

func URLExcludeScheme(scheme string) URL {
	return func(u *url.URL) error {
		if u.Scheme == scheme {
			return fmt.Errorf("url should not use %q scheme", scheme)
		}

		return nil
	}
}
