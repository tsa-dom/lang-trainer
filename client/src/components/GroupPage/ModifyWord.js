import React, { useState } from 'react'
import useWords from '../../hooks/words'
import { useDispatch } from 'react-redux'
import { addWordToSelectedGroup } from '../../features/groupSlice'
import { setNotification } from '../../features/notificationSlice'
import FormBody from './FormBody'

const ModifyWord = ({ word, setSelected, group }) => {
  const [items, setItems] = useState(word.items.map(item => {
    return { name: item.name, description: item.description, id: item.id }
  }))
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
        name: word.name,
        description: word.description
      }}
    />
  )
}

export default ModifyWord