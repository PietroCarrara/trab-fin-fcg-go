package graphics

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gl/gl/v3.3-core/gl"
)

func loadShaders() (uint32, error) {
	vertex_shader_id, err := loadShaderVertex("data/shaders/shader.vert.glsl")
	if err != nil {
		return 0, err
	}

	fragment_shader_id, err := loadShaderFragment("data/shaders/shader.frag.glsl")
	if err != nil {
		return 0, err
	}

	return createGpuProgram(vertex_shader_id, fragment_shader_id), nil
}

func loadShaderVertex(filename string) (uint32, error) {
	shaderID := gl.CreateShader(gl.VERTEX_SHADER)
	return shaderID, loadShader(filename, shaderID)
}

func loadShaderFragment(filename string) (uint32, error) {
	shaderID := gl.CreateShader(gl.FRAGMENT_SHADER)
	return shaderID, loadShader(filename, shaderID)
}

func loadShader(filename string, shaderID uint32) error {
	bts, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	str, free := gl.Strs(string(bts) + "\x00")
	defer free()

	gl.ShaderSource(shaderID, 1, str, nil)
	gl.CompileShader(shaderID)

	var compiledOK int32
	var logLen int32
	gl.GetShaderiv(shaderID, gl.COMPILE_STATUS, &compiledOK)
	gl.GetShaderiv(shaderID, gl.INFO_LOG_LENGTH, &logLen)
	if logLen != 0 {
		logContents := make([]uint8, logLen)
		gl.GetShaderInfoLog(shaderID, logLen, &logLen, &logContents[0])
		if compiledOK == 0 {
			return fmt.Errorf("error while compiling \"%s\":\n%s", filename, string(logContents))
		} else {
			log.Printf("compilation of \"%s\" did not go smoothly:\n%s", filename, string(logContents))
		}
	}

	return nil
}

func createGpuProgram(vertex_shader_id, fragment_shader_id uint32) uint32 {
	program_id := gl.CreateProgram()

	gl.AttachShader(program_id, vertex_shader_id)
	gl.AttachShader(program_id, fragment_shader_id)
	gl.LinkProgram(program_id)

	var linkedOK int32
	gl.GetProgramiv(program_id, gl.LINK_STATUS, &linkedOK)
	if linkedOK == 0 {
		var logLen int32
		gl.GetShaderiv(program_id, gl.INFO_LOG_LENGTH, &logLen)
		logContents := make([]uint8, logLen)
		gl.GetShaderInfoLog(program_id, logLen, &logLen, &logContents[0])
		log.Printf("linking of shaders did no go smoothly:\n%s", string(logContents))
	}

	return program_id
}
