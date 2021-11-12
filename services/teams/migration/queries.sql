SELECT * FROM teams WHERE id IN (
  SELECT id FROM team_members WHERE user_id = $1
);

SELECT * FROM teams;     

DELETE FROM teams WHERE id = $1;

INSERT INTO teams (name, description) VALUES ($1, $2) RETURNING *;

INSERT INTO  team_members (user_id, team_id) VALUES ($1, $2) RETURNING *;