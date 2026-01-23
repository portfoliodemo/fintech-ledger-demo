import { useState } from "react"
import { getBalance, addCredit, addDebit } from "./api"

function App() {
  const [userId, setUserId] = useState(1)
  const [amount, setAmount] = useState("")
  const [balance, setBalance] = useState<number | null>(null)
  const [error, setError] = useState("")

  const loadBalance = async () => {
    try {
      setError("")
      const balance = await getBalance(userId)
      setBalance(balance)
    } catch (err) {
      setError((err as Error).message)
    }
  }
  

  const handleCredit = async () => {
    try {
      setError("")
      await addCredit(userId, Number(amount))
      await loadBalance()
    } catch (err) {
      setError((err as Error).message)
      }    
  }

  const handleDebit = async () => {
    try {
      setError("")
      await addDebit(userId, Number(amount))
      await loadBalance()
    } catch (err) {
      setError((err as Error).message)
      }    
  }

  return (
    <div>
      <h2>Fintech Ledger Demo</h2>

      <label>User ID</label>
      <input
        type="number"
        value={userId}
        onChange={e => setUserId(Number(e.target.value))}
      />

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
