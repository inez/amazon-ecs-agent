package main

import (
	"fmt"
)

type DockerTaskEngine struct {
	indexToArn map[int]string
}

func NewDockerTaskEngine() *DockerTaskEngine {
	dockerTaskEngine := &DockerTaskEngine{
		indexToArn: make(map[int]string),
	}
	return dockerTaskEngine
}

func (engine *DockerTaskEngine) addByArn(arn string) int {
	index := 0
	for {
		_, ok := engine.indexToArn[index]
		if !ok {
			engine.indexToArn[index] = arn
			return index
		}
		index++
	}
}

func (engine *DockerTaskEngine) removeByArn(arn string) {
	for key, value := range engine.indexToArn {
		if value == arn {
			delete(engine.indexToArn, key)
		}
	}
}

func main() {
	fmt.Println("Hello, playground")
	taskEngine := NewDockerTaskEngine()
	fmt.Println("indexToArn:", taskEngine.indexToArn)
	// taskEngine.addByArn("t0")
	var index int = taskEngine.addByArn("t0")
	fmt.Sprintf("%v,%v,%v", index*3, index*3+1, index*3+2)
	fmt.Println("indexToArn:", taskEngine.indexToArn)
}
