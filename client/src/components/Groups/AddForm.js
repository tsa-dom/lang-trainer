import React from 'react'
import Input from '../Styled/Input'
import Button from '../Styled/Button'
import { useTranslation } from 'react-i18next'
import TextArea from '../Styled/TextArea'
import { Formik } from 'formik'
import useGroups from '../../hooks/groups'
import { addGroup as storeGroup } from '../../features/groupSlice'
import { useDispatch } from 'react-redux'

const AddForm = ({ setSelected }) => {
  const { t } = useTranslation('translation')
  const dispatch = useDispatch()
  const { addGroup } = useGroups()

  const validate = () => {}

  const onSubmit = async (values) => {
    const group = await addGroup(values)
    if (group) dispatch(storeGroup(group))
    setSelected('list')
  }

  return (
    <Formik
      validate={validate}
      initialValues={{
        name: '',
        description: ''
      }}
      onSubmit={onSubmit}
    >
      {({ handleSubmit, handleChange, values }) => {
        return (
          <div className="groups-add-body">
            <Input
              id="name"
              className="groups-add-input"
              label={t('groups-list-name')}
              onChange={handleChange}
              value={values.name}
            />
            <TextArea
              id="description"
              className="groups-add-textarea"
              label={t('groups-list-description')}
              onChange={handleChange}
              value={values.description}
            />
            <Button
              className="groups-add-button"
              text={t('groups-addnew')}
              onClick={handleSubmit}
            />
          </div>
        )
      }}
    </Formik>
  )
}

export default AddForm