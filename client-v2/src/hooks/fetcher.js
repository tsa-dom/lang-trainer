import { useDispatch, useSelector } from 'react-redux'
import { setWordsToGroup, setGroupAsFetched, setGroups } from '../features/groupSlice'
import { setTemplates } from '../features/templateSlice'
import { setUser } from '../features/userSlice'
import { getGroups } from '../services/groups'
import { getTemplates } from '../services/templates'
import { authorize } from '../services/users'
import { getWordsInGroup } from '../services/words'

const useFetch = () => {
  const dispatch = useDispatch()
  const store = useSelector(state => state)

  const fetchWords = async group => {
    if (!group.fetched) {
      const words = await getWordsInGroup({
        id: group.id
      })
      dispatch(setWordsToGroup({
        words,
        groupId: group.id
      }))
      dispatch(setGroupAsFetched(group))
    }
  }

  const fetchGroups = async () => {
    if (!store.groups.fetched) {
      const groups = await getGroups()
      dispatch(setGroups(groups))
    }
  }

  const fetchTemplates = async () => {
    if(!store.templates.fetched) {
      const templates = await getTemplates()
      if (templates) dispatch(setTemplates(templates))
    }
  }

  const fetchUser = async () => {
    const token = localStorage.getItem('app-token')
    if(!store.users.currentUser && token) {
      const user = await authorize(token)
      if (user) dispatch(setUser(user))
      else localStorage.removeItem('app-token')
    }
  }

  return { fetchGroups, fetchWords, fetchTemplates, fetchUser }
}

export default useFetch