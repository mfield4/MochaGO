package systems

// Meant to be rendering system responsible for rendering chunks

type ChunkRenderer struct {
}

func NewChunkRenderer() *ChunkRenderer {
	return &ChunkRenderer{}
}

func (c *ChunkRenderer) Render(dt float32) {
	panic("implement me")
}
