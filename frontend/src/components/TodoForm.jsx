import { useState } from "react";

export default function TodoForm({ onSubmit, initialData, onCancel }) {
    const [title, setTitle] = useState(initialData?.title || "");
    const [description, setDescription] = useState(
        initialData?.description || "",
    );
    const [error, setError] = useState("");

    const isEditing = !!initialData;

    async function handleSubmit(e) {
        e.preventDefault();
        setError("");
        try {
            await onSubmit({ title: title.trim(), description: description.trim() });
            if (!isEditing) {
                setTitle("");
                setDescription("");
            }
        } catch (err) {
            setError(err.message);
        }
    }

    return (
        <form className="todo-form" onSubmit={handleSubmit}>
            <h2>{isEditing ? "Edit task" : "Add new task"}</h2>
            {error && <p className="form-error">{error}</p>}
            <input
                type="text"
                placeholder="Task title"
                value={title}
                onChange={(e) => setTitle(e.target.value)}
                required
            />
            <textarea
                placeholder="Description (optional)"
                value={description}
                onChange={(e) => setDescription(e.target.value)}
                rows={3}
            />
            <div className="form-actions">
                <button type="submit">{isEditing ? "Update" : "Add"}</button>
                <button type="button" className="btn-cancel" onClick={onCancel}>
                    Cancel
                </button>
            </div>
        </form>
    );
}