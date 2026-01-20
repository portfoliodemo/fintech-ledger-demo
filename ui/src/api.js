const BASE_URL = "http://localhost:8080"

export async function getBalance(userId) {
  const res = await fetch(`${BASE_URL}/balance?user_id=${userId}`)
  if (!res.ok) throw new Error("Failed to fetch balance")
  return res.json()
}

export async function credit(userId, amount) {
  return post("/credit", { user_id: userId, amount })
}

export async function debit(userId, amount) {
  return post("/debit", { user_id: userId, amount })
}

async function post(path, body) {
  const res = await fetch(`${BASE_URL}${path}`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(body)
  })

  if (!res.ok) {
    const text = await res.text()
    throw new Error(text)
  }

  return res.json()
}
