import { Dialog, TextField } from '@mui/material'
import React, { useEffect, useState } from 'react'
import { useTranslation } from 'react-i18next'
import { Button } from '@material-ui/core'
import SendIcon from '@mui/icons-material/Send'
import { Formik } from 'formik'
import { useDispatch, useSelector } from 'react-redux'
import { setNotification } from '../../features/notificationSlice'
import { addTemplate as add, setSelected } from '../../features/templateSlice'
import { addTemplate } from '../../services/templates'
import { fetchTemplates } from '../../utils/fetcher'
import Select from 'react-select'

const FormBody = ({
  onSubmit,
  initialValues,
  items,
  setItems
}) => {
  const { t } = useTranslation()
  const [openDialog, setOpenDialog] = useState(false)
  const [templateName, setTemplateName] = useState('')
  const dispatch = useDispatch()
  const template = useSelector(state => state.templates.selected)
  const templates = useSelector(state => state.templates.values)

  useEffect(() => {
    if (template) {
      setItems(template.descriptions.map((description, id) => {
        return {
          description,
          name: '',
          id
        }
      }))
    } else setItems([])
  }, [template])

  const validate = () => {}

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

  const handleAddTemplate = async () => {
    const template = await addTemplate({
      name: templateName,
      descriptions: items.map(item => item.description)
    })
    if (template) {
      dispatch(add(template))
      dispatch(setNotification({
        message: 'Added a new template successfully',
        type: 'success'
      }))
      setTemplateName('')
      setOpenDialog(false)
    } else {
      dispatch(setNotification({
        message: 'Server error',
        type: 'error'
      }))
    }
  }

  const handleTemplateSelect = id => {
    const select = templates.find(t => t.id === id)
    dispatch(setSelected(select))
  }

  return (
    <Formik
      validate={validate}
      initialValues={initialValues}
      onSubmit={onSubmit}
    >
      {({ handleSubmit, handleChange, values }) => {
        return (
          <div className="words-body">
            <div style={{ fontWeight: 'bold' }} className="words-header">{t('template')}</div>
            <div style={{ width: 250 }}>
              <Select
                cacheOptions
                onMenuOpen={fetchTemplates}
                defaultValue={template ? { value: template.id, label: template.name } : undefined}
                options={[{ value: '', label: <em>None</em> }].concat(templates.map(t => {
                  return { value: t.id, label: t.name }
                }))}
                styles={{
                  control: styles => ({
                    ...styles,
                    padding: 5,
                    backgroundColor: 'rgb(231, 233, 240)',
                    borderColor: 'black',
                    '&:hover': {
                      borderColor: 'black'
                    }
                  }),
                  option: styles => ({ ...styles })
                }}
                onChange={e => handleTemplateSelect(e.value)}
              />
            </div>

            {/* <FormControl>
              <Select
                value={template ? template.id : ''}
                onChange={e => handleTemplateSelect(e.target.value)}
                onOpen={fetchTemplates}
                displayEmpty
                inputProps={{ 'aria-label': 'Without label' }}
                style={{ marginBottom: 20, width: 200 }}
              >
                <MenuItem value="">
                  <em>None</em>
                </MenuItem>
                {templates.map(t => {
                  return <MenuItem key={t.id} value={t.id}>{t.name}</MenuItem>
                })}
              </Select>
            </FormControl> */}
            <div style={{ fontWeight: 'bold', marginTop: 20 }} className="words-header">{t('word-info')}</div>
            <TextField
              id="name"
              required
              defaultValue={values.name}
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
              defaultValue={values.description}
              onChange={handleChange}
              style={{ width: 400 }}
            />
            <div style={{ marginTop: 20, marginBottom: 20 }}>
              <Button
                variant="outlined"
                style={{ minWidth: 150, color: 'rgb(5, 23, 71)', borderColor: 'rgb(5, 23, 71)' }}
                onClick={handleAddItem}
              >
                {t('add-item')}
              </Button>
            </div>
            {items.length > 0 &&
              <div style={{ marginBottom: 15, fontWeight: 'bold' }} className="words-item-header">{t('word-items')}</div>
            }
            {items.map(item => {
              return (
                <div key={item.id} style={{ display: 'flex' }}>
                  <TextField
                    required
                    defaultValue={item.name}
                    variant="standard"
                    label={t('name')}
                    style={{ marginRight: 30, marginBottom: 10 }}
                    onChange={value => handleModifyItem(item.id, value, 'name')}
                  />
                  <TextField
                    variant="standard"
                    label={t('description')}
                    multiline
                    defaultValue={item.description}
                    style={{ marginRight: 30, width: 400 }}
                    onChange={event => handleModifyItem(item.id, event, 'description')}
                  />
                  <Button
                    onClick={() => handleRemoveItem(item.id)}
                    style={{ color: 'red' }}
                  >
                    {t('word-remove-item')}
                  </Button>
                </div>
              )
            })}
            <div style={{ marginTop: 30 }}>
              <Button
                style={{ minWidth: 150, backgroundColor: 'rgb(5, 23, 71)', color: 'white' }}
                variant="contained"
                endIcon={<SendIcon />}
                onClick={handleSubmit}
              >
                {t('words-add-word')}
              </Button>
              {!!items.length &&
                <Button
                  variant="outlined"
                  style={{ marginLeft: 20, minWidth: 150, color: 'rgb(5, 23, 71)', borderColor: 'rgb(5, 23, 71)' }}
                  onClick={() => setOpenDialog(true)}
                >
                  {t('create-template')}
                </Button>
              }
              <Dialog onClose={() => setOpenDialog(false)} open={openDialog}>
                <div style={{ padding: 30 }}>
                  <div style={{ marginBottom: 20 }} >
                    <TextField
                      variant="standard"
                      label={t('template-name')}
                      style={{ marginRight: 30, width: 300 }}
                      defaultValue={templateName}
                      onChange={e => setTemplateName(e.target.value)}
                    />
                  </div>
                  <Button
                    style={{ minWidth: '100%', backgroundColor: 'rgb(5, 23, 71)', color: 'white' }}
                    variant="contained"
                    endIcon={<SendIcon />}
                    onClick={handleAddTemplate}
                  >
                    {t('create-template')}
                  </Button>
                </div>
              </Dialog>
            </div>
          </div>
        )
      }}
    </Formik>
  )
}

export default FormBody