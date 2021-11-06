import axios from 'axios'
import { BACKEND_URL } from '../config'
const BASE_URL = BACKEND_URL + '/api/user/'

const useUser = () => {

  const authorize = async (token) => {
    try {
      const res = await axios.get(BASE_URL, {
        headers: {
          'Authorization': `Bearer ${token}`
        }
      })
      return res.data
    } catch (err) {
      return {
        errors: err
      }
    }
  }

  return { authorize }
}

export default useUser