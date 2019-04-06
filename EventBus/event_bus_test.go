package EventBus

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	b := New()
	assert.NotNil(t, b)
}

func TestEventBus_Subscribe(t *testing.T) {
	b := New()
	c := b.Subscribe("topic", func() {
		fmt.Println("on event")
	})
	b.Fire("topic")

	b.Fire("topic")
	c()
	b.Fire("topic")
}
