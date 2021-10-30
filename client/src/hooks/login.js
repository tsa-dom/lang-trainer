import axios from 'axios'
import { useState } from 'react'

const useLogin = () => {
  const [result, setResult] = useState()
  const [loading, setLoading] = useState(true)

  const login = async (username, password) => {
    try {
      const res = await axios.post('http://localhost:8080/api/user/login/', {
        username,
        password
      })
      setResult(res.data)
    } catch (err) {
      setResult({
        errors: err
      })
    }
    setLoading(false)
  }

  return { login, result, loading }
}

export default useLogin