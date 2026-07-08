import { useState, useEffect, useRef } from "react";
import "./App.css";
import { getTodos, createTodo, updateTodo, deleteTodo } from "./api";
import TodoForm from "./components/TodoForm";
import TodoList from "./components/TodoList";
import SearchBar from "./components/SearchBar";
import ConfirmModal from "./components/ConfirmModal";
import Pagination from "./components/Pagination";
import Toast from "./components/Toast";

function App() {
    const [todos, setTodos] = useState([]);
    const [search, setSearch] = useState("");
    const [filterStatus, setFilterStatus] = useState("");
    const [editingTodo, setEditingTodo] = useState(null);
    const [showForm, setShowForm] = useState(false);
    const [loaded, setLoaded] = useState(false);
    const [deleteTarget, setDeleteTarget] = useState(null);
    const [error, setError] = useState("");
    const [toast, setToast] = useState(null);
    const [page, setPage] = useState(1);
    const [totalPages, setTotalPages] = useState(1);
    const fetchIdRef = useRef(0);

    useEffect(() => {
        const id = ++fetchIdRef.current;

        getTodos(search, filterStatus, page)
            .then((data) => {
                if (id !== fetchIdRef.current) return;
                setTodos(data.data || []);
                setTotalPages(Math.ceil(data.total / data.limit) || 1);
                setLoaded(true);
                setError("");
            })
            .catch((err) => {
                if (id !== fetchIdRef.current) return;
                setTodos([]);
                setError(err.message);
                setLoaded(true);
            });
    }, [search, filterStatus, page]);

    function showToast(message, type = 'success') {
        setToast({ message, type })
    }

    async function handleCreate({ title, description }) {
        const todo = await createTodo(title, description);
        setTodos((prev) => [todo, ...prev]);
        setShowForm(false);
        showToast('Task created successfully');
    }

    async function handleUpdate({ title, description }) {
        const payload = {};
        if (title !== editingTodo.title) payload.title = title;
        if (description !== editingTodo.description)
            payload.description = description;
        const updated = await updateTodo(editingTodo.id, payload);
        setTodos((prev) => prev.map((t) => (t.id === updated.id ? updated : t)));
        setEditingTodo(null);
        showToast('Task updated successfully');
    }

    async function handleToggle(id, isCompleted) {
        const updated = await updateTodo(id, { is_completed: isCompleted });
        setTodos((prev) => prev.map((t) => (t.id === updated.id ? updated : t)));
    }

    function handleDeleteClick(id, title) {
        setDeleteTarget({ id, title })
    }

    async function confirmDelete() {
        if (!deleteTarget) return
        await deleteTodo(deleteTarget.id)
        setTodos((prev) => prev.filter((t) => t.id !== deleteTarget.id))
        setDeleteTarget(null)
        showToast('Task deleted successfully');
    }

    function handleSearchChange(value) {
        setSearch(value)
        setPage(1)
    }

    function handleFilterChange(value) {
        setFilterStatus(value)
        setPage(1)
    }

    return (
        <div className="app">
            <header className="app-header">
                <h1>Todo List</h1>
            </header>

            <SearchBar
                search={search}
                onSearchChange={handleSearchChange}
                filterStatus={filterStatus}
                onFilterChange={handleFilterChange}
            />

            <div className="app-actions">
                {!showForm && !editingTodo && (
                    <button className="btn-add" onClick={() => setShowForm(true)}>
                        + New task
                    </button>
                )}
            </div>

            {showForm && (
                <TodoForm onSubmit={handleCreate} onCancel={() => setShowForm(false)} />
            )}

            {editingTodo && (
                <TodoForm
                    onSubmit={handleUpdate}
                    initialData={editingTodo}
                    onCancel={() => setEditingTodo(null)}
                />
            )}

            {error && <p className="error-msg">{error}</p>}

            {!loaded ? (
                <p className="loading-msg">Loading...</p>
            ) : (
                <>
                    <TodoList
                        todos={todos}
                        onToggle={handleToggle}
                        onEdit={setEditingTodo}
                        onDelete={handleDeleteClick}
                    />
                    <Pagination
                        page={page}
                        totalPages={totalPages}
                        onPageChange={setPage}
                    />
                </>
            )}
            {deleteTarget && (
                <ConfirmModal
                    title={`Delete task "${deleteTarget.title}"?`}
                    onConfirm={confirmDelete}
                    onCancel={() => setDeleteTarget(null)}
                />
            )}
            {toast && (
                <Toast message={toast.message} type={toast.type} onClose={() => setToast(null)} />
            )}
        </div>
    );
}

export default App;
