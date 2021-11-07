package groups

func addGroup() string {
	return `
		INSERT INTO Groups (owner_id, name, description)
		VALUES ($1, $2, $3)
		RETURNING id
	`
}

func addWord() string {
	return `
		INSERT INTO Words (owner_id, word, description)
		VALUES ($1, $2, $3)
		RETURNING id
	`
}

func wordById() string {
	return `
		SELECT id, owner_id, word, description FROM Words WHERE id=$1;
	`
}

func wordItemsByWordId() string {
	return `
		SELECT I.id, I.word, I.description FROM Words W, WordItems I Where W.id=$1 AND W.id=I.word_id;
	`
}

func addWordItem() string {
	return `
		INSERT INTO WordItems (word_id, word, description)
		VALUES ($1, $2, $3)
	`
}

func deleteItemsByWordId() string {
	return `
		DELETE FROM WordItems WHERE word_id=$1
	`
}

func deleteWordById() string {
	return `
		DELETE FROM Words WHERE id=$1;
	`
}

func linkWordToGroup() string {
	return `
		INSERT INTO GroupLinks (group_id, word_id)
		VALUES ($1, $2)
		RETURNING id
	`
}

func deleteWordLink() string {
	return `
		DELETE FROM GroupLinks WHERE group_id=$1 AND word_id=$2
	`
}

func getGroups() string {
	return `
		SELECT id, owner_id, name, description 
		FROM Groups WHERE owner_id=$1
	`
}
