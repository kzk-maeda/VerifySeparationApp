SELECT
 id, uuid, user_id
FROM session_table 
WHERE user_id = $1