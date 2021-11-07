import axios from 'axios'
import { useState } from 'react'
import { BACKEND_URL } from '../config'
const BASE_URL = BACKEND_URL + '/api'

const useLogin = () => {
  const [user, setUser] = useState()
  const [loading, setLoading] = useState(true)

  const login = async (username, password) => {
    try {
      const res = await axios.post(`${BASE_URL}/login/`, {
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