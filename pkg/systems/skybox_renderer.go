package systems

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"go_entity_component_services/pkg/components"
	"go_entity_component_services/pkg/rendering"
	"log"
)

type SkyboxRenderer struct {
	skyboxID             uint32
	skyboxVAO, skyboxVBO uint32

	skyboxShader rendering.Shader
}

func NewSkyboxRenderer() *SkyboxRenderer {
	v := components.CubeVertices()
	vao, vbo := rendering.InitVBOPos((v)[:])

	return &SkyboxRenderer{
		skyboxID: func() uint32 {
			mapID, err := rendering.NewCubeMapD([6]string{
				"resources/textures/skybox/right.jpg",
				"resources/textures/skybox/left.jpg",
				"resources/textures/skybox/top.jpg",
				"resources/textures/skybox/bottom.jpg",
				"resources/textures/skybox/front.jpg",
				"resources/textures/skybox/back.jpg"})

			if err != nil {
				log.Fatalf("Failed to create skybox %v\n", err)
				return 0
			}

			return mapID
		}(),
		skyboxVAO: vao,
		skyboxVBO: vbo,
		skyboxShader: func() rendering.Shader {
			shader, err := rendering.NewShader("shaders/skybox.vert", "shaders/skybox.frag", "")
			if err != nil {
				log.Fatalf("Failed to create skybox %v\n", err)
				return rendering.Shader{}
			}

			uniIdx := gl.GetUniformBlockIndex(shader.Program, gl.Str("Matrices\x00"))
			gl.UniformBlockBinding(shader.Program, uniIdx, 0)

			shader.Use()
			shader.SetInt32("skybox", 0)

			return shader
		}(),
	}
}

func (s *SkyboxRenderer) Render(dt float32) {
	// RENDER MY CUBE
	s.skyboxShader.Use()

	gl.DepthFunc(gl.LEQUAL)
	gl.BindVertexArray(s.skyboxVAO)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_CUBE_MAP, s.skyboxID)
	gl.DrawArrays(gl.TRIANGLES, 0, 36)
	gl.BindVertexArray(0)
	gl.DepthFunc(gl.LESS)
}
