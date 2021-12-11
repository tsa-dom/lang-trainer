import React from 'react'
import { useTranslation } from 'react-i18next'
import { Formik } from 'formik'
import { addGroup as storeGroup } from '../../features/groupSlice'
import { useDispatch } from 'react-redux'
import { TextField } from '@mui/material'
import SendIcon from '@mui/icons-material/Send'
import { Button } from '@material-ui/core'
import { setNotification } from '../../features/notificationSlice'
import { addGroup } from '../../services/groups'

const AddForm = ({ setSelected }) => {
  const { t } = useTranslation()
  const dispatch = useDispatch()

  const validate = () => {}

  const onSubmit = async (values) => {
    const group = await addGroup(values)
    if (group) {
      if (group) dispatch(storeGroup(group))
      setSelected('groups-list')
      dispatch(setNotification({
        message: 'Added a new group successfully',
        type: 'success'
      }))
    } else {
      dispatch(setNotification({
        message: 'Server error',
        type: 'error'
      }))
    }
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
      {({ handleSubmit, handleChange }) => {
        return (
          <div className="groups-add-body">
            <TextField
              id="name"
              required
              variant="standard"
              label={t('name')}
              fullWidth
              style={{ marginBottom: 20 }}
              onChange={handleChange}
            />
            <TextField
              id="description"
              variant="standard"
              label={t('description')}
              multiline
              onChange={handleChange}
              style={{ marginBottom: 20 }}
              fullWidth
            />
            <Button
              style={{ marginTop: 30, minWidth: 150, backgroundColor: 'rgb(5, 23, 71)', color: 'white' }}
              variant="contained"
              endIcon={<SendIcon />}
              onClick={handleSubmit}
            >
              {t('groups-addnew')}
            </Button>
          </div>
        )
      }}
    </Formik>
  )
}

export default AddForm