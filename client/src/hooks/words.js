import axios from 'axios'
import { BACKEND_URL } from '../config'
const TEACHER_URL = BACKEND_URL + '/api/teacher'
const USER_URL = BACKEND_URL + '/api/my'

const getHeader = () => {
  const token = localStorage.getItem('app-token')
  return {
    'Authorization': `Bearer ${token}`
  }
}

const useWords = () => {
  const getWordsInGroup = async (values) => {
    try {
      const res = await axios.post(`${USER_URL}/words/`, values, {
        headers: getHeader()
      })
      return res.data.words
    } catch (err) {
      return null
    }
  }

  const addWordToGroup = async (values) => {
    try {
      const res = await axios.post(`${TEACHER_URL}/words/`, values, {
        headers: getHeader()
      })
      return res.data.word
    } catch (err) {
      return null
    }
  }

  const removeWords = async (values) => {
    try {
      const res = await axios.delete(`${TEACHER_URL}/words/`, {
        headers: getHeader(),
        data: values
      })
      return res.data.wordIds
    } catch (err) {
      return null
    }
  }

  return { getWordsInGroup, addWordToGroup, removeWords }
}

export default useWords