import React from 'react'
import { useTranslation } from 'react-i18next'
import { useDispatch } from 'react-redux'
import WordModal from '.'
import { modifyWord as modify } from '../../../features/groupSlice'
import { setNotification } from '../../../features/notificationSlice'
import { modifyWord } from '../../../services/words'

const ModifyWord = ({ word, setWord }) => {
  const dispatch = useDispatch()
  const { t } = useTranslation()

  const handleClose = () => {
    setWord(undefined)
  }

  const initialValues = {
    name: word.name,
    description: word.description
  }

  const onSubmit = async (values) => {
    let itemsAreValid = true
    values.items.forEach(item => {
      if (item.name === '') itemsAreValid = false
    })
    if (itemsAreValid) {
      const modifiedWord = await modifyWord({
        ...values,
        id: word.id
      })
      if (modifiedWord) {
        dispatch(modify(modifiedWord))
        dispatch(setNotification({
          message: 'A word modified successfully',
          type: 'success'
        }))
      } else {
        dispatch(setNotification({
          message: 'Server error',
          type: 'error'
        }))
      }
    }
  }

  return (
    <WordModal
      initialValues={initialValues}
      onSubmit={onSubmit}
      open={word ? true : false}
      onClose={handleClose}
      title={`${t('word')} - ${word.name}`}
      initialItems={word.items.map(item => {
        return { name: item.name, description: item.description, id: item.id }
      })}
    />
  )
}

export default ModifyWord
