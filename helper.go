package main

import (
	"fmt"
	"sync"
)

type IndexToArnSafe struct {
	mu sync.Mutex
	v  map[int]string
}

var indexToArnSafe = IndexToArnSafe{v: make(map[int]string)}

type DockerTaskEngine struct {
}

func NewDockerTaskEngine() *DockerTaskEngine {
	dockerTaskEngine := &DockerTaskEngine{}
	return dockerTaskEngine
}

func (engine *DockerTaskEngine) addByArn(arn string) int {
	indexToArnSafe.mu.Lock()
	defer indexToArnSafe.mu.Unlock()
	index := 0
	for {
		_, ok := indexToArnSafe.v[index]
		if !ok {
			indexToArnSafe.v[index] = arn
			return index
		}
		index++
	}
}

func (engine *DockerTaskEngine) removeByArn(arn string) {
	indexToArnSafe.mu.Lock()
	defer indexToArnSafe.mu.Unlock()
	for key, value := range indexToArnSafe.v {
		if value == arn {
			delete(indexToArnSafe.v, key)
		}
	}
}

func main() {
	fmt.Println("Hello, playground")
	taskEngine := NewDockerTaskEngine()
	fmt.Println("global_MAP:", indexToArnSafe.v)
	taskEngine.addByArn("t0")
	var index int = taskEngine.addByArn("t0")
	fmt.Sprintf("%v,%v,%v", index*3, index*3+1, index*3+2)
	fmt.Println("global_MAP:", indexToArnSafe.v)
}
