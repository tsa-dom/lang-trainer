import store from '..'
import { setWordsToGroup, setGroupAsFetched, setGroups } from '../features/groupSlice'
import { setTemplates } from '../features/templateSlice'
import { setUser } from '../features/userSlice'
import { getGroups } from '../services/groups'
import { getTemplates } from '../services/templates'
import { authorize } from '../services/users'
import { getWordsInGroup } from '../services/words'

export const fetchWords = async group => {
  if (!group.fetched) {
    const words = await getWordsInGroup({
      id: group.id
    })
    store.dispatch(setWordsToGroup({
      words,
      groupId: group.id
    }))
    store.dispatch(setGroupAsFetched(group))
  }
}

export const fetchGroups = async () => {
  if (!store.getState().groups.fetched) {
    const groups = await getGroups()
    store.dispatch(setGroups(groups))
  }
}

export const fetchTemplates = async () => {
  if(!store.getState().templates.fetched) {
    const templates = await getTemplates()
    if (templates) store.dispatch(setTemplates(templates))
  }
}

export const fetchUser = async () => {
  const token = localStorage.getItem('app-token')
  if(!store.getState().users.currentUser && token) {
    const user = await authorize(token)
    if (user) store.dispatch(setUser(user))
    else localStorage.removeItem('app-token')
  }
}