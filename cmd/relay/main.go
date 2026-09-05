package main

import "log"

// Phase C0 placeholder. The real outbox-relay logic (SKIP LOCKED + publish
// to Kafka) lands in Phase C4/C5, not before.
func main() {
	log.Println("outbox-relay: placeholder — no outbox logic yet (Phase C4)")
}
