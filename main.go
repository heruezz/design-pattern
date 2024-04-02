package main

import (
	"design-pattern/factory"
	"fmt"
	"time"
)

// interface di go itu dibuat dimana tempat dia dipake, bukan tempat dimana dia di implementasi
type ContentCreator interface {
	Produce(time time.Time) factory.Content
	Type() factory.ContentType
}

func main() {
	var contentCreator ContentCreator
	contentCreator = &factory.PiewDiePie{}
	content := contentCreator.Produce(time.Now())
	content.Play()
	fmt.Println(contentCreator.Type())
}
