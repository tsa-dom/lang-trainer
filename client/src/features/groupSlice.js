import { createSlice } from '@reduxjs/toolkit'

export const groupSlice = createSlice({
  name: 'groups',
  initialState: {
    values: [],
    fetched: false,
    selectedGroup: null
  },
  reducers: {
    setGroups: (state, groups) => {
      state.values = groups.payload
      state.fetched = true
    },
    addGroup: (state, group) => {
      state.values = state.values.concat(group.payload)
    },
    setSelectedGroup: (state, group) => {
      state.selectedGroup = group.payload
    }
  }
})

export const { setGroups, addGroup, setSelectedGroup } = groupSlice.actions

export default groupSlice.reducer