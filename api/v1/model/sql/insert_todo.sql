INSERT INTO
 todo_table ( uuid, title, body, deadline, user_id, created_at )
values
 ($1, $2, $3, $4, $5, $6)
returning
 id, uuid, title, body, deadline, user_id, created_at