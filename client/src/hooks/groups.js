import axios from 'axios'
import { BACKEND_URL, getHeader } from '../config'
const USER_URL = BACKEND_URL + '/my/groups/'
const TEACHER_URL = BACKEND_URL + '/teacher/groups/'

const useGroups = () => {
  const getGroups = async () => {
    try {
      const res = await axios.get(USER_URL, {
        headers: getHeader()
      })
      return res.data.groups
    } catch (err) {
      return null
    }
  }

  const addGroup = async (values) => {
    try {
      const res = await axios.post(TEACHER_URL, values, {
        headers: getHeader()
      })
      return res.data.group
    } catch (err) {
      return null
    }
  }

  const modifyGroup = async (values) => {
    try {
      const res = await axios.put(TEACHER_URL, values, {
        headers: getHeader()
      })
      return res.data.group
    } catch (err) {
      return null
    }
  }

  const removeGroups = async (values) => {
    try {
      const res = await axios.delete(TEACHER_URL, {
        headers: getHeader(),
        data: values
      })
      return res.data.groupIds
    } catch (err) {
      return null
    }
  }

  return { getGroups, addGroup, removeGroups, modifyGroup }
}

export default useGroups