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
    },
    addGroup: (state, group) => {
      state.values = state.values.concat(group.payload)
    }
  }
})

export const { setGroups, addGroup } = groupSlice.actions

export default groupSlice.reducer