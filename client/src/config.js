export const BACKEND_URL = process.env.REACT_APP_API_URI ? `${process.env.REACT_APP_API_URI}/api` : 'http://localhost:8080/api'

export const getHeader = () => {
  const token = localStorage.getItem('app-token')
  return {
    'Authorization': `Bearer ${token}`
  }
}