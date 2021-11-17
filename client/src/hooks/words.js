import axios from 'axios'
import { BACKEND_URL } from '../config'
const TEACHER_URL = BACKEND_URL + '/api/teacher'
const USER_URL = BACKEND_URL + '/api/my'

const useWords = () => {
  const getWordsInGroup = async (values) => {
    try {
      const token = localStorage.getItem('app-token')
      const res = await axios.post(`${USER_URL}/words/`, values, {
        headers: {
          'Authorization': `Bearer ${token}`
        }
      })
      return res.data.words
    } catch (err) {
      return null
    }
  }

  const addWordToGroup = async (values) => {
    try {
      const token = localStorage.getItem('app-token')
      const res = await axios.post(`${TEACHER_URL}/word/`, values, {
        headers: {
          'Authorization': `Bearer ${token}`
        }
      })
      return res.data.word
    } catch (err) {
      return null
    }
  }

  return { getWordsInGroup, addWordToGroup }
}

export default useWords