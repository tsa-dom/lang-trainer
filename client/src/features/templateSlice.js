import { createSlice } from '@reduxjs/toolkit'

export const templateSlice = createSlice({
  name: 'templates',
  initialState: {
    values: [],
    selected: null
  },
  reducers: {
    setTemplates: (state, templates) => {
      state.values = templates.payload
    },
    addTemplate: (state, template) => {
      state.values = state.values.concat(template.payload)
    },
    setSelected: (state, template) => {
      state.selected = template.payload
    }
  }
})

export const {
  setSelected,
  setTemplates,
  addTemplate
} = templateSlice.actions

export default templateSlice.reducer