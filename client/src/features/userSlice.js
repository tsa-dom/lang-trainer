import { createSlice } from '@reduxjs/toolkit'

export const userSlice = createSlice({
  name: 'users',
  initialState: {
    currentUser: null,
    language: 'fi'
  },
  reducers: {
    setUser: (state, user) => {
      state.currentUser = user.payload
    },
    clearUser: state => {
      state.currentUser = null
    },
    setLanguage: (state, language) => {
      state.language = language.payload
    }
  }
})

export const { setUser, clearUser, setLanguage } = userSlice.actions

export default userSlice.reducer