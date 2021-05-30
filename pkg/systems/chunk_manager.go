package systems

// Handles the asynchronous loading of chunks while safely unloading chunks on the main thread

type ChunkManager struct {
	running bool
}

func NewChunkManager() *ChunkManager {
	return &ChunkManager{}
}

func (c *ChunkManager) Poll() {
	// panic("implement me")
}

func (c *ChunkManager) Step(dt float32) bool {
	// panic("implement me")

	return c.running
}

func (c *ChunkManager) Stop() {
	c.running = false
}
