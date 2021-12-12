import React from 'react'
import { Formik } from 'formik'
import { useTranslation } from 'react-i18next'
import { useHistory } from 'react-router'
import { useDispatch } from 'react-redux'
import { setUser } from '../../features/userSlice'
import { Button, TextField } from '@material-ui/core'
import { login } from '../../services/users'

const LoginForm = () => {
  const { t } = useTranslation()
  const dispatch = useDispatch()
  const history = useHistory()

  const validate = () => {}

  const onSubmit = async (values, formik) => {
    const user = await login(values.username, values.password)
    if (user) {
      localStorage.setItem('app-token', user.token)
      dispatch(setUser(user))
      history.push('/')
    }
    formik.setFieldValue('username', '')
    formik.setFieldValue('password', '')
  }

  return (
    <>
      <div className="page-container-head">
        <div className="page-container-header">{t('loginform-header')}</div>
        <hr className="page-container-linebreak"></hr>
      </div>
      <div className="page-container-body">
        <Formik
          validate={validate}
          initialValues={{
            username: '',
            password: ''
          }}
          onSubmit={onSubmit}
        >
          {({ handleSubmit, handleChange, values }) => {
            return (
              <div className="loginform-body">
                <TextField
                  id="username"
                  variant="standard"
                  label={t('loginform-username')}
                  style={{ width: 400, marginBottom: 20 }}
                  onChange={handleChange}
                  value={values.username}
                />
                <TextField
                  id="password"
                  variant="standard"
                  label={t('loginform-password')}
                  type="password"
                  onChange={handleChange}
                  style={{ width: 400 }}
                  value={values.password}
                />
                <div style={{ marginTop: 20, marginBottom: 20 }}>
                  <Button
                    style={{ marginTop: 30, width: 400, backgroundColor: 'rgb(5, 23, 71)', color: 'white' }}
                    variant="contained"
                    onClick={handleSubmit}
                  >
                    {t('loginform-button')}
                  </Button>
                </div>
              </div>
            )
          }}
        </Formik>
      </div>
    </>
  )
}

export default LoginForm