package utils

import (
	"errors"

	"github.com/tsa-dom/lang-trainer/app/models/words"
)

type Word struct {
	Base  words.Word
	Items []words.WordItem
}

func AddWordWithItems(word Word) error {
	base, err := words.CreateWord(word.Base)
	if err != nil {
		return errors.New("failed to add word")
	}
	err = words.AddItemsToWord(base.Id, word.Items)
	if err != nil {
		return errors.New("failed to add items to word")
	}
	return nil
}
