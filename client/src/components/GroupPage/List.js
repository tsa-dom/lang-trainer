import React, { useEffect } from 'react'
import { useDispatch } from 'react-redux'
import useWords from '../../hooks/words'
import { setWordsToGroup, setGroupAsFetched } from '../../features/groupSlice'
import ItemList from '../Styled/ItemList'
import { useTranslation } from 'react-i18next'

const List = ({ group }) => {
  const { getWordsInGroup } = useWords()
  const { t } = useTranslation()
  const dispatch = useDispatch()

  useEffect(async () => {
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
  }, [])

  const columns = [
    { field: 'name', headerName: t('words-list-name'), flex: 1 },
    { field: 'description', headerName: t('words-list-description'), flex: 3 },
  ]

  return (
    <ItemList
      rows={group.words}
      columns={columns}
    />
  )
}

export default List