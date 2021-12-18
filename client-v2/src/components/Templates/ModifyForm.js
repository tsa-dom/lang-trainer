import { Formik } from 'formik'
import React from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { unSelect, modifyTemplate as modify } from '../../features/templateSlice'
import { useTranslation } from 'react-i18next'
import { modifyTemplate } from '../../services/templates'
import { Button, Form } from 'react-bootstrap'

const ModifyForm = () => {
  const template = useSelector(state => state.templates.selected)
  const dispatch = useDispatch()
  const { t } = useTranslation()

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
    <Formik
      initialValues={{
        name: template.name,
        descriptions: template.descriptions.map((d, id) => {
          return { name: d, id }
        })
      }}
      onSubmit={onSubmit}
    >
      {({ values, setFieldValue, handleChange }) => {
        return (
          <div>
            <div style={{ fontWeight: 'bold' }} className="words-header">{t('template')}</div>
            <Form.Group>
              <Form.Label>{t('name')}</Form.Label>
              <Form.Control
                id="name"
                placeholder={t('name')}
                value={values.name}
                onChange={handleChange}
                as="textarea"
              />
            </Form.Group>
            {values.descriptions.map(description => {
              return (
                <div key={description.id} style={{ display: 'flex' }}>
                  <Form.Control
                    id="name"
                    placeholder={t('description')}
                    value={description.name}
                    onChange={e => {
                      setFieldValue('descriptions', values.descriptions.map(value => {
                        if (description.id === value.id) return { ...value, name: e.target.value }
                        else return value
                      }))
                    }}
                  />{/*
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
                  /> */}
                  <Button
                    className='button-menu'
                    style={{ width: '100%' }}
                    type="submit"
                    onClick={() => {
                      setFieldValue('descriptions', values.descriptions.filter(value => {
                        return description.id !== value.id
                      }))
                    }}
                  >
                    {t('remove-description')}
                  </Button>
                </div>
              )}
            )}
            <div style={{ marginTop: 10 }}>
              {/* <Button
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
              </Button> */}
            </div>
          </div>
        )
      }}
    </Formik>
  )
}

export default ModifyForm