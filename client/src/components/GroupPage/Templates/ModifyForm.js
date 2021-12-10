import { Dialog, TextField } from '@mui/material'
import { Formik } from 'formik'
import React from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { unSelect, modifyTemplate as modify } from '../../../features/templateSlice'
import { Button } from '@material-ui/core'
import { useTranslation } from 'react-i18next'
import SendIcon from '@mui/icons-material/Send'
import useTemplates from '../../../hooks/templates'

const ModifyForm = () => {
  const template = useSelector(state => state.templates.selected)
  const dispatch = useDispatch()
  const { t } = useTranslation()
  const { modifyTemplate } = useTemplates()

  if (!template) return <></>

  const onSubmit = async (values) => {
    const res = await modifyTemplate({
      id: template.id,
      name: values.name,
      descriptions: values.descriptions.map(d => d.name)
    })
    dispatch(modify(res))
    dispatch(unSelect())
  }

  return (
    <Dialog onClose={() => dispatch(unSelect())} open={!!template} >
      <div style={{ padding: 20 }}>
        <Formik
          initialValues={{
            name: template.name,
            descriptions: template.descriptions.map((d, id) => {
              return { name: d, id }
            })
          }}
          onSubmit={onSubmit}
        >
          {({ values, setFieldValue, handleSubmit, handleChange }) => {
            return (
              <div>
                <div style={{ fontWeight: 'bold' }} className="words-header">{t('template')}</div>
                <TextField
                  id="name"
                  required
                  defaultValue={values.name}
                  variant="standard"
                  label={t('name')}
                  style={{ marginRight: 30, marginBottom: 20 }}
                  onChange={handleChange}
                />
                {values.descriptions.map(description => {
                  return (
                    <div key={description.id} style={{ display: 'flex' }}>
                      <TextField
                        required
                        defaultValue={description.name}
                        variant="standard"
                        style={{ marginRight: 30, marginBottom: 10 }}
                        label={t('description')}
                        onChange={e => {
                          setFieldValue('descriptions', values.descriptions.map(value => {
                            if (description.id === value.id) return { ...value, name: e.target.value }
                            else return value
                          }))
                        }}
                      />
                      <Button
                        onClick={() => {
                          setFieldValue('descriptions', values.descriptions.filter(value => {
                            return description.id !== value.id
                          }))
                        }}
                        style={{ color: 'red' }}
                      >
                        {t('remove-description')}
                      </Button>
                    </div>
                  )}
                )}
                <div style={{ marginTop: 10 }}>
                  <Button
                    style={{ minWidth: 150, backgroundColor: 'rgb(5, 23, 71)', color: 'white' }}
                    variant="contained"
                    endIcon={<SendIcon />}
                    onClick={handleSubmit}
                  >
                    {t('create-template')}
                  </Button>
                  <Button
                    variant="outlined"
                    style={{ marginLeft: 20, minWidth: 150, color: 'rgb(5, 23, 71)', borderColor: 'rgb(5, 23, 71)' }}
                    onClick={() => {
                      const desc = values.descriptions
                      const newTemplate = {
                        id: desc.length ? desc[desc.length - 1].id + 1 : 1,
                        name: ''
                      }
                      setFieldValue('descriptions', values.descriptions.concat(newTemplate))
                    }}
                  >
                    {t('add-description')}
                  </Button>
                </div>
              </div>
            )
          }}
        </Formik>
      </div>
    </Dialog>
  )
}

export default ModifyForm