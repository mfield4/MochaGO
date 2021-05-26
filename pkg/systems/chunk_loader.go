package systems

type ChunkLoader struct {
}

func NewChunkLoader() *ChunkLoader {
	return &ChunkLoader{}
}

func (c *ChunkLoader) Start() {
	panic("implement me")
}

func (c *ChunkLoader) Step(dt float32) {
	panic("implement me")
}

func (c *ChunkLoader) Stop() {
	panic("implement me")
}
