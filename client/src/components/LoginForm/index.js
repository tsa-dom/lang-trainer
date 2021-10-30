import React, { useEffect } from 'react'
import { Formik } from 'formik'
import { useTranslation } from 'react-i18next'
import Input from '../Styled/Input'
import Button from '../Styled/Button'
import './index.css'
import { useHistory } from 'react-router'
import useLogin from '../../hooks/login'

const LoginForm = ({ setCurrentUser }) => {
  const { t } = useTranslation('translation')
  const { login, result } = useLogin()
  const history = useHistory()

  useEffect(() => {
    if (result && !result.errors) {
      localStorage.setItem('app-token', result.token)
      setCurrentUser({
        username: result.username,
        priviledges: result.priviledges
      })
    }
  }, [result])

  const initialValues = {
    username: '',
    password: ''
  }

  const validate = (values) => {
    console.log(values)
  }

  const onSubmit = async (values) => {
    await login(values.username, values.password)
    history.push('/')
  }

  return (
    <div className="loginform-container">
      <div className="loginform-header">{t('loginform-header')}</div>
      <Formik
        validate={validate}
        initialValues={initialValues}
        onSubmit={onSubmit}
      >
        {({ handleSubmit, handleChange }) => {
          return (
            <div className="loginform-body">
              <Input
                id="username"
                className="loginform-input"
                label={t('loginform-username')}
                onChange={handleChange}
              />
              <Input
                id="password"
                className="loginform-input"
                label={t('loginform-password')}
                type="password"
                onChange={handleChange}
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
  )
}

export default LoginForm