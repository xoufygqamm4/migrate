package migrate

import (
	"context"
	"time"
)

// ... existing code ...

func (m *Migrate) markDirty(version int) error {
	// Use a background context with a timeout to ensure the dirty flag is written
	// even if the original migration context has been cancelled or timed out.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	return m.store.SetVersion(ctx, version, true)
}

// Inside the migration execution loop:
// if err != nil {
//     if err == context.DeadlineExceeded || err == context.Canceled {
//         _ = m.markDirty(version)
//     }
//     return err
// }