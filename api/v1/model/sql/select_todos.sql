SELECT
 id, uuid, title, body, deadline, user_id, created_at
FROM
 todo_table
WHERE
 user_id = $1
ORDER BY created_at DESC 