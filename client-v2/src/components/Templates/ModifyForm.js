import { Formik } from 'formik'
import React from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { unSelect, modifyTemplate as modify } from '../../features/templateSlice'
import { useTranslation } from 'react-i18next'
import { modifyTemplate } from '../../services/templates'
import { Modal } from 'react-bootstrap'
import { Button, Form, Row, Col } from 'react-bootstrap'
import { RiDeleteBinLine } from 'react-icons/ri'

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
          <Modal
            show={show}
            onHide={handleClose}
            backdrop="static"
            keyboard={false}
          >
            <Modal.Header closeButton>
              <Modal.Title>{`${t('template')} - ${template.name}`}</Modal.Title>
            </Modal.Header>
            <Modal.Body>
              <Form.Group as={Row}>
                <Form.Label column sm="2" style={{ fontWeight: 'bold' }}>{t('name')}</Form.Label>
                <Col sm="10">
                  <Form.Control
                    id="name"
                    placeholder={t('name')}
                    value={values.name}
                    onChange={handleChange}
                  />
                </Col>
              </Form.Group>

              <Form.Label style={{ fontWeight: 'bold' }}>{t('descriptions')}</Form.Label>
              {values.descriptions.map((description, i) => {
                return (
                  <Form.Group key={i} as={Row} className="mb-3">
                    <Col sm="10">
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
                    </Col>
                    <Col sm="1">
                      <Button style={{ backgroundColor: 'red', borderColor: 'red' }}>
                        <RiDeleteBinLine onClick={() => {
                          setFieldValue('descriptions', values.descriptions.filter(value => {
                            return description.id !== value.id
                          }))
                        }}/>
                      </Button>
                    </Col>
                  </Form.Group>
                )}
              )}
              <Button
                className='button-menu'
                onClick={() => {
                  const desc = values.descriptions
                  const newTemplate = {
                    id: desc.length ? desc[desc.length - 1].id + 1 : 1,
                    name: ''
                  }
                  setFieldValue('descriptions', values.descriptions.concat(newTemplate))
                }}
                style={{ minWidth: 100 }}
              >
                {t('add-description')}
              </Button>
            </Modal.Body>
            <Modal.Footer>
              <Button
                style={{ backgroundColor: 'grey' }}
                onClick={handleClose}
              >{t('close')}</Button>
              <Button
                className='button-menu'
                onClick={handleSubmit}
                style={{ minWidth: 100 }}
              >
                {t('create-template')}
              </Button>
            </Modal.Footer>
          </Modal>
        )
      }}
    </Formik>
  )
}

export default ModifyForm