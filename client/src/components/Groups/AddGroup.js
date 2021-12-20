import React from 'react'
import { useTranslation } from 'react-i18next'
import { useDispatch } from 'react-redux'
import { addGroup } from '../../services/groups'
import GroupModal from './GroupModal'
import { addGroup as add } from '../../features/groupSlice'

const AddGroup = ({ handleClose, show }) => {
  const dispatch = useDispatch()
  const { t } = useTranslation()

  const initialValues = {
    name: '',
    description: ''
  }

  const onSubmit = async (values, formik) => {
    const { name, description } = values
    const res = await addGroup({ name, description })
    dispatch(add(res))
    formik.resetForm()
    handleClose()
  }

  return (
    <GroupModal
      initialValues={initialValues}
      onSubmit={onSubmit}
      handleClose={handleClose}
      show={show}
      title={t('add-group')}
      submitButtonName={t('add-group')}
    />
  )
}

export default AddGroup
