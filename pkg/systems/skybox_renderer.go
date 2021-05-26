package systems

type SkyboxRenderer struct {
}

func NewSkyboxRenderer() *SkyboxRenderer {
	return &SkyboxRenderer{}
}

func (s *SkyboxRenderer) Render(dt float32) {
	panic("implement me")
}
