import React, { useEffect } from 'react'
import { useDispatch } from 'react-redux'
import useWords from '../../hooks/words'
import { setWordsToGroup, setGroupAsFetched } from '../../features/groupSlice'
import { removeWords as remove } from '../../features/groupSlice'
import ItemList from '../Styled/ItemList'
import { useTranslation } from 'react-i18next'

const List = ({ group, setSelectedWord, setSelectedPage }) => {
  const { getWordsInGroup, removeWords } = useWords()
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

  const handleWordClick = (values) => {
    setSelectedWord(values.row)
    setSelectedPage(values.row.name)
  }

  const handleWordRemove = async (values) => {
    if (values.length <= 0) return
    const ids = await removeWords({ wordIds: values })
    if (ids) dispatch(remove(ids))
  }

  return (
    <ItemList
      rows={group.words}
      columns={columns}
      onCellClick={handleWordClick}
      handleItemRemove={handleWordRemove}
    />
  )
}

export default List