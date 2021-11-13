import React, { useEffect } from 'react'
import { Formik } from 'formik'
import { useTranslation } from 'react-i18next'
import './index.css'
import { useHistory } from 'react-router'
import useLogin from '../../hooks/login'
import { useDispatch, useSelector } from 'react-redux'
import { setUser } from '../../features/userSlice'
import { Button, TextField } from '@material-ui/core'

const LoginForm = () => {
  const { t } = useTranslation('translation')
  const user = useSelector(state => state.users.currentUser)
  const dispatch = useDispatch()
  const { login, user: fetchedUser } = useLogin()
  const history = useHistory()

  useEffect(() => {
    if (fetchedUser && !fetchedUser.errors) {
      localStorage.setItem('app-token', fetchedUser.token)
      dispatch(setUser(fetchedUser))
    }
  }, [fetchedUser])

  useEffect(() => {
    if (user) history.push('/')
  }, [user])

  const validate = () => {

  }

  const onSubmit = async (values, formik) => {
    await login(values.username, values.password)
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
          {({ handleSubmit, handleChange }) => {
            return (
              <div className="loginform-body">
                <TextField
                  id="username"
                  variant="standard"
                  label={t('loginform-username')}
                  style={{ width: 400, marginBottom: 20 }}
                  onChange={handleChange}
                />
                <TextField
                  id="password"
                  variant="standard"
                  label={t('loginform-password')}
                  type="password"
                  onChange={handleChange}
                  style={{ width: 400 }}
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