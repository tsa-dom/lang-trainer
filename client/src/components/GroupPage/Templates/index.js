import React, { useEffect } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import useTemplates from '../../../hooks/templates'
import { setTemplates, removeTemplates as remove, setSelected } from '../../../features/templateSlice'
import ItemList from '../../Styled/ItemList'
import { useTranslation } from 'react-i18next'
import ModifyForm from './ModifyForm'

const Templates = () => {
  const templates = useSelector(state => state.templates.values)
  const fetched = useSelector(state => state.templates.fetched)
  const { getTemplates, removeTemplates } = useTemplates()
  const dispatch = useDispatch()
  const { t } = useTranslation()

  useEffect(async () => {
    if (!fetched) {
      const templates = await getTemplates()
      if (templates) dispatch(setTemplates(templates))
    }
  }, [])

  const columns = [
    { field: 'name', headerName: t('name'), flex: 1 },
    { field: 'descriptions', headerName: t('descriptions'), flex: 3 }
  ]

  const handleTemplateRemove = async (values) => {
    if (values.length <= 0) return
    const ids = await removeTemplates({ templateIds: values })
    if (ids) dispatch(remove(ids))
  }

  const handleTemplateClick = (values) => {
    dispatch(setSelected({
      ...values.row,
      descriptions: values.row.descriptions.split(', ')
    }))
  }

  return (
    <>
      <ItemList
        rows={templates.map(template => {
          return { ...template, descriptions: template.descriptions.join(', ') }
        })}
        columns={columns}
        onCellClick={handleTemplateClick}
        handleItemRemove={handleTemplateRemove}
      />
      <ModifyForm />
    </>
  )
}

export default Templates