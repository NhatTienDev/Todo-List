-- name: GetTodos :many
SELECT id, title, description, is_completed, created_at, updated_at 
FROM todos
WHERE 
    (sqlc.arg(search)::text = '' OR title ILIKE '%' || sqlc.arg(search) || '%')
    AND (
        sqlc.arg(filter_status)::text = 'all' 
        OR (sqlc.arg(filter_status)::text = 'completed' AND is_completed = TRUE)
        OR (sqlc.arg(filter_status)::text = 'pending' AND is_completed = FALSE)
    )
ORDER BY created_at DESC
LIMIT sqlc.arg(page_limit)::int OFFSET sqlc.arg(page_offset)::int;

-- name: CountTodos :one
SELECT COUNT(*) FROM todos
WHERE 
    (sqlc.arg(search)::text = '' OR title ILIKE '%' || sqlc.arg(search) || '%')
    AND (
        sqlc.arg(filter_status)::text = 'all' 
        OR (sqlc.arg(filter_status)::text = 'completed' AND is_completed = TRUE)
        OR (sqlc.arg(filter_status)::text = 'pending' AND is_completed = FALSE)
    );

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

-- name: DeleteTodo :execrows
DELETE FROM todos WHERE id = $1;
