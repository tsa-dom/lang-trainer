import React, { useState } from 'react'
import { Formik } from 'formik'
import { useTranslation } from 'react-i18next'
import TextField from '@mui/material/TextField'
import { Button } from '@mui/material'
import SendIcon from '@mui/icons-material/Send'
import useWords from '../../hooks/words'
import { useDispatch } from 'react-redux'
import { addWordToSelectedGroup } from '../../features/groupSlice'

const AddWord = ({ setSelected, group }) => {
  const { t } = useTranslation('translation')
  const [items, setItems] = useState([])
  const { addWordToGroup } = useWords()
  const dispatch = useDispatch()

  const validate = () => {}

  const onSubmit = async (values) => {
    const body = {
      ...values,
      groupId: group.id,
      items: items.map(item => {
        delete item.id
        return item
      })
    }
    const word = await addWordToGroup(body)
    dispatch(addWordToSelectedGroup(word))
    setSelected('group-word-list')
  }

  const handleAddItem = () => {
    const newItem = {
      id: items.length ? items[items.length - 1].id + 1 : 1,
      name: '',
      description: ''
    }
    setItems(items.concat(newItem))
  }

  const handleRemoveItem = (id) => {
    setItems(items.filter(item => item.id !== id))
  }

  const handleModifyItem = (id, event, fieldName) => {
    setItems(items.map(item => {
      if (item.id === id) {
        item[fieldName] = event.target.value
      }
      return item
    }))
  }

  return (
    <Formik
      validate={validate}
      initialValues={{
        name: '',
        description: ''
      }}
      onSubmit={onSubmit}
    >
      {({ handleSubmit, handleChange }) => {
        return (
          <div className="words-add-body">
            <div className="words-add-header">{t('word-info')}</div>
            <TextField
              id="name"
              required
              variant="standard"
              label={t('name')}
              style={{ marginRight: 30 }}
              onChange={handleChange}
            />
            <TextField
              id="description"
              variant="standard"
              label={t('description')}
              multiline
              onChange={handleChange}
              style={{ width: 400 }}
            />
            <div style={{ marginTop: 20, marginBottom: 20 }}>
              <Button
                variant="outlined"
                style={{ minWidth: 150 }}
                onClick={handleAddItem}
              >
                {t('add-item')}
              </Button>
            </div>
            {items.length > 0 &&
              <div style={{ marginBottom: 20 }} className="words-item-header">{t('word-items')}</div>
            }
            {items.map(item => {
              return (
                <div key={item} style={{ display: 'flex' }}>
                  <TextField
                    required
                    variant="standard"
                    label={t('name')}
                    style={{ marginRight: 30, marginBottom: 10 }}
                    onChange={value => handleModifyItem(item.id, value, 'name')}
                  />
                  <TextField
                    variant="standard"
                    label={t('description')}
                    multiline
                    style={{ marginRight: 30, width: 400 }}
                    onChange={event => handleModifyItem(item.id, event, 'description')}
                  />
                  <Button
                    onClick={() => handleRemoveItem(item.id)}
                    color="error"
                  >
                    {t('word-remove-item')}
                  </Button>
                </div>
              )
            })}
            <Button
              style={{ marginTop: 30, minWidth: 150, backgroundColor: 'rgb(5, 23, 71)', color: 'white' }}
              variant="contained"
              endIcon={<SendIcon />}
              onClick={handleSubmit}
            >
              {t('words-add-word')}
            </Button>
          </div>
        )
      }}
    </Formik>
  )
}

export default AddWord