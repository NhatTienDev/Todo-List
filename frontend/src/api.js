const BASE_URL = '/api/v1/todos'

export async function getTodos(search = '', status = '', page = 1, limit = 10) {
  const params = new URLSearchParams()
  if (search) params.set('search', search)
  if (status) params.set('status', status)
  params.set('page', page)
  params.set('limit', limit)
  const res = await fetch(`${BASE_URL}?${params}`)
  const json = await res.json()
  if (!res.ok) throw new Error(json.error || 'Failed to fetch todos')
  return json.data
}

export async function getTodoByID(id) {
  const res = await fetch(`${BASE_URL}/${id}`)
  const json = await res.json()
  if (!res.ok) throw new Error(json.error || 'Todo not found')
  return json.data
}

export async function createTodo(title, description) {
  const res = await fetch(BASE_URL, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ title, description }),
  })
  const json = await res.json()
  if (!res.ok) throw new Error(json.error || 'Failed to create todo')
  return json.data
}

export async function updateTodo(id, updates) {
  const res = await fetch(`${BASE_URL}/${id}`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(updates),
  })
  const json = await res.json()
  if (!res.ok) throw new Error(json.error || 'Failed to update todo')
  return json.data
}

export async function deleteTodo(id) {
  const res = await fetch(`${BASE_URL}/${id}`, { method: 'DELETE' })
  const json = await res.json()
  if (!res.ok) throw new Error(json.error || 'Failed to delete todo')
  return json
}
