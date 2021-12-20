import React from 'react'
import { useDispatch, useSelector } from 'react-redux'
import TemplateModal from './TemplateModal'
import { unSelect, modifyTemplate as modify } from '../../features/templateSlice'
import { modifyTemplate } from '../../services/templates'
import { useTranslation } from 'react-i18next'

const ModifyTemplate = ({ handleClose, show }) => {
  const template = useSelector(state => state.templates.selected)
  const dispatch = useDispatch()
  const { t } = useTranslation()

  if (!template) return <></>

  const initialValues = {
    name: template.name,
    descriptions: template.descriptions.map((d, id) => {
      return { name: d, id }
    })
  }

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
    <TemplateModal
      initialValues={initialValues}
      onSubmit={onSubmit}
      handleClose={() => {
        handleClose()
        dispatch(unSelect())
      }}
      show={show}
      title={`${t('template')} - ${template.name}`}
      submitButtonName={t('modify-template')}
    />
  )
}

export default ModifyTemplate
