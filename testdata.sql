INSERT INTO Users (username, password_hash, priviledges) VALUES ('exists', 'hash', 'default');
INSERT INTO Users (username, password_hash, priviledges) VALUES ('scrum', 'hash2', 'admin');
INSERT INTO Users (username, password_hash, priviledges) VALUES ('valuable', 'hash3', 'value');
INSERT INTO Users (username, password_hash, priviledges) VALUES ('desirable', 'hash4', 'desire');

INSERT INTO Groups (owner_id, name, description) VALUES (2, 'Group', 'This is awesome');
INSERT INTO Groups (owner_id, name, description) VALUES (2, 'Propably empty', 'This sould be empty');

INSERT INTO Words (owner_id, word, description) VALUES (2, 'This is word', 'Awesome text');
INSERT INTO Words (owner_id, word, description) VALUES (2, 'This is word2', 'Awesome text2'); /* id = 2 */
INSERT INTO Words (owner_id, word, description) VALUES (1, 'This is word3', 'Awesome text3');
INSERT INTO Words (owner_id, word, description) VALUES (4, 'This is word4', 'Awesome text4');

INSERT INTO WordItems (word_id, word, description) VALUES (2, 'Item1', 'item desc1');
INSERT INTO WordItems (word_id, word, description) VALUES (2, 'Item2', 'item desc2');
INSERT INTO WordItems (word_id, word, description) VALUES (2, 'Item3', 'item desc3');
INSERT INTO WordItems (word_id, word, description) VALUES (1, 'Item4', 'item desc4');
INSERT INTO WordItems (word_id, word, description) VALUES (1, 'Item5', 'item desc5');
INSERT INTO WordItems (word_id, word, description) VALUES (3, 'Item6', 'item desc6');
INSERT INTO WordItems (word_id, word, description) VALUES (3, 'Item7', 'item desc7');
INSERT INTO WordItems (word_id, word, description) VALUES (3, 'Item8', 'item desc8');
INSERT INTO WordItems (word_id, word, description) VALUES (3, 'Item9', 'item desc9');

INSERT INTO GroupLinks (group_id, word_id) VALUES (1, 2);
INSERT INTO GroupLinks (group_id, word_id) VALUES (1, 1);