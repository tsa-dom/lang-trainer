import React, { useEffect, useState } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { removeTemplates as remove, setSelected } from '../../features/templateSlice'
import ItemList from '../ItemList'
import { useTranslation } from 'react-i18next'
import { removeTemplates } from '../../services/templates'
import useFetch from '../../hooks/fetcher'
import ModifyForm from './ModifyForm'

const Templates = () => {
  const templates = useSelector(state => state.templates.values)
  const { fetchTemplates } = useFetch()
  const dispatch = useDispatch()
  const { t } = useTranslation()
  const [show, setShow] = useState(false)

  useEffect(fetchTemplates, [])

  const columns = [
    { field: 'name', headerName: t('name'), flex: 1 },
    { field: 'descriptions', headerName: t('descriptions'), flex: 3 }
  ]

  const handleTemplateRemove = async (values) => {
    if (values.length <= 0) return
    const ids = await removeTemplates({ templateIds: values })
    if (ids) dispatch(remove(ids))
  }

  const handleTemplateClick = (row) => {
    dispatch(setSelected({
      ...row,
      descriptions: row.descriptions.split(', ')
    }))
    setShow(true)
  }

  const handleClose = () => setShow(false)

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
      <ModifyForm handleClose={handleClose} show={show} />
    </>
  )
}

export default Templates