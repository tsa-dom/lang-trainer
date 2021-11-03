import { configureStore } from '@reduxjs/toolkit'
import groupReducer from './features/groupSlice'
import userReducer from './features/userSlice'

export default configureStore({
  reducer: {
    groups: groupReducer,
    users: userReducer
  }
})