import { createSlice } from '@reduxjs/toolkit'

export const notificationSlice = createSlice({
  name: 'notifications',
  initialState: {
    message: '',
    type: '',
  },
  reducers: {
    setNotification: (state, notification) => {
      const { message, type } = notification.payload
      state.message = message
      state.type = type
    },
    resetNotification: (state) => {
      state.message = ''
      state.type = ''
    }
  }
})

export const {
  setNotification,
  resetNotification
} = notificationSlice.actions

export default notificationSlice.reducer