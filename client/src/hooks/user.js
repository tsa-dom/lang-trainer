import axios from 'axios'
import { useState } from 'react'

const useUser = () => {
  const [user, setUser] = useState()
  const [loading, setLoading] = useState(true)

  const authorize = async (token) => {
    try {
      const res = await axios.get('http://localhost:8080/api/user/', {
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