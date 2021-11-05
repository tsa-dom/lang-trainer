import React, { useEffect } from 'react'
import { Formik } from 'formik'
import { useTranslation } from 'react-i18next'
import Input from '../Styled/Input'
import Button from '../Styled/Button'
import './index.css'
import { useHistory } from 'react-router'
import useLogin from '../../hooks/login'
import { useDispatch, useSelector } from 'react-redux'
import { setUser } from '../../features/userSlice'

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
          {({ handleSubmit, handleChange, values }) => {
            return (
              <div className="loginform-body">
                <Input
                  id="username"
                  className="loginform-input"
                  label={t('loginform-username')}
                  onChange={handleChange}
                  value={values.username}
                />
                <Input
                  id="password"
                  className="loginform-input"
                  label={t('loginform-password')}
                  type="password"
                  onChange={handleChange}
                  value={values.password}
                />
                <Button
                  className="loginform-submit-button"
                  onClick={handleSubmit}
                  text={t('loginform-button')}
                />
              </div>
            )
          }}
        </Formik>
      </div>
    </>
  )
}

export default LoginForm