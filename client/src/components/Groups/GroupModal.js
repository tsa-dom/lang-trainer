/* eslint-disable no-unused-vars */
import { Formik } from 'formik'
import React from 'react'
import { Button, Form, Modal } from 'react-bootstrap'
import { useTranslation } from 'react-i18next'

const GroupModal = ({
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
      {({ values, handleChange, handleSubmit }) => {
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
              <Form.Group>
                <Form.Label column sm="2" style={{ fontWeight: 'bold' }}>{t('name')}</Form.Label>
                <Form.Control
                  id="name"
                  placeholder={t('name')}
                  value={values.name}
                  onChange={handleChange}
                />
              </Form.Group>
              <Form.Group>
                <Form.Label column sm="2" style={{ fontWeight: 'bold' }}>{t('description')}</Form.Label>
                <Form.Control
                  id="description"
                  placeholder={t('description')}
                  value={values.description}
                  onChange={handleChange}
                />
              </Form.Group>
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

export default GroupModal
