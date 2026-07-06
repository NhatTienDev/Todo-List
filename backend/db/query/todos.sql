-- name: GetTodos :many
SELECT id, title, description, is_completed, created_at, updated_at 
FROM todos
WHERE 
    -- Search: If an empty string is passed, skip it; if it contains text, search using ILIKE
    (sqlc.arg(search)::text = '' OR title ILIKE '%' || sqlc.arg(search) || '%')
    AND (
        -- Filter by status: 'all', 'completed', 'pending'
        sqlc.arg(filter_status)::text = 'all' 
        OR (sqlc.arg(filter_status)::text = 'completed' AND is_completed = TRUE)
        OR (sqlc.arg(filter_status)::text = 'pending' AND is_completed = FALSE)
    )
ORDER BY created_at DESC;

-- name: GetTodoByID :one
SELECT * FROM todos WHERE id = $1 LIMIT 1;

-- name: CreateTodo :one
INSERT INTO todos (title, description)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateTodo :one
UPDATE todos
SET 
    title = $2, 
    description = $3, 
    is_completed = $4, 
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM todos WHERE id = $1;