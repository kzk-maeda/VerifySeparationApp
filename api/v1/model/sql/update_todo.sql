UPDATE todo_table
SET
 title = $2,
 body = $3,
 deadline = $4,
 created_at = $5
where
 id = $1
returning
 id, title, body, deadline