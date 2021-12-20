import React from 'react'
import { Formik } from 'formik'
import { useTranslation } from 'react-i18next'
import { useNavigate } from 'react-router-dom'
import { useDispatch } from 'react-redux'
import { setUser } from '../../features/userSlice'
import { login } from '../../services/users'
import { Button, Container, Form } from 'react-bootstrap'

const Login = () => {
  const { t } = useTranslation()
  const dispatch = useDispatch()
  const navigate = useNavigate()

  const validate = () => { /* TODO */ }

  const onSubmit = async (values, formik) => {
    const user = await login(values.username, values.password)
    if (user) {
      localStorage.setItem('app-token', user.token)
      dispatch(setUser(user))
      navigate('/')
    }
    formik.setFieldValue('username', '')
    formik.setFieldValue('password', '')
  }

  return (
    <Container className="login-container">
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
              <Form.Group className='mb-3'>
                <Form.Label>{t('username')}</Form.Label>
                <Form.Control
                  id="username"
                  placeholder={t('username')}
                  value={values.username}
                  onChange={handleChange}
                />
              </Form.Group>
              <Form.Group className="mb-3">
                <Form.Label>{t('password')}</Form.Label>
                <Form.Control
                  id="password"
                  type="password"
                  placeholder={t('password')}
                  value={values.password}
                  onChange={handleChange}
                />
              </Form.Group>
              <br />
              <Button
                className='button-menu'
                style={{ width: '100%' }}
                type="submit"
                onClick={handleSubmit}
              >
                {t('login-button')}
              </Button>
            </div>
          )
        }}
      </Formik>
    </Container>
  )
}

export default Login