import axios from 'axios'
import { useState } from 'react'

const useGroups = () => {
  const [groups, setGroups] = useState()
  const [loading, setLoading] = useState(true)

  const getGroups = async () => {
    try {
      const token = localStorage.getItem('app-token')
      const res = await axios.get('http://localhost:8080/api/group/', {
        headers: {
          'Authorization': `Bearer ${token}`
        }
      })
      setGroups(res.data.groups)
    } catch (err) {
      setGroups({
        errors: err
      })
    }
    setLoading(false)
  }

  return { getGroups, groups, loading }
}

export default useGroups