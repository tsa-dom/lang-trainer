/* eslint-disable no-unused-vars */
import React from 'react'
import { useTranslation } from 'react-i18next'
import { useDispatch } from 'react-redux'
import { addTemplate as add } from '../../features/templateSlice'
import { addTemplate } from '../../services/templates'
import TemplateModal from './TemplateModal'

const AddTemplate = ({ handleClose, show }) => {
  const dispatch = useDispatch()
  const { t } = useTranslation()

  const intialValues = {
    name: '',
    descriptions: []
  }

  const onSubmit = async (values, formik) => {
    const { name, descriptions } = values
    const res = await addTemplate({ name, descriptions: descriptions.map(d => d.name) })
    dispatch(add(res))
    formik.resetForm()
    handleClose()
  }

  return (
    <TemplateModal
      initialValues={intialValues}
      onSubmit={onSubmit}
      handleClose={handleClose}
      show={show}
      title={t('create-template')}
      submitButtonName={t('create-template')}
    />
  )
}

export default AddTemplate
