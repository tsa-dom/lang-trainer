import axios from 'axios'
import { useState } from 'react'

const usePing = () => {
  const [result, setResult] = useState()
  const [loading, setLoading] = useState(true)

  const getPing = async () => {
    try {
      const res = await axios.get('http://localhost:8080/ping')
      setResult(res.data)
    } catch (err) {
      setResult({
        errors: err
      })
    }
    setLoading(false)
  }

  return { getPing, result, loading }
}

export default usePing