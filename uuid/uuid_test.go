package uuid_test

import (
	"fmt"
	"testing"

	"github.com/ironzhang/pearls/uuid"
)

func TestUUID(t *testing.T) {
	for i := 0; i < 10; i++ {
		u := uuid.New()
		fmt.Println(u)
	}
}
