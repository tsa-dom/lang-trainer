import { configureStore } from '@reduxjs/toolkit'
import groupReducer from './features/groupSlice'
import userReducer from './features/userSlice'
import notificationReducer from './features/notificationSlice'
import templateReducer from './features/templateSlice'

export default configureStore({
  reducer: {
    groups: groupReducer,
    users: userReducer,
    notifications: notificationReducer,
    templates: templateReducer,
  }
})