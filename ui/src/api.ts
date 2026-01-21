const API_BASE = "http://localhost:8080";

export async function addCredit(userId: number, amount: number): Promise<void> {
  const res = await fetch(`${API_BASE}/credit`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ user_id: userId, amount }),
  });

  if (!res.ok) {
    throw new Error(await res.text());
  }
}

export async function addDebit(userId: number, amount: number): Promise<void> {
  const res = await fetch(`${API_BASE}/debit`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ user_id: userId, amount }),
  });

  if (!res.ok) {
    throw new Error(await res.text());
  }
}

export async function getBalance(userId: number): Promise<number> {
  const res = await fetch(`${API_BASE}/balance?user_id=${userId}`);
  if (!res.ok) {
    throw new Error(await res.text());
  }

  const data = await res.json();
  return data.balance;
}
