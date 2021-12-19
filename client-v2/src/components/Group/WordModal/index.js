import { Formik } from 'formik'
import React, { useEffect, useState } from 'react'
import { Button, Col, Form, Modal, Row } from 'react-bootstrap'
import { useTranslation } from 'react-i18next'
import { useDispatch, useSelector } from 'react-redux'
import Select from 'react-select'
import { setSelected, unSelect } from '../../../features/templateSlice'
import useFetch from '../../../hooks/fetcher'
import { RiDeleteBinLine } from 'react-icons/ri'

const WordModal = ({
  initialValues,
  onSubmit,
  open,
  onClose,
  title,
  initialItems
}) => {
  const { fetchTemplates } = useFetch()
  const template = useSelector(state => state.templates.selected)
  const templates = useSelector(state => state.templates.values)
  const [items, setItems] = useState(initialItems)
  const dispatch = useDispatch()
  const { t } = useTranslation()

  useEffect(() => {
    if (template) {
      setItems(template.descriptions.map((description, id) => {
        return {
          description,
          name: '',
          id
        }
      }))
    }
  }, [template])

  const handleTemplateSelect = id => {
    if (id !== '') {
      const select = templates.find(t => t.id === id)
      dispatch(setSelected(select))
    } else {
      setItems([])
      dispatch(unSelect())
    }
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
      initialValues={initialValues}
      onSubmit={(values) => {
        onSubmit({
          ...values,
          items
        })
        onClose()
      }}
    >
      {({ values, handleChange, handleSubmit }) => {
        return (
          <Modal
            size="lg"
            show={open}
            onHide={onClose}
            backdrop="static"
            keyboard={false}
          >
            <Modal.Header closeButton>
              <Modal.Title>{title}</Modal.Title>
            </Modal.Header>
            <Modal.Body>
              <Form.Group as={Row} className="mb-3">
                <Form.Label column sm="2" style={{ fontWeight: 'bold' }}>{t('template')}</Form.Label>
                <Col sm="10">
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
                        padding: 2
                      }),
                      option: styles => ({ ...styles })
                    }}
                    onChange={e => handleTemplateSelect(e.value)}
                  />
                </Col>
              </Form.Group>
              <Form.Group as={Row} className="mb-3">
                <Col sm="4">
                  <Form.Label column style={{ fontWeight: 'bold' }} >{t('name')}</Form.Label>
                  <Form.Control
                    id="name"
                    placeholder={t('name')}
                    value={values.name}
                    onChange={handleChange}
                  />
                </Col>
                <Col sm="8">
                  <Form.Label style={{ fontWeight: 'bold' }} column >{t('test')}</Form.Label>
                  <Form.Control
                    id="description"
                    placeholder={t('description')}
                    value={values.description}
                    onChange={handleChange}
                  />
                </Col>
              </Form.Group>
              <Form.Label style={{ fontWeight: 'bold' }}>{t('word-items')}</Form.Label>
              <Row>
                <Col sm="4">
                  <Form.Label column >{t('name')}</Form.Label>
                </Col>
                <Col sm="8">
                  <Form.Label column >{t('description')}</Form.Label>
                </Col>
              </Row>

              {items.map(item => {
                return (
                  <Form.Group key={item.id} as={Row} className="mb-3">
                    <Col sm="4">
                      <Form.Control
                        id="name"
                        placeholder={t('name')}
                        value={item.name}
                        onChange={e => handleModifyItem(item.id, e, 'name')}
                      />
                    </Col>
                    <Col sm="7">
                      <Form.Control
                        id="description"
                        placeholder={t('description')}
                        value={item.description}
                        onChange={e => handleModifyItem(item.id, e, 'description')}
                      />
                    </Col>
                    <Col sm="1">
                      <Button style={{ backgroundColor: 'red', borderColor: 'red' }}>
                        <RiDeleteBinLine onClick={() => handleRemoveItem(item.id)}/>
                      </Button>
                    </Col>
                  </Form.Group>
                )
              })}
              <Button className="button-menu" onClick={() => handleAddItem(values)}>
                {t('add-item')}
              </Button>
            </Modal.Body>
            <Modal.Footer>
              <Button
                style={{ backgroundColor: 'grey' }}
                onClick={onClose}
              >{t('close')}</Button>
              <Button
                className='button-menu'
                onClick={handleSubmit}
                style={{ minWidth: 100 }}
              >
                {title === t('add-word') ? t('add-word') : t('modify-word')}
              </Button>
            </Modal.Footer>
          </Modal>
        )
      }}
    </Formik>
  )
}

export default WordModal
