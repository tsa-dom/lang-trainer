import { createSlice } from '@reduxjs/toolkit'

export const groupSlice = createSlice({
  name: 'groups',
  initialState: {
    values: [],
    fetched: false
  },
  reducers: {
    setGroups: (state, groups) => {
      state.values = groups.payload
      state.fetched = true
    }
  }
})

export const { setGroups } = groupSlice.actions

export default groupSlice.reducer