import React, { useState } from 'react'
import { useDispatch } from 'react-redux'
import { modifyWord as modify } from '../../features/groupSlice'
import { setNotification } from '../../features/notificationSlice'
import Form from '../Group/deprecated/Form'
import { modifyWord } from '../../services/words'

const ModifyWord = ({ word }) => {
  const [items, setItems] = useState(word.items.map(item => {
    return { name: item.name, description: item.description, id: item.id }
  }))
  const dispatch = useDispatch()

  const onSubmit = async (values) => {
    let itemsAreValid = true
    items.forEach(item => {
      if (item.name === '') itemsAreValid = false
    })
    if (itemsAreValid) {
      const modifiedWord = await modifyWord({
        ...values,
        id: word.id,
        items,
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
    <Form
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