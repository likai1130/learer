package typed

import "learner/example/mogodemo/basic"

type EventGetter interface {

}

type EventInterface interface {

}

type events struct {
	client basic.DataSource
}