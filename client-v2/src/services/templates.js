import axios from 'axios'
import { BACKEND_URL, getHeader } from '../config'
const TEACHER_URL = BACKEND_URL + '/teacher/templates/'

export const getTemplates = async () => {
  try {
    const res = await axios.get(TEACHER_URL, {
      headers: getHeader()
    })
    return res.data.templates
  } catch (err) {
    return null
  }
}

export const addTemplate = async (values) => {
  try {
    const res = await axios.post(TEACHER_URL, values, {
      headers: getHeader()
    })
    return res.data.template
  } catch (err) {
    return null
  }
}

export const modifyTemplate = async (values) => {
  try {
    const res = await axios.put(TEACHER_URL, values, {
      headers: getHeader()
    })
    return res.data.template
  } catch (err) {
    return null
  }
}

export const removeTemplates = async (values) => {
  try {
    const res = await axios.delete(TEACHER_URL, {
      headers: getHeader(),
      data: values
    })
    return res.data.templateIds
  } catch (err) {
    return null
  }
}