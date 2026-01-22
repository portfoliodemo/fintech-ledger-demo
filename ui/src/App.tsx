import { useState } from "react"
import { getBalance, addCredit, addDebit } from "./api"

function App() {
  const [userId, setUserId] = useState(1)
  const [amount, setAmount] = useState("")
  const [balance, setBalance] = useState(null)
  const [error, setError] = useState("")

  const loadBalance = async () => {
    try {
      setError("")
      const data = await getBalance(userId)
      setBalance(data.balance)
    } catch (err) {
      setError(err.message)
    }
  }

  const handleCredit = async () => {
    try {
      setError("")
      await addCredit(userId, Number(amount))
      await loadBalance()
    } catch (err) {
      setError(err.message)
    }
  }

  const handleDebit = async () => {
    try {
      setError("")
      await addDebit(userId, Number(amount))
      await loadBalance()
    } catch (err) {
      setError(err.message)
    }
  }

  return (
    <div>
      <h2>Fintech Ledger Demo</h2>

      <label>User ID</label>
      <input value={userId} onChange={e => setUserId(e.target.value)} />

      <label>Amount</label>
      <input value={amount} onChange={e => setAmount(e.target.value)} />

      <div style={{ marginTop: 12 }}>
        <button onClick={handleCredit}>Credit</button>
        <button onClick={handleDebit}>Debit</button>
        <button onClick={loadBalance}>Get Balance</button>
      </div>

      {balance !== null && <p>Balance: {balance}</p>}
      {error && <p style={{ color: "red" }}>{error}</p>}
    </div>
  )
}

export default App
