package factory

import (
	"fmt"
	"time"
)

type Content interface {
	Play()
}

type CloudContent struct{}

func (c *CloudContent) Play() {
	fmt.Println("this is cloud content about kubernetes")
}

type GrebekRumahContent struct{}

func (c *GrebekRumahContent) Play() {
	fmt.Println("ini adalah content yang sangat diminati para remaja")
}

type CodingContent struct{}

func (c *CodingContent) Play() {
	fmt.Println("ini adalah content pembelajaran Coding")
}

type ContentType string

const (
	Hiburan ContentType = "hiburan"
	Edukasi ContentType = "edukasi"
)

// pindah kedepan
// type ContentCreator interface {
// 	Produce(time time.Time) Content
// 	Type() ContentType
// }

type Heru struct {
}

func (i *Heru) Produce(time time.Time) Content {
	if time.Weekday().String() == "Tuesday" {
		return &CloudContent{}
	} else if time.Weekday().String() == "Thursday" {
		return &CodingContent{}
	} else {
		return nil
	}
}

func (i Heru) Type() ContentType {
	return Edukasi
}

type PiewDiePie struct {
}

func (a *PiewDiePie) Produce(time time.Time) Content {
	return &GrebekRumahContent{}
}

func (a PiewDiePie) Type() ContentType {
	return Hiburan
}
