import axios from 'axios'
import { useState } from 'react'
import { BACKEND_URL } from '../config'
const BASE_URL = BACKEND_URL + '/api/group/'

const useGroups = () => {
  const [groups, setGroups] = useState()

  const getGroups = async () => {
    try {
      const token = localStorage.getItem('app-token')
      const res = await axios.get(BASE_URL, {
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
      const res = await axios.post(BASE_URL, values, {
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