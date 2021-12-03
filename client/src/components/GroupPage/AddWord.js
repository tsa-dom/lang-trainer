import React, { useState } from 'react'
import useWords from '../../hooks/words'
import { useDispatch } from 'react-redux'
import { addWordToSelectedGroup } from '../../features/groupSlice'
import { setNotification } from '../../features/notificationSlice'
import FormBody from './FormBody'

const AddWord = ({ setSelected, group }) => {
  const [items, setItems] = useState([])
  const { addWordToGroup } = useWords()
  const dispatch = useDispatch()

  const onSubmit = async (values) => {
    let itemsAreValid = true
    items.forEach(item => {
      if (item.name === '') itemsAreValid = false
    })
    if (itemsAreValid) {
      const body = {
        ...values,
        groupId: group.id,
        items,
      }
      const word = await addWordToGroup(body)
      if (word) {
        dispatch(addWordToSelectedGroup(word))
        dispatch(setNotification({
          message: 'Added a new word successfully',
          type: 'success'
        }))
        setSelected('group-word-list')
      } else {
        dispatch(setNotification({
          message: 'Server error',
          type: 'error'
        }))
      }
    }
  }

  return (
    <FormBody
      onSubmit={onSubmit}
      items={items}
      setItems={setItems}
      initialValues={{
        name: '',
        description: ''
      }}
    />
  )
}

export default AddWord