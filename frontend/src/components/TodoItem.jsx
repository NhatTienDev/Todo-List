export default function TodoItem({ todo, onToggle, onEdit, onDelete }) {
    return (
        <div className={`todo-item ${todo.is_completed ? "completed" : ""}`}>
            <div className="todo-content">
                <label className="todo-checkbox">
                    <input
                        type="checkbox"
                        checked={todo.is_completed}
                        onChange={() => onToggle(todo.id, !todo.is_completed)}
                    />
                    <span className="checkmark" />
                </label>
                <div className="todo-info">
                    <span className={`status-badge ${todo.is_completed ? "completed" : "pending"}`}>
                        <span className="status-dot" />
                        {todo.is_completed ? "Completed" : "Pending"}
                    </span>
                    <span className="todo-title">{todo.title}</span>
                    {todo.description && <p className="todo-desc">{todo.description}</p>}
                </div>
            </div>
            <div className="todo-actions">
                {!todo.is_completed && (
                    <button className="btn-edit" onClick={() => onEdit(todo)} title="Edit">
                        <svg
                            width="16"
                            height="16"
                            viewBox="0 0 24 24"
                            fill="none"
                            stroke="currentColor"
                            strokeWidth="2"
                        >
                            <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7" />
                            <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z" />
                        </svg>
                    </button>
                )}
                <button
                    className="btn-delete"
                    onClick={() => onDelete(todo.id, todo.title)}
                    title="Delete"
                >
                    <svg
                        width="16"
                        height="16"
                        viewBox="0 0 24 24"
                        fill="none"
                        stroke="currentColor"
                        strokeWidth="2"
                    >
                        <polyline points="3 6 5 6 21 6" />
                        <path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2" />
                    </svg>
                </button>
            </div>
        </div>
    );
}