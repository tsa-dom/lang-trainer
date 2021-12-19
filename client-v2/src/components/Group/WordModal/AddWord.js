import React from 'react'
import { useDispatch, useSelector } from 'react-redux'
import WordModal from '.'
import { addWordToSelectedGroup } from '../../../features/groupSlice'
import { addWordToGroup } from '../../../services/words'
import { setNotification } from '../../../features/notificationSlice'
import { useTranslation } from 'react-i18next'

const AddWord = ({ open, setOpen }) => {
  const group = useSelector(state => state.groups.selectedGroup)
  const dispatch = useDispatch()
  const { t } = useTranslation()

  const initialValues = {
    name: '',
    description: '',
    items: []
  }

  const onSubmit = async (values) => {
    let itemsAreValid = true
    values.items.forEach(item => {
      if (item.name === '') itemsAreValid = false
    })
    if (itemsAreValid) {
      const body = {
        ...values,
        groupId: group.id,
      }
      const word = await addWordToGroup(body)
      if (word) {
        dispatch(addWordToSelectedGroup(word))
        dispatch(setNotification({
          message: 'Added a new word successfully',
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

  const handleClose = () => {
    setOpen(false)
  }

  return (
    <WordModal
      initialValues={initialValues}
      onSubmit={onSubmit}
      open={open}
      onClose={handleClose}
      title={t('add-word')}
    />
  )
}

export default AddWord
