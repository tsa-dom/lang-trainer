import { createSlice } from '@reduxjs/toolkit'

export const templateSlice = createSlice({
  name: 'templates',
  initialState: {
    values: [],
    selected: undefined,
    fetched: false
  },
  reducers: {
    setTemplates: (state, templates) => {
      state.values = templates.payload
      state.fetched = true
    },
    addTemplate: (state, template) => {
      state.values = state.values.concat(template.payload)
    },
    setSelected: (state, template) => {
      state.selected = template.payload
    },
    modifyTemplate: (state, template) => {
      state.values = state.values.map(t => {
        if (t.id === template.payload.id) return template.payload
        else return t
      })
    },
    removeTemplates: (state, ids) => {
      state.values = state.values.filter(template => !ids.payload.includes(template.id))
      if (state.selected && ids.payload.includes(state.selected.id)) {
        state.selected = null
      }
    },
    unSelect: state => {
      state.selected = null
    }
  }
})

export const {
  setSelected,
  setTemplates,
  addTemplate,
  removeTemplates,
  unSelect,
  modifyTemplate
} = templateSlice.actions

export default templateSlice.reducer
