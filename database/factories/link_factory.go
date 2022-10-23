package factories

import (
	"zerodot618/huango/app/models/link"

	"github.com/bxcodec/faker/v3"
)

func MakeLinks(count int) []link.Link {

	var objs []link.Link

	for i := 0; i < count; i++ {
		model := link.Link{
			Name: faker.Username(),
			URL:  faker.URL(),
		}
		objs = append(objs, model)
	}

	return objs
}
