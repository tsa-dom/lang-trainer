import axios from 'axios'
import { BACKEND_URL } from '../config'
const BASE_URL = BACKEND_URL + '/api/my/groups/'

const useGroups = () => {
  const getGroups = async () => {
    try {
      const token = localStorage.getItem('app-token')
      const res = await axios.get(BASE_URL, {
        headers: {
          'Authorization': `Bearer ${token}`
        }
      })
      return res.data.groups
    } catch (err) {
      return null
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

  return { getGroups, addGroup }
}

export default useGroups