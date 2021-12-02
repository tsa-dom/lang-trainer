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
    modifyGroup: (state, group) => {
      state.selectedGroup = {
        ...state.selectedGroup,
        ...group.payload
      }
      state.values = state.values.map(g => {
        if (g.id !== group.payload.id) return g
        else return {
          ...g,
          ...group.payload
        }
      })
    },
    removeGroups: (state, ids) => {
      state.values = state.values.filter(group => !ids.payload.includes(group.id))
    },
    setSelectedGroup: (state, group) => {
      state.selectedGroup = group.payload
    },
    setWordsToGroup: (state, group) => {
      const words = group.payload.words
      const groupId = group.payload.groupId
      state.values = state.values.map(value => {
        if (groupId === value.id) {
          return {
            ...value,
            words: words
          }
        } else {
          return { ...value }
        }
      })
      state.selectedGroup = groupId === state.selectedGroup.id ? {
        ...state.selectedGroup,
        words
      } : state.selectedGroup
    },
    addWordToSelectedGroup: (state, word) => {
      const group = state.selectedGroup
      group.words.push(word.payload)
      state.selectedGroup = group
    }
  }
})

export const {
  setGroups,
  addGroup,
  setSelectedGroup,
  setWordsToGroup,
  addWordToSelectedGroup,
  removeGroups,
  modifyGroup
} = groupSlice.actions

export default groupSlice.reducer