import axios from 'axios'
import { useState } from 'react'
import { BACKEND_URL } from '../config'
const BASE_URL = BACKEND_URL + '/api/user/'

const useUser = () => {
  const [user, setUser] = useState()
  const [loading, setLoading] = useState(true)

  const authorize = async (token) => {
    try {
      const res = await axios.get(BASE_URL, {
        headers: {
          'Authorization': `Bearer ${token}`
        }
      })
      setUser(res.data)
    } catch (err) {
      setUser({
        errors: err
      })
    }
    setLoading(false)
  }

  return { authorize, user, loading }
}

export default useUser