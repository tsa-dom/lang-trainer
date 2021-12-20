import axios from 'axios'
import { BACKEND_URL } from '../config'
const BASE_URL = BACKEND_URL

export const login = async (username, password) => {
  try {
    const res = await axios.post(`${BASE_URL}/login/`, {
      username,
      password
    })
    return res.data
  } catch (err) {
    return null
  }
}

export const authorize = async (token) => {
  try {
    const res = await axios.get(`${BASE_URL}/my/`, {
      headers: {
        'Authorization': `Bearer ${token}`
      }
    })
    return res.data.user
  } catch (err) {
    return null
  }
}