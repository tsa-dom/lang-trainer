import axios from 'axios'
import { useState } from 'react'

const useLogin = () => {
  const [user, setUser] = useState()
  const [loading, setLoading] = useState(true)

  const login = async (username, password) => {
    try {
      const res = await axios.post('http://localhost:8080/api/user/login/', {
        username,
        password
      })
      setUser(res.data)
    } catch (err) {
      setUser({
        errors: err
      })
    }
    setLoading(false)
  }

  return { login, user, loading }
}

export default useLogin