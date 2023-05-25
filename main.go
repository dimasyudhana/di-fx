package main

import (
	"log"

	"go.uber.org/fx"
)

func main() {
	// t := Title("testing")
	// p := NewPublisher(&t)
	// m := NewMainService(p)
	// m.Run()

	fx.New(
		fx.Provide(NewMainService),
		fx.Provide(
			fx.Annotate(
				// Annotate our constructor
				NewPublisher,
				fx.As(new(IPublisher)),
				fx.ParamTags(`group:"titles"`),
			),
		),
		fx.Provide(
			titleComponent("testing1"),
		),
		fx.Provide(
			titleComponent("testing1.5"),
		),
		fx.Provide(
			titleComponent("testing2"),
		),
		fx.Invoke(func(service *MainService) {
			service.Run()
		}),
	).Run()
}

func titleComponent(title string) any {
	return fx.Annotate(
		func() *Title {
			t := Title(title)
			return &t
		},
		fx.ResultTags(`group:"titles"`),
	)
}

// Main service
type MainService struct {
	publisher IPublisher
}

func NewMainService(publisher IPublisher) *MainService {
	return &MainService{publisher: publisher}
}

func (service *MainService) Run() {
	service.publisher.Publish()
	log.Print("main program")
}

// Dependency
type IPublisher interface {
	Publish()
}

type Publisher struct {
	titles []*Title
}

// We will takes slice of title, using variadic func
func NewPublisher(titles ...*Title) *Publisher {
	return &Publisher{titles: titles}
}

func (publisher *Publisher) Publish() {
	for _, title := range publisher.titles {
		log.Print("publisher: ", *title)
	}
}

// Dependency of publisher
type Title string
