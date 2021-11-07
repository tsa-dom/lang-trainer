package utils

import (
	"errors"

	"github.com/tsa-dom/lang-trainer/app/models/groups"
)

type Word struct {
	Base  groups.Word
	Items []groups.WordItem
}

func AddWordWithItems(word Word) error {
	base, err := groups.CreateWord(word.Base)
	if err != nil {
		return errors.New("failed to add word")
	}
	err = groups.AddItemsToWord(base.Id, word.Items)
	if err != nil {
		return errors.New("failed to add items to word")
	}
	return nil
}
