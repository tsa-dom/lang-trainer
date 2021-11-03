import axios from 'axios'
import { useState } from 'react'

const useGroups = () => {
  const [groups, setGroups] = useState()

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
  }

  const addGroup = async (values) => {
    try {
      const token = localStorage.getItem('app-token')
      const res = await axios.post('http://localhost:8080/api/group/', values, {
        headers: {
          'Authorization': `Bearer ${token}`
        },
      })
      return res.data.group
    } catch (err) {
      return null
    }
  }

  return { getGroups, addGroup, groups }
}

export default useGroups