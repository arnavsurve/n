package n

import (
	"fmt"
	"time"
)

func DailyNote() string {
	now := time.Now()
	return fmt.Sprintf(`# %s

## Notes
-

## Todos
- [ ]

## Log
-
`, now.Format("Monday, January 02, 2006"))
}

func ProjectNote(slug string) string {
	now := time.Now()
	return fmt.Sprintf(`# %s
_Created: %s_

## Context

## Notes
-

## Next
- [ ]
`, slug, now.Format("2006-01-02 15:04"))
}
