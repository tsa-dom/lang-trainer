import { createSlice } from '@reduxjs/toolkit'

export const userSlice = createSlice({
  name: 'users',
  initialState: {
    currentUser: null
  },
  reducers: {
    setUser: (state, user) => {
      state.currentUser = user.payload
    },
    clearUser: state => {
      state.currentUser = null
    }
  }
})

export const { setUser, clearUser } = userSlice.actions

export default userSlice.reducer