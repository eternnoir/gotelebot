package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplyKeyboardMarkup(t *testing.T) {
	assert := assert.New(t)
	row1 := []string{"1", "2", "3"}
	row2 := []string{"4", "5", "6"}
	rkm := ReplyKeyboardMarkup{}
	rkm.Keyboard = append(rkm.Keyboard, row1)
	rkm.Keyboard = append(rkm.Keyboard, row2)
	_, err := rkm.ToJson()
	if err != nil {
		assert.Fail(err.Error())
	}
}
