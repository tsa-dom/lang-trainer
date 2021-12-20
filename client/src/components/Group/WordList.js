import React, { useEffect, useState } from 'react'
import { useDispatch } from 'react-redux'
import { removeWords as remove } from '../../features/groupSlice'
import ItemList from '../ItemList'
import { useTranslation } from 'react-i18next'
import { removeWords } from '../../services/words'
import { unSelect } from '../../features/templateSlice'
import useFetch from '../../hooks/fetcher'
import { Button } from 'react-bootstrap'
import AddWord from './WordModal/AddWord'

const WordList = ({ group, setWord }) => {
  const { t } = useTranslation()
  const dispatch = useDispatch()
  const { fetchWords } = useFetch()
  const [showAddModal, setShowAddModal] = useState(false)

  useEffect(() => fetchWords(group), [])

  const columns = [
    { field: 'name', headerName: t('words-list-name'), width: 200 },
    { field: 'description', headerName: t('words-list-description') },
  ]

  const handleWordClick = (values) => {
    dispatch(unSelect())
    setWord(values)
  }

  const handleWordRemove = async (values) => {
    if (values.length <= 0) return
    const ids = await removeWords({ wordIds: values })
    if (ids) dispatch(remove(ids))
  }

  const handleAddWord = () => setShowAddModal(true)

  return (
    <>
      <ItemList
        rows={group.words}
        columns={columns}
        onCellClick={handleWordClick}
        handleItemRemove={handleWordRemove}
        title={`${t('words')} – ${group.name}`}
      />
      <AddWord
        open={showAddModal}
        setOpen={setShowAddModal}
      />
      <Button
        className='button-menu'
        onClick={handleAddWord}
      >
        {t('add-word')}
      </Button>
    </>
  )
}

export default WordList