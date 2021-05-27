package systems

// Handles the asynchronous loading of chunks while safely unloading chunks on the main thread

type ChunkManager struct {
}

func (c *ChunkManager) Poll() {
	panic("implement me")
}

func NewChunkManager() *ChunkManager {
	return &ChunkManager{}
}

func (c *ChunkManager) Step(dt float32) bool {
	panic("implement me")
}

func (c *ChunkManager) Stop() {
	panic("implement me")
}
