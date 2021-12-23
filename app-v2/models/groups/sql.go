package models

func linkWordToGroup() string {
	return `
		INSERT INTO GroupLinks (group_id, word_id)
		VALUES ($1, $2)
	`
}

func wordsInGroup() string {
	return `
		SELECT 
			W.id, W.owner_id, W.word, W.description, 
			I.id, I.word, I.description 
		FROM GroupLinks G 
		LEFT JOIN Words W ON G.group_id=$1 AND G.word_id=W.id 
		LEFT JOIN WordItems I ON W.id=I.word_id
	`
}
