import { useState, useEffect, useRef } from 'react'
import './App.css'
import { getTodos, createTodo, updateTodo, deleteTodo } from './api'
import TodoForm from './components/TodoForm'
import TodoList from './components/TodoList'
import SearchBar from './components/SearchBar'

function App() {
  const [todos, setTodos] = useState([])
  const [search, setSearch] = useState('')
  const [filterStatus, setFilterStatus] = useState('')
  const [editingTodo, setEditingTodo] = useState(null)
  const [showForm, setShowForm] = useState(false)
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')
  const fetchIdRef = useRef(0)

  useEffect(() => {
    const id = ++fetchIdRef.current
    setLoading(true)

    getTodos(search, filterStatus)
      .then((data) => {
        if (id !== fetchIdRef.current) return
        setTodos(data || [])
        setError('')
        setLoading(false)
      })
      .catch((err) => {
        if (id !== fetchIdRef.current) return
        setError(err.message)
        setLoading(false)
      })
  }, [search, filterStatus])

  async function handleCreate({ title, description }) {
    const todo = await createTodo(title, description)
    setTodos((prev) => [todo, ...prev])
    setShowForm(false)
  }

  async function handleUpdate({ title, description }) {
    const payload = {}
    if (title !== editingTodo.title) payload.title = title
    if (description !== editingTodo.description) payload.description = description
    const updated = await updateTodo(editingTodo.id, payload)
    setTodos((prev) => prev.map((t) => (t.id === updated.id ? updated : t)))
    setEditingTodo(null)
  }

  async function handleToggle(id, isCompleted) {
    const updated = await updateTodo(id, { is_completed: isCompleted })
    setTodos((prev) => prev.map((t) => (t.id === updated.id ? updated : t)))
  }

  async function handleDelete(id) {
    await deleteTodo(id)
    setTodos((prev) => prev.filter((t) => t.id !== id))
  }

  return (
    <div className="app">
      <header className="app-header">
        <h1>Todo List</h1>
      </header>

      <SearchBar
        search={search}
        onSearchChange={setSearch}
        filterStatus={filterStatus}
        onFilterChange={setFilterStatus}
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

      {loading ? (
        <p className="loading-msg">Loading...</p>
      ) : (
        <TodoList
          todos={todos}
          onToggle={handleToggle}
          onEdit={setEditingTodo}
          onDelete={handleDelete}
        />
      )}
    </div>
  )
}

export default App
