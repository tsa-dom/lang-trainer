import { Formik } from 'formik'
import React from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { unSelect, modifyTemplate as modify } from '../../features/templateSlice'
import { useTranslation } from 'react-i18next'
import { modifyTemplate } from '../../services/templates'
import { Modal } from 'react-bootstrap'
import { Button, Form } from 'react-bootstrap'

const ModifyForm = ({ show, handleClose }) => {
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
      {({ values, setFieldValue, handleChange, handleSubmit }) => {
        return (
          <Modal show={show} onHide={handleClose}>
            <Modal.Header closeButton>
              <Modal.Title>test</Modal.Title>
            </Modal.Header>
            <Modal.Body>
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
                    />
                    <Button
                      className='button-menu'
                      style={{ width: '100%' }}
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
            </Modal.Body>
            <Modal.Footer>
              <Button
                className='button-menu'
                style={{ width: '100%' }}
                type="submit"
                onClick={handleSubmit}
              >
                {t('create-template')}
              </Button>
              <Button
                className='button-menu'
                style={{ width: '100%' }}
                type="submit"
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
            </Modal.Footer>
          </Modal>
        )
      }}
    </Formik>
  )
}

export default ModifyForm