import axios from 'axios'
import { BACKEND_URL } from '../config'
const USER_URL = BACKEND_URL + '/api/my/groups/'
const TEACHER_URL = BACKEND_URL + '/api/teacher/groups/'

const useGroups = () => {
  const getGroups = async () => {
    try {
      const token = localStorage.getItem('app-token')
      const res = await axios.get(USER_URL, {
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
      const res = await axios.post(TEACHER_URL, values, {
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