import axios from 'axios'
import { BACKEND_URL, getHeader } from '../config'
const TEACHER_URL = BACKEND_URL + '/teacher/words/'
const USER_URL = BACKEND_URL + '/my/words/'

const useWords = () => {
  const getWordsInGroup = async (values) => {
    try {
      const res = await axios.post(USER_URL, values, {
        headers: getHeader()
      })
      return res.data.words
    } catch (err) {
      console.error(err)
      return null
    }
  }

  const addWordToGroup = async (values) => {
    try {
      const res = await axios.post(TEACHER_URL, values, {
        headers: getHeader()
      })
      return res.data.word
    } catch (err) {
      console.error(err)
      return null
    }
  }

  const removeWords = async (values) => {
    try {
      const res = await axios.delete(TEACHER_URL, {
        headers: getHeader(),
        data: values
      })
      return res.data.wordIds
    } catch (err) {
      console.error(err)
      return null
    }
  }

  const modifyWord = async (values) => {
    try {
      const res = await axios.put(TEACHER_URL, values, {
        headers: getHeader()
      })
      return res.data.word
    } catch (err) {
      console.error(err)
      return null
    }
  }

  return { getWordsInGroup, addWordToGroup, removeWords, modifyWord }
}

export default useWords