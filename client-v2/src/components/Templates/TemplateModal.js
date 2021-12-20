import { Formik } from 'formik'
import React from 'react'
import { useTranslation } from 'react-i18next'
import { Modal } from 'react-bootstrap'
import { Button, Form, Row, Col } from 'react-bootstrap'
import { RiDeleteBinLine } from 'react-icons/ri'

const TemplateModal = ({
  show,
  handleClose,
  initialValues,
  onSubmit,
  title,
  submitButtonName
}) => {
  const { t } = useTranslation()

  return (

    <Formik
      initialValues={initialValues}
      onSubmit={onSubmit}
      enableReinitialize
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
              <Modal.Title>{title}</Modal.Title>
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

              <Form.Group>
                <Form.Label column as={Row} style={{ fontWeight: 'bold' }}>{t('descriptions')}</Form.Label>
              </Form.Group>
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
                      <Button style={{ backgroundColor: 'red', borderColor: 'red' }} onClick={() => {
                        setFieldValue('descriptions', values.descriptions.filter(value => {
                          return description.id !== value.id
                        }))
                      }}>
                        <RiDeleteBinLine />
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
                {submitButtonName}
              </Button>
            </Modal.Footer>
          </Modal>
        )
      }}
    </Formik>
  )
}

export default TemplateModal